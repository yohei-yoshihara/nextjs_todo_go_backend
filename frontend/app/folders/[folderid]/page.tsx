import { Folder } from "@/app/lib/definitions";
import TaskList from "@/app/ui/task-list";
import Container from "@/app/ui/container";
import Title from "@/app/ui/title";
import LinkButton from "@/app/ui/link-button";

export default async function TaskListPage({
  params,
}: {
  params: { folderid: string };
}) {
  const folderId = parseInt((await params).folderid);

  const backendUrl = process.env.NEXT_PUBLIC_BACKEND_URL;
  const folderData = await fetch(backendUrl + `/folders/${folderId}`);
  if (!folderData.ok) {
    return <div>フォルダ情報の取得に失敗しました</div>;
  }
  const folder: Folder = await folderData.json();

  return (
    <Container>
      <Title>{folder.name}のタスク一覧</Title>
      <TaskList folderId={folderId} />
      <div className="flow-root items-center justify-center">
        <div className="float-left">
          <LinkButton href={`/folders/${folderId}/tasks/create`}>
            作成
          </LinkButton>
        </div>
        <div className="float-right">
          <LinkButton href="/folders">フォルダ一覧へ戻る</LinkButton>
        </div>
      </div>
    </Container>
  );
}
