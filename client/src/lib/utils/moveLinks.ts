
import type { Link } from '$lib/types/link';
import { getSession } from './getSession';
import { apiURL, links } from '../../stores/stores';
import { resetLinksCut } from './resetLinksCut';
import { resetSelectedLinks } from './resetSelectedLinks';


let myLinks: Partial<Link>[];
let baseUrl: string;

export async function moveLinks(linkss: Partial<Link>[], folderID: string) {
  const unsub = apiURL.subscribe((value) => {
    baseUrl = value;
  });

  unsub();

  const response = await fetch(`${baseUrl}/private/link/move`, {
    method: 'PATCH', // *GET, POST, PUT, DELETE, etc.
    mode: 'cors', // no-cors, *cors, same-origin
    cache: 'no-cache', // *default, no-cache, reload, force-cache, only-if-cached
    credentials: 'include', // include, *same-origin, omit
    headers: {
      'Content-Type': 'application/json',
      authorization: `Bearer${JSON.parse(getSession()).access_token}`
      // 'Content-Type': 'application/x-www-form-urlencoded',
    },
    redirect: 'follow', // manual, *follow, error
    referrerPolicy: 'no-referrer', // no-referrer, *no-referrer-when-downgrade, origin, origin-when-cross-origin, same-origin, strict-origin, strict-origin-when-cross-origin, unsafe-url
    body: JSON.stringify({
      links: linkss.map((li) => li.link_id),
      folder_id: folderID
    }) // body data type must match "Content-Type" header
  });

  const result = await response.json();

  const linksMoved: Partial<Link>[] = result[0];


  for (let index = 0; index < linksMoved.length; index++) {
    const element = linksMoved[index];
    const getDomLinks = links.subscribe((value) => {
      myLinks = value;
    });

    getDomLinks();

    myLinks = myLinks.filter((ml) => ml.link_id !== element.link_id);

    links.set(myLinks);
  }

  resetSelectedLinks();

  resetLinksCut();

  // hideMoveItemsPopup();

}
