import { twMerge } from "tailwind-merge";
import cx, { ArgumentArray } from "classnames";

export function cm(...args: ArgumentArray): string {
  return twMerge(cx(args));
}
