import Container from "@/app/ui/container";
import DeleteTaskForm from "@/app/ui/delete-task-form";
import Title from "@/app/ui/title";

type Props = {
  params: { folderid: string; taskid: string };
};

export default async function CreateTaskPage({ params }: Props) {
  const folderId = parseInt((await params).folderid);
  const taskId = parseInt((await params).taskid);
  const backendUrl = process.env.NEXT_PUBLIC_BACKEND_URL;
  const taskData = await fetch(backendUrl + `/tasks/${taskId}`);
  if (!taskData.ok) {
    return <div>タスクの読み込みに失敗しました</div>;
  }
  const task = await taskData.json();
  return (
    <Container>
      <Title>タスクの削除</Title>
      <DeleteTaskForm
        folderId={folderId}
        task={task}
        path={`/folders/${folderId}`}
      />
    </Container>
  );
}
