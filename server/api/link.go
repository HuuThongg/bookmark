package api

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"go-bookmark/auth"
	"go-bookmark/db/sqlc"
	"go-bookmark/util"
	"go-bookmark/vultr"
	"log"
	"net/http"
	"net/url"
	"strings"
	"sync"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
)

type URL struct {
	URL      string `json:"url"`
	FolderId string `json:"folder_id"`
}

func (u URL) validate(requestValidateChan chan error) error {
	validationError := validation.ValidateStruct(&u,
		validation.Field(&u.URL, validation.Required.Error("url is required"), is.URL.Error("url must be a valid url")),
	)

	requestValidateChan <- validationError
	return validationError
}

func (h *BaseHandler) AddLink(w http.ResponseWriter, r *http.Request) {
	rBody := json.NewDecoder(r.Body)
	rBody.DisallowUnknownFields()

	var req URL

	if err := rBody.Decode(&req); err != nil {
		ErrorDecodingRequest(w, err)
		return
	}
	requestValidationChan := make(chan error, 1)
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		req.validate(requestValidationChan)
	}()
	validationError := <-requestValidationChan

	if validationError != nil {
		ErrorInvalidRequest(w, validationError)
		return
	}

	var host string
	var urlToOpen string

	if strings.Contains(req.URL, "?") {
		u, err := url.ParseRequestURI(req.URL)
		if err != nil {
			ErrorInternalServerError(w, err)
			return
		}
		if u.Scheme == "https" {
			host = u.Host
			urlToOpen = fmt.Sprintf(`%v`, u)
		} else {
			util.Response(w, "invalid url", http.StatusBadRequest)
			return
		}
	} else {
		parsedUrl, err := url.Parse(req.URL)
		if err != nil {
			ErrorInternalServerError(w, err)
			return
		}

		if parsedUrl.Scheme == "https" {
			host = parsedUrl.Host

			urlToOpen = req.URL
		} else {
			host = parsedUrl.String()

			urlToOpen = fmt.Sprintf(`https://%s`, req.URL)
		}
	}
	resp, err := http.Get(fmt.Sprintf("https://www.google.com/s2/favicons?domain=%v&sz=64", req.URL))
	if err != nil {
		util.Response(w, "something went wrong", http.StatusInternalServerError)
		return
	}

	var favicon string
	log.Printf("resp HEader get: %v", resp.Header.Get("content-location"))
	if err := util.DownloadFavicon(resp.Header.Get("content-location"), "favicon.icon"); err != nil {
		favicon = resp.Header.Get("content-location")
	}
	log.Printf("favicon : %v", favicon)

	log.Printf("host: %v", host)
	log.Printf("urlToOpen : %v", urlToOpen)

	if favicon == "" {
		urlFaviconChan := make(chan string, 1)
		wg.Add(1)
		go func() {
			defer wg.Done()
			vultr.UploadLinkFavicon(urlFaviconChan)
		}()
		favicon = <-urlFaviconChan
	}
	u := launcher.New().UserDataDir("~/.config/google-chrome").Leakless(true).NoSandbox(true).Headless(true).MustLaunch()
	browser := rod.New().ControlURL(u).MustConnect()
	page := browser.MustPage(urlToOpen).MustWaitLoad()

	defer browser.MustClose()

	var urlTitle string
	urlTiltleChan := make(chan string, 1)
	urlHeadingChan := make(chan string, 1)

	wg.Add(1)
	go func() {
		defer wg.Done()

		util.GetUrlTitle(page, urlTiltleChan)
	}()

	wg.Add(1)
	go func() {

		defer wg.Done()
		util.GetUrlHeading(page, urlHeadingChan)
	}()

	title := strings.TrimSpace(<-urlTiltleChan)
	heading := strings.TrimSpace(<-urlHeadingChan)

	if title != "" {
		if heading != "" {
			if len(heading) > len(title) {
				urlTitle = heading
			} else {
				urlTitle = title
			}
		} else {
			urlTitle = title
		}
	} else {
		if heading != "" {
			urlTitle = heading
		} else {
			urlTitle = req.URL
		}
	}

	payload := r.Context().Value("payload").(*auth.PayLoad)
	var folderID sql.NullString

	if req.FolderId != "" {
		folderID = sql.NullString{String: req.FolderId, Valid: true}
	}

	util.RodGetUrlScreenshot(page)

	urlScreenshotChan := make(chan string, 1)
	wg.Add(1)
	go func() {
		defer wg.Done()
		vultr.UploadLinkThumbnail(urlScreenshotChan)
	}()
	urlScreenshotLink := <-urlScreenshotChan
	stringChan := make(chan string, 1)
	wg.Add(1)
	go func() {
		defer wg.Done()
		util.RandomStringGenerator(stringChan)
	}()

	linkID := <-stringChan
	addLinkParams := sqlc.AddLinkParams{
		LinkID:        linkID,
		LinkTitle:     urlTitle,
		LinkHostname:  host,
		LinkUrl:       req.URL,
		LinkFavicon:   favicon,
		AccountID:     payload.AccountID,
		FolderID:      folderID,
		LinkThumbnail: urlScreenshotLink,
	}
	q := sqlc.New(h.db)
	link, err := q.AddLink(r.Context(), addLinkParams)
	if err != nil {
		ErrorInternalServer(w, err)
		return
	}
	util.JsonResponse(w, link)
	wg.Wait()
}
