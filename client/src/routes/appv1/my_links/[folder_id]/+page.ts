
import { error, redirect } from '@sveltejs/kit';
import type { PageLoad } from "./$types";
import { browser } from '$app/environment';
import type { Session } from '$lib/types/sessions';
import type { Folder } from '$lib/types/folder';
import { apiURL } from '../../../../stores/stores';
import { GetNewAccessToken } from '$lib/utils/getNewAccessToken';
import { resetAncestorsOfCurrentFolder } from '$lib/utils/resetAncestorsOfCurrentFolder';
import type { Link } from '$lib/types/link';
import { GetLinksAndFolders } from '$lib/utils/getLinksAndFolders';
import { GetCurrentFolderAncestors } from '$lib/utils/getCurrentFolderAncestors';

let s: Partial<Session>
let folders: Partial<Folder>[] = [];
let links: Partial<Link>[] = [];
let currentFolderAncestors: Partial<Folder>[];
let apiEndPoint: string

export const load: PageLoad = async ({ params, fetch, url }) => {
  if (browser) {
    const sessionString: string | null = window.localStorage.getItem('session')
    if (sessionString === '' || sessionString === null) {
      window.localStorage.clear()
      redirect(302, `${url.origin}`)
    }
    if (sessionString) {
      s = JSON.parse(sessionString)
      const getApitEndPoint = apiURL.subscribe((value) => {
        apiEndPoint = value
      })
      getApitEndPoint()
      const res = await fetch(
        `${apiEndPoint}/private/getLinksAndFolders/${s.Account?.id}/${params.folder_id}`,
        {
          method: 'GET',
          mode: 'cors',
          credentials: 'include',
          headers: {
            'Content-Type': 'application/json',
            authorization: `Bearer${s.access_token}`
          }
        }
      );

      const result = await res.json();

      // if token has expired 
      if (result.message) {
        // this token has expired
        //alert(result.message);

        // get new session
        const newSession: Partial<Session> = await GetNewAccessToken(fetch, url);

        // add new session to local storage
        window.localStorage.removeItem('session');

        window.localStorage.setItem('session', JSON.stringify(newSession));


        const result = await GetLinksAndFolders(fetch, params, newSession)
        folders = result.folders
        links = result.links;

        (currentFolderAncestors = await GetCurrentFolderAncestors(fetch, params, newSession.access_token!))

        return { links, folders, currentFolderAncestors };

      }
      folders = result[0].folders;

      links = result[0].links;

      // get ancestors or current folder
      (currentFolderAncestors = await GetCurrentFolderAncestors(fetch, params, s.access_token!))

      return { links, folders, currentFolderAncestors };
    }
  }
}
