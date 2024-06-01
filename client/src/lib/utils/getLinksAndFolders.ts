import type { Folder } from "$lib/types/folder";
import type { Link } from "$lib/types/link";
import type { Session } from "$lib/types/sessions";
import { apiURL } from "../../stores/stores";

interface RouteParams {
  folder_id: string;
}

export interface GetLinksAndFoldersResponse {
  folders: Partial<Folder>[];
  links: Partial<Link>[];
}
let apiEndPoint: string

export async function GetLinksAndFolders(
  fetch: typeof window.fetch,
  params: RouteParams,
  newSession: Partial<Session>
): Promise<GetLinksAndFoldersResponse> {
  let folders: Partial<Folder>[] = [];
  let links: Partial<Link>[] = [];

  // Get API endpoint
  const getApiEndPoint = apiURL.subscribe((value) => {
    apiEndPoint = value;
  });

  getApiEndPoint();

  // Fetch links using new session
  const res = await fetch(
    `${apiEndPoint}/private/getLinksAndFolders/${newSession.Account?.id}/${params.folder_id}`,
    {
      method: 'GET',
      mode: 'cors',
      credentials: 'include',
      headers: {
        'Content-Type': 'application/json',
        authorization: `Bearer ${newSession.access_token}`, // Ensure a space after 'Bearer'
      },
    }
  );

  const newResult = await res.json();

  // Get folders and links from new result
  folders = newResult[0].folders;
  links = newResult[0].links;

  // Return as an object
  return { folders, links };
}
