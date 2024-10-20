import { ReactNode } from "react";
import Link from "next/link";

type Props = {
  className?: string;
  href: string;
  children: ReactNode;
};

export default function LinkButton(props: Props) {
  return (
    <Link href={props.href}>
      <button className="bg-blue-500 p-2 rounded-xl">{props.children}</button>
    </Link>
  );
}
