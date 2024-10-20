import { Suspense } from "react";
import FolderList from "../ui/folder-list";
import Container from "@/app/ui/container";
import Title from "@/app/ui/title";

export default function FoldersPage() {
  return (
    <Container>
      <Title>フォルダの一覧</Title>
      <Suspense fallback={<h2>読み込み中...</h2>}>
        <FolderList />
      </Suspense>
    </Container>
  );
}
