import { page } from "$app/stores";
import type { Session } from "$lib/types/sessions"
import { apiURL, invalid_email, invalid_password, loggedInAs, session } from "../../stores/stores";
let s: Partial<Session> = {}

let origin: string
let baseURL = '';
let el: HTMLDivElement | null;

export async function SignIn(email: string, password: string) {
  const unsub = apiURL.subscribe((value) => {
    baseURL = value
  })
  unsub()

  const response = await fetch(`${baseURL}/public/account/signin`, {
    method: 'POST',
    mode: 'cors',
    cache: 'no-cache',
    credentials: 'include',
    headers: {
      'Content-Type': 'application/json',
    },
    redirect: 'follow',
    referrerPolicy: 'no-referrer',
    body: JSON.stringify({
      email: email,
      password: password
    })
  })

  const promise = await response.json();
  if (promise.message) {
    promise.message === 'invalid email' ? invalid_email.set(true) : invalid_password.set(true)
    return
  }

  console.log("promise: ", promise)
  s = promise[0];
  if (s === null) {
    return
  }

  session.set(s)
  window.localStorage.removeItem('session')
  window.localStorage.setItem('session', JSON.stringify(s))

  // MakeCheckMarkLotieVisible
  if (s.Account?.email) {
    loggedInAs.set(s.Account.email)
  }

  const getPageOrigin = page.subscribe((value) => {
    origin = value.url.origin
  })

  getPageOrigin()

  window.location.replace(`${origin}/appv1/my_links`)
}

