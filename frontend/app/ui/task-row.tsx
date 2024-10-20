import { Task } from "@/app/lib/definitions";
import { MdDelete } from "react-icons/md";
import { SiTask } from "react-icons/si";
import Link from "next/link";

type Props = {
  folderId: number;
  task: Task;
};

export default function TaskRow(props: Props) {
  const task = props.task;
  return (
    <div className="flex flex-row items-center">
      <SiTask className="text-green-200 mr-2" />
      <div className="w-32">
        <div className="truncate font-bold text-gray-300 hover:text-blue-300 mr-2">
          {task.title}
        </div>
      </div>
      <div>
        <Link href={`/folders/${props.folderId}/tasks/${task.id}/delete`}>
          <MdDelete className="text-red-400" />
        </Link>
      </div>
    </div>
  );
}
