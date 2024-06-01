import type { Folder } from "$lib/types/folder";
import { apiURL, folders } from "../../stores/stores";
import { getSession } from "./getSession";
import { resetFoldersCut } from "./resetFoldersCut";
import { resetSelectedFolders } from "./resetSelectedFolders";
import { resetSelectedLinks } from "./resetSelectedLinks";
let baseURL: string;

let domFolders: Partial<Folder>[];

export async function moveFolders(folderz: Partial<Folder>[], folderID: string | undefined) {
  if (folderID) {
    await moveFoldersToAnotherFolder(folderz, folderID)
  }
}
async function moveFoldersToAnotherFolder(folderz: Partial<Folder>[], folderID: string) {
  const unsub = apiURL.subscribe((value) => {
    baseURL = value
  })
  unsub()
  const response = await fetch(`${baseURL}/private/folder/moveFolders`, {
    method: 'PATCH',
    mode: 'cors',
    cache: 'no-cache',
    credentials: 'include',
    headers: {
      'Content-Type': 'application/json',
      authorization: `Bearer${JSON.parse(getSession()).access_token}`
    },
    redirect: 'follow',
    referrerPolicy: 'same-origin',
    body: JSON.stringify({
      destination_folder_id: folderID,
      folder_ids: folderz.map((f) => f.folder_id)
    })
  })

  const result = await response.json()
  const fs: Partial<Folder>[] = result[0]

  for (let index = 0; index < fs.length; index++) {
    const element = fs[index]
    const getDomFolders = folders.subscribe((value) => { domFolders = value })
    getDomFolders()
    folders.set(domFolders.filter((dmf) => dmf.folder_id !== element.folder_id));
  }
  resetSelectedFolders();

  resetSelectedLinks();

  resetFoldersCut();

  // hideMoveItemsPopup();

}
