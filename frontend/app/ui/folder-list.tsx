import FolderRow from "./folder-row";

interface Folder {
  id: number;
  name: string;
}

export default async function FolderList() {
  const backendUrl = process.env.NEXT_PUBLIC_BACKEND_URL;
  const foldersData = await fetch(backendUrl + "/folders");
  if (!foldersData.ok) {
    return <div>フォルダー一覧の取得に失敗しました</div>;
  }
  const folders: Folder[] = await foldersData.json();

  return (
    <div className="m-4">
      <ul className="divider-y divide-dotted divide-gray-600 divide-y-2">
        {folders.map((folder) => {
          return (
            <li key={folder.id} className="p-2">
              <FolderRow folder={folder} />
            </li>
          );
        })}
      </ul>
    </div>
  );
}
