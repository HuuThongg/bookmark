import type { Folder } from "$lib/types/folder";
import type { Session } from "$lib/types/sessions";
import { get } from "svelte/store";
import { apiURL } from "../../stores/stores";
import { resetAncestorsOfCurrentFolder } from "./resetAncestorsOfCurrentFolder";

interface RouteParams {
  folder_id: string
}
let currentFolderAncestors: Partial<Folder>[] = []
export async function GetCurrentFolderAncestors(fetch: typeof window.fetch, params: RouteParams, access_token: string): Promise<Partial<Folder>[]> {
  const apiEndPoint = get(apiURL)
  const res2 = await fetch(
    `${apiEndPoint}/private/folder/getFolderAncestors/${params.folder_id}`,
    {
      method: 'GET',
      mode: 'cors',
      credentials: 'include',
      headers: {
        'Content-Type': 'application/json',
        authorization: `Bearer${access_token}`
      }
    }
  );

  const result2 = await res2.json();

  if (result2[0] === undefined) {
    console.log(result2.message);
  } else {
    currentFolderAncestors = result2[0];
    console.log("getFolderAncestors,", result2[0])

    resetAncestorsOfCurrentFolder();
  }
  return currentFolderAncestors;
}
