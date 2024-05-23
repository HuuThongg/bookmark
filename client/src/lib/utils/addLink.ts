import type { Link } from "$lib/types/link";
import { getSession } from "./getSession";
import { addLinkMode, apiURL, links, newLink, successMessage } from "../../stores/stores";
import { get } from "svelte/store";
import { page } from "$app/stores";
import { goto } from "$app/navigation";

const myLinks: Link[] = []
let origin: string
let path: string
let baseURL: string
export let errorInvalidUrl = ''
export async function addLink(url: string, folderID: string) {
  if (url === 'https://example.com') {
    errorInvalidUrl = 'Please enter a valid url'
    return
  }

  // showLoadingToss()
  addLinkMode.set(false)
  const getApiPath = apiURL.subscribe((value) => { baseURL = value })
  getApiPath()


  const response = await fetch(`${baseURL}/private/link/add`, {
    method: 'POST', // *GET, POST, PUT, DELETE, etc.
    mode: 'cors', // no-cors, *cors, same-origin
    cache: 'no-cache', // *default, no-cache, reload, force-cache, only-if-cached
    credentials: 'include', // include, *same-origin, omit
    headers: {
      'Content-Type': 'application/json',
      authorization: `Bearer${JSON.parse(getSession()).access_token}`
      // 'Content-Type': 'application/x-www-form-urlencoded',
    },
    redirect: 'follow', // manual, *follow, error
    referrerPolicy: 'no-referrer', // no-referrer, *no-referrer-when-downgrade, origin, origin-when-cross-origin, same-origin, strict-origin, strict-origin-when-cross-origin, unsafe-url
    body: JSON.stringify({
      url: url,
      folder_id: folderID
    }) // body data type must match "Content-Type" header
  });
  newLink.set('')
  try {
    const result = await response.json()
    if (result[0] === null) {
      console.log("no result returned")
      return
    }
    if (result.message) {
      console.log(result.message)
    }
    const link: Link = result[0]
    if (get(links) !== null) {
      links.update((values) => [link, ...values])
    } else {
      links.set([...myLinks, link])
    }
    newLink.set('')
    const unsub = page.subscribe((values) => path = values.url.pathname)
    unsub()
    const getPageOrigin = page.subscribe((value) => {
      origin = value.url.origin
    })
    getPageOrigin()
    if (path === '/appv1/my_links/recycle_bin') {
      goto(`${origin}/appv1/my_links`)
    }
    successMessage.set("Link was added successfully")
    // hidingLoadingToss()
    // showSuccessNotification()
  } catch (error) {
    alert(error)

  }
}
