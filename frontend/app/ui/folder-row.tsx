import { Folder } from "@/app/lib/definitions";
import Link from "next/link";
import { MdDelete } from "react-icons/md";
import { FaFolder } from "react-icons/fa";

type Props = {
  folder: Folder;
};

export default function FolderRow(props: Props) {
  const folder = props.folder;
  return (
    <div className="flex flex-row items-center">
      <FaFolder className="text-orange-200 mr-2" />
      <Link href={`/folders/${folder.id}`} className="w-32">
        <div className="truncate font-bold text-gray-300 hover:text-blue-300 mr-2">
          {folder.name}
        </div>
      </Link>
      <div>
        <MdDelete className="text-red-400" />
      </div>
    </div>
  );
}
