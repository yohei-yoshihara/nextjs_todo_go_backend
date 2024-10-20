import { ReactNode } from "react";

type Props = {
  children: ReactNode;
};
export default function Container(props: Props) {
  return <div className="m-5">{props.children}</div>;
}
