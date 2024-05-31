import { createFolderMode, folderName } from "../../stores/stores";
import { SwitchOffCreateMode } from './switchOffCreateMode'
export function showCreateFolderForm() {
  SwitchOffCreateMode()
  createFolderMode.set(true)
  folderName.set('Untitled collection')
  setTimeout(() => {
    console.log("highlightinital input content")
  }, 100)
}
