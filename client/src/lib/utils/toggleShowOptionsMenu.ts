import { showOptionsMenu } from "../../stores/stores";

let showOptionsMenuStatus: boolean;

export function toggleShowOptionsMenu() {
  const unsubscribe = showOptionsMenu.subscribe((value) => {
    showOptionsMenuStatus = value
  })
  unsubscribe();
  showOptionsMenuStatus = !showOptionsMenuStatus;
  showOptionsMenu.set(showOptionsMenuStatus)
}

export function hideShowOptionsMenu() {
  showOptionsMenu.set(false)
}
