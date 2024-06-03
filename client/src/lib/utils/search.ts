import type { Link } from "$lib/types/link";
import type { Folder } from "$lib/types/folder";
import { apiURL, foldersFound, linksFound, searching } from "../../stores/stores";
import { get } from "svelte/store";
import { getSession } from "./getSession";
let l: Partial<Link>[] = []
let f: Partial<Folder>[] = []
let baseUrl: string
export async function searchLinksAndFolders(searchPhrase: string) {
  searching.set(true)
  baseUrl = get(apiURL)

  const searchLinksPromise = await fetch(`${baseUrl}/private/link/searchLinks/${searchPhrase}`, {
    method: 'GET',
    mode: 'cors',
    credentials: 'include',
    headers: {
      'Content-Type': 'application/json',
      authorization: `Bearer${JSON.parse(getSession()).access_token}`
    }
  });

  const searchFoldersPromise = await fetch(
    `${baseUrl}/private/folder/searchFolders/${searchPhrase}`,
    {
      method: 'GET',
      mode: 'cors',
      credentials: 'include',
      headers: {
        'Content-Type': 'application/json',
        authorization: `Bearer${JSON.parse(getSession()).access_token}`
      }
    }
  );

  const searchResults = await Promise.all([searchLinksPromise, searchFoldersPromise]);

  const searchLinksRes = await searchResults[0].json();

  const searchFoldersRes = await searchResults[1].json();

  l = searchLinksRes[0];

  f = searchFoldersRes[0];

  if (l !== null) {
    linksFound.set(l);
  } else {
    linksFound.set([]);
  }

  if (f !== null) {
    foldersFound.set(f);
  } else {
    foldersFound.set([]);
  }

  searching.set(false);
}
