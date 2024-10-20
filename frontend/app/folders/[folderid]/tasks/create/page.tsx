import Container from "@/app/ui/container";
import CreateTaskForm from "@/app/ui/create-task-form";
import Title from "@/app/ui/title";

type Props = {
  params: { folderid: string };
};

export default async function CreateTaskPage({ params }: Props) {
  const folderId = parseInt((await params).folderid);
  return (
    <Container>
      <Title>新規タスクの作成</Title>
      <CreateTaskForm folderId={folderId} path={`/folders/${folderId}`} />
    </Container>
  );
}
