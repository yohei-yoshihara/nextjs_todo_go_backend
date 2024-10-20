import { Task } from "../lib/definitions";
import TaskRow from "./task-row";

type Props = {
  folderId: number;
};

export default async function TaskList(props: Props) {
  const backendUrl = process.env.NEXT_PUBLIC_BACKEND_URL;
  const tasksData = await fetch(
    backendUrl + `/tasks?folderId=${props.folderId}`
  );
  if (!tasksData.ok) {
    return <div>タスク一覧の取得に失敗しました</div>;
  }
  const tasks: Task[] = await tasksData.json();
  return (
    <div className="m-4">
      <ul className="divider-y divide-dotted divide-gray-600 divide-y-2">
        {tasks.map((task) => {
          return (
            <li key={task.id} className="p-2">
              <TaskRow folderId={props.folderId} task={task} />
            </li>
          );
        })}
      </ul>
    </div>
  );
}
