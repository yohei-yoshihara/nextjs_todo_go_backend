"use client";

import Form from "next/form";
import { useActionState } from "react";
import LinkButton from "./link-button";
import Spinning from "./spinning";

type Props = {
  folderId: number;
  path: string;
};

type ActionState = {
  finished: boolean;
  errorMessage?: string;
};

export default function CreateTaskForm(props: Props) {
  async function handleAction(prevState: ActionState, formData: FormData) {
    const backendUrl = process.env.NEXT_PUBLIC_BACKEND_URL;
    const resp = await fetch(backendUrl + "/tasks/create", {
      method: "POST",
      body: JSON.stringify({
        title: formData.get("title"),
        folder_id: props.folderId,
      }),
    });
    if (!resp.ok) {
      return {
        finished: true,
        errorMessage: "サーバーからエラーが返されました",
      };
    }
    return { finished: true };
  }

  const [state, formAction, isPending] = useActionState<ActionState, FormData>(
    handleAction,
    {
      finished: false,
    }
  );

  if (state.finished && !state.errorMessage) {
    return (
      <div>
        <p className="mb-5">新しいタスクを作成しました。</p>
        <LinkButton href={props.path}>タスク一覧へ戻る</LinkButton>
      </div>
    );
  }

  if (state.finished && state.errorMessage) {
    return (
      <div>
        <p>新しいタスクの作成に失敗しました。</p>
        <p className="mb-5">{state.errorMessage}</p>
        <LinkButton href={props.path}>タスク一覧へ戻る</LinkButton>
      </div>
    );
  }

  return (
    <Form action={formAction}>
      <div className="flex flex-row justify-center items-center mb-4">
        <div className="w-1/3">
          <label htmlFor="title">タイトル</label>
        </div>
        <div className="w-2/3">
          <input
            id="title"
            name="title"
            type="text"
            required
            className="bg-gray-400 text-gray-700 p-2 rounded-lg focus:outline-none focus:outline-blue-500"
          />
        </div>
      </div>
      <div className="flex flex-row mb-5">
        <div className="w-1/3"></div>
        <div className="w-2/3">
          <div className="flow-root">
            <button
              type="submit"
              className="float-left bg-blue-500 disabled:bg-gray-600 p-2 rounded-xl"
              disabled={isPending}>
              作成
            </button>
            <div className="float-right">
              <LinkButton href={props.path}>戻る</LinkButton>
            </div>
          </div>
        </div>
      </div>
      {isPending && <Spinning />}
    </Form>
  );
}
