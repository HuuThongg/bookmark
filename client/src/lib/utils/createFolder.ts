import type { Folder } from "$lib/types/folder"
import { page } from "$app/stores";
import { get } from "svelte/store";
import { validateFolderName } from "./validationFolderName";
import { apiURL, createFolderMode, folderName, folders, lastCreatedFolder } from "../../stores/stores";
import { getSession } from "./getSession";
let folderID: string = ''
let myFolders: Folder[] = [];
let baseURL: string;
let path: string;
export async function CreateFolder(folder_name: string, parent_folder_id: string) {
  const err = validateFolderName(folder_name)
  console.log("create folder ,")
  console.log("folder_name: ", folder_name)
  if (err != '') {
    console.log(err)
    return
  }
  if (parent_folder_id === undefined) {
    folderID = 'null'
  } else {
    folderID = parent_folder_id
  }

  createFolderMode.set(false)
  const unsub = apiURL.subscribe((value) => { baseURL = value })
  unsub();
  const response = await fetch(`${baseURL}/private/folder/create`, {
    method: 'POST',
    mode: 'cors',
    credentials: 'include',
    headers: {
      'Content-Type': 'application/json',
      authorization: `Bearer${JSON.parse(getSession()).access_token}`
    },
    redirect: 'follow',
    referrerPolicy: 'no-referrer',
    body: JSON.stringify({
      folder_name: folder_name,
      parent_folder_id: folderID
    })
  })
  try {
    const result = await response.json()
    console.log("result : ,", result)
    const folder: Folder = result[0]
    console.log("folder : ", folder)
    if (folder === null) return
    if (get(folders) !== null) {
      folders.update((values) => [folder, ...values])
    } else {
      folders.set([...myFolders, folder])
    }

    lastCreatedFolder.set(folder)
    folderName.set('Untitled collection')
    const unsub = page.subscribe((values) => {
      path = values.url.pathname
    })
    unsub()

    // loading.set(false)
    // go to folder creawted?
  } catch (error) {
    console.log(error)
    // loading.set(false)
  }
}
