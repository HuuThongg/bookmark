import type { Folder } from "$lib/types/folder";
import { foldersCut, selectedFolders } from "../../stores/stores";

export function cutFolders(fs: Partial<Folder>[]) {
  emptyFoldersCut()
  foldersCut.update((values) => [...values, ...fs])
  selectedFolders.set([])
}
export function emptyFoldersCut() {
  foldersCut.set([])
}
