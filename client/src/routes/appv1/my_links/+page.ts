import { browser } from "$app/environment";
import { redirect } from "@sveltejs/kit";
import type { PageLoad } from "./$types";
import type { Session } from "$lib/types/sessions";
import type { Folder } from "$lib/types/folder";
import type { Link } from "$lib/types/link";
import { apiURL } from "../../../stores/stores";
import { GetNewAccessToken } from "$lib/utils/getNewAccessToken";

let s: Partial<Session>
let folders: Partial<Folder>[] = [];
let links: Partial<Link>[] = []
let apiEndPoint: string

export const load: PageLoad = async ({ fetch, url }) => {
  if (browser) {
    console.log("page ts in browser mode")
    const sessionString: string | null = window.localStorage.getItem("session")
    console.log("session:\n", sessionString)
    if (sessionString === '' || sessionString === null) {
      window.localStorage.clear()
      console.log("url : ", url)
      redirect(302, `${url.origin}`)
    }
    if (sessionString) {
      console.log("sessionString: ", sessionString)
      s = JSON.parse(sessionString)
      const getApitEndPoint = apiURL.subscribe((value) => apiEndPoint = value)
      getApitEndPoint()
      const res = await fetch(`${apiEndPoint}/private/getLinksAndFolders/${s.Account?.id}/null`, {
        method: 'GET',
        mode: 'cors',
        credentials: 'include',
        headers: {
          'Content-Type': 'application/json',
          authorization: `Bearer${s.access_token}`
        }
      });

      const result = await res.json();
      console.log("result :", result)

      if (result.message) {
        // this token has expired
        console.log("This token has expired")
        //
        // get new session 
        const newSession: Partial<Session> = await GetNewAccessToken(fetch, url)
        console.log("new session: ", newSession)
        window.localStorage.removeItem('session')
        window.localStorage.setItem('session', JSON.stringify(newSession))
        console.log("refetch")
        const resp = await fetch(
          `${apiEndPoint}/private/getLinksAndFolders/${newSession.Account?.id}/null`,
          {
            method: 'GET',
            mode: 'cors',
            credentials: 'include',
            headers: {
              'Content-Type': 'application/json',
              authorization: `Bearer${newSession.access_token}`
            }
          }
        );
        const newResult = await resp.json();
        console.log(newResult)
        console.log("resul1t\n", newResult[0])
        folders = newResult[0].folders
        links = newResult.links
        return { links, folders }
      }

      console.log("result\n", result[0])
      folders = result[0].folders;
      links = result[0].folders;
      return { links, folders }
    }
  }
}
