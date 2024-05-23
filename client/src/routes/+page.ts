// since there's no dynamic data here, we can prerender 
//  it so that it gets served as a static asset in production 

import { browser } from '$app/environment'

let s: string | null

export async function load({ fetch, params, url }: any) {
  if (browser) {
    s = window.localStorage.getItem('session')
  }
  return { s }
}
export const prerender = true;
