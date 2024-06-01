import type { Folder } from "$lib/types/folder";
import type { Link } from "$lib/types/link";
import { draggedFolder, draggedLink, selectedFolders, selectedLinks } from "../../stores/stores";
import { getDomFolders } from "./getDomFolders";
import { moveFolders } from "./moveFolderToAnotherFolder";

let el: HTMLElement;
let f: Partial<Folder> = {}
let l: Partial<Link> = {}
let folderID = ''

const foldersToMove: Partial<Folder>[] = []
const linksToMove: Partial<Link>[] = []

let currentDomFolders: NodeListOf<HTMLDivElement>
export async function ListenToDrop(e: DragEvent) {
  console.log("ListenToDrop")
  selectedFolders.set([])
  selectedLinks.set([])

  el = e.target as HTMLDivElement;
  el = el.closest('.folder') as HTMLDivElement

  const unsubscribe = draggedFolder.subscribe((value) => {
    f = value
  })
  const unsub = draggedLink.subscribe((value) => { l = value })
  unsubscribe();
  unsub()

  if (el.dataset.folderid) {
    folderID = el.dataset.folderid;
  }
  console.log("f.folder_id: ", f.folder_id, " folderID ", folderID)
  if (f.folder_id && l.link_id === undefined) {
    if (f.folder_id !== folderID) {
      await moveFolders([...foldersToMove, f], folderID)
    }
  } else if (l.link_id && f.folder_id === undefined) {
    // movelinks
  }

  draggedFolder.set({})
  draggedLink.set({})

  currentDomFolders = getDomFolders()
  currentDomFolders.forEach((cdf) => {

    cdf.classList.remove('folder_drag_over');

    cdf.classList.remove('dragged_folder');
  })
}
