import type { Folder } from "$lib/types/folder"
import { draggedFolder } from "../../stores/stores";
import { getDomFolders } from "./getDomFolders";

let el: HTMLElement
let f: Partial<Folder> = {}
let folderID: string;

let currentDomFolders: NodeListOf<HTMLDivElement>
export function FolderDragOver(e: DragEvent) {
  console.log("FolderDragOver")
  el = e.target as HTMLDivElement
  el = el.closest('.folder') as HTMLDivElement

  if (el.dataset.folderid) {
    folderID = el.dataset.folderid;
  }

  const unsubscribe = draggedFolder.subscribe((value) => { f = value })
  unsubscribe();

  currentDomFolders = getDomFolders()
  currentDomFolders.forEach((df) => {
    df.classList.remove('folder_drag_over')
  })
  if (folderID !== f.folder_id) {
    el.classList.add("folder_drag_over")
  }

}
