import type { newUser } from "$lib/types/newUser";
import type { Session } from "$lib/types/sessions";
import { accessToken, apiURL, email_exists, session } from "../../stores/stores";
let baseURL = '';
let s: Partial<Session> = {}
export async function createNewAccount(a: Partial<newUser>) {

  console.log("sign up with email", a)
  const unsub = apiURL.subscribe((value) => { baseURL = value })
  unsub()

  const response = await fetch(`${baseURL}/public/account/create`, {
    method: 'POST',
    mode: 'cors',
    cache: 'no-cache',
    headers: {
      'Content-Type': 'applicaton/json'
    },
    redirect: 'follow',
    referrerPolicy: 'no-referrer',
    body: JSON.stringify({
      full_name: a.full_name,
      email_address: a.email_address,
      password: a.password
    })
  })
  try {
    const data = await response.json()

    if (data.message) {
      if (data.message === 'email already exists') {
        email_exists.set(true)
      } else {
        console.log(data.message)
      }
      return
    }
    s = data[0]
    if (s === null) return
    if (s.access_token) {
      accessToken.set(s.access_token)
    }

    session.set(s)
    window.localStorage.removeItem('session')
    window.localStorage.setItem('session', JSON.stringify(s))
    window.location.reload()
  } catch (error) {

    console.log(error)
  }
}
