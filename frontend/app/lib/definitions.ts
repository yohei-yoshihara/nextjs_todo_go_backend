export interface Folder {
  id: number;
  name: string;
}

export interface Task {
  id: number;
  title: string;
  folderId: number;
}
