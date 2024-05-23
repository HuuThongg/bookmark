import { browser } from "$app/environment";
import { redirect } from "@sveltejs/kit";

import type { PageLoad } from "./$types";
export const load: PageLoad = async ({ url }) => {
  if (browser) {
    const sessionString: string | null = window.localStorage.getItem('session')
    if (sessionString) {
      redirect(302, `${url.origin}/appv1/my_links`)
    }
  }
  return;
}
