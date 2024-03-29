import { LinkIcon } from "@heroicons/react/24/outline";
import { cm } from "../utils/tailwindMerge";
import { AnchorHTMLAttributes } from "react";

export type NavLinkProps = {
  text: string;
  link: string;
  target?: AnchorHTMLAttributes<HTMLAnchorElement>["target"];
};
export function NavLink(props: NavLinkProps) {
  return (
    <a
      href={props.link}
      target={props.target}
      className={cm(
        "text-blue-300 flex flex-row items-center gap-0.5",
        "hover:text-blue-200 hover:underline",
        "focus:text-blue-200 focus:underline"
      )}
    >
      <LinkIcon className="h-4 w-4" />
      {props.text}
    </a>
  );
}
