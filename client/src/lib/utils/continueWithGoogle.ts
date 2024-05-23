import type { Session } from "$lib/types/sessions"
import { apiURL, session } from "../../stores/stores";
import { page } from '$app/stores'
let s: Partial<Session> = {}

let origin: string;

let baseURL = '';

export async function continueWithGoogle(v: any) {
  const unsub = apiURL.subscribe((value) => {
    baseURL = value;
  })

  unsub();
  const url = `${baseURL}/public/continueWithGoogle`;
  const res = await fetch(url, {
    method: 'POST',
    mode: 'cors',
    referrerPolicy: 'no-referrer-when-downgrade',
    headers: {
      'Content-Type': 'application/json'
    },
    body: JSON.stringify({
      name: v.name,
      email: v.email,
      picture: v.picture
    })
  })

  const data = await res.json();
  if (data[0] === null) return

  s = data[0]
  session.set(s);
  window.localStorage.setItem('session', JSON.stringify(data[0]))

  const getPageOrigin = page.subscribe((value) => {
    origin = value.url.origin;
  })
  getPageOrigin();
  window.location.replace(`${origin}/appv1/my_links`)
}
