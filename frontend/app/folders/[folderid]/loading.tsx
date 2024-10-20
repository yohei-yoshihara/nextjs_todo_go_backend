import Container from "@/app/ui/container";

export default function Loading() {
  return (
    <Container>
      <div className="flex flex-col items-center justify-center">
        <div className="mb-4">読み込み中...</div>
        <div>
          <div className="animate-spin h-10 w-10 border-4 border-blue-500 rounded-full border-t-transparent"></div>
        </div>
      </div>
    </Container>
  );
}
