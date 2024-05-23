import type { Session } from "$lib/types/sessions"
import { redirect } from "@sveltejs/kit";
import { apiURL } from "../../stores/stores"
// import type { PageLoad } from ".";
let apiEndPoint: string;
let ses: Partial<Session>;

export const GetNewAccessToken = async (fetch: typeof window.fetch, url: any): Promise<Partial<Session>> => {
  const getApiEndPoint = apiURL.subscribe((value) => { apiEndPoint = value })
  getApiEndPoint()
  const response = await fetch(`${apiEndPoint}/public/refreshToken`, {
    method: 'POST',
    mode: 'cors',
    cache: 'no-cache',
    credentials: 'include',
    redirect: 'follow',
    referrerPolicy: 'no-referrer',
  })

  const data = await response.json()
  if (data.message) {
    window.localStorage.clear()
    redirect(302, `${url.origin}/accounts/sign_in`)
  }
  console.log("create new GetNewAccessToken ", data)
  ses = data[0]
  return ses;
}
