import { ButtonHTMLAttributes, PropsWithChildren } from "react";
import { cm } from "../utils/tailwindMerge";

export type ButtonVariant =
  | "default"
  | "primary"
  | "secondary"
  | "success"
  | "warning"
  | "danger"
  | "disabled";

export const buttonColors: Record<ButtonVariant, string> = {
  default:
    "bg-zinc-500 focus:bg-zinc-400 hover:bg-zinc-400 text-zinc-50 focus:ring-zinc-300",
  primary:
    "bg-indigo-600 focus:bg-indigo-500 hover:bg-indigo-500 text-indigo-50 focus:ring-indigo-300",
  secondary:
    "bg-slate-600 focus:bg-slate-500 hover:bg-slate-500 text-slate-50 focus:ring-slate-300",
  success:
    "bg-emerald-600 focus:bg-emerald-500 hover:bg-emerald-500 text-emerald-50 focus:ring-emerald-300",
  warning:
    "bg-amber-600 focus:bg-amber-500 hover:bg-amber-500 text-amber-50 focus:ring-amber-300",
  danger:
    "bg-red-600 focus:bg-red-500 hover:bg-red-500 text-red-50 focus:ring-red-300",
  disabled: "bg-gray-400 text-gray-200 cursor-not-allowed",
} as const;

export type ButtonProps = {
  variant?: ButtonVariant;
  fullWidth?: boolean;
} & PropsWithChildren<ButtonHTMLAttributes<HTMLButtonElement>>;

export function Button(props: ButtonProps) {
  const {
    variant = "default",
    className,
    children,
    disabled,
    fullWidth,
    ...rest
  } = props;
  return (
    <button
      className={cm(
        "rounded px-2.5 py-1 focus:outline-none focus:ring-2",
        buttonColors[disabled ? "disabled" : variant],
        {
          "w-full": fullWidth,
        },
        className
      )}
      disabled={disabled}
      {...rest}
    >
      {children}
    </button>
  );
}
