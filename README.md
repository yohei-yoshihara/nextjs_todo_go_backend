# Go のバックエンドへ接続する Next.js ToDo アプリ

Go のバックエンドが API サーバとして動作し、そこへ Next.js アプリが接続する。
Go バックエンドはリバースプロキシとしても動作するので、Next.js へはリバースプロキシを経由してアクセスする。

## 起動

### バックエンド

```bash
cd backend
go build
./server seed
./server serve
```

### フロントエンド

```bash
cd frontend
pnpm install
pnpm run dev
```

`http://localhost:8000`へアクセスする。
（バックエンド内で動作しているリバースプロキシへ接続するため 3000 ポートでないことに注意）
