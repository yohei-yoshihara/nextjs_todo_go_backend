import { ReactNode } from "react";

type Props = {
  children: ReactNode;
};

export default function Title(props: Props) {
  return (
    <h1 className="text-xl text-gray-300 font-bold mb-3">{props.children}</h1>
  );
}
