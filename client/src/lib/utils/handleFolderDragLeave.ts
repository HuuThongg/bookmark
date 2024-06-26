import { getDomFolders } from "./getDomFolders"

let currentDomFolders: NodeListOf<HTMLDivElement>
export function FolderDragLeave(e: DragEvent) {
  console.log("FolderDragLeave")
  currentDomFolders = getDomFolders()
  currentDomFolders.forEach((cdf) => {
    cdf.classList.remove('folder_drag_over')
  })
}
