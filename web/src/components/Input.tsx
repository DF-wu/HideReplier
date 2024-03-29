import { HTMLProps, Ref } from "react";
import { cm } from "../utils/tailwindMerge";

export type CommonInputProps = {
  containerClassName?: string;
  labelClassName?: string;
  fullWidth?: boolean;
};

export type InputProps = {
  multiLine?: false;
  inputRef?: Ref<HTMLInputElement>;
} & HTMLProps<HTMLInputElement> &
  CommonInputProps;

export type TextAreaProps = {
  multiLine: true;
  inputRef?: Ref<HTMLTextAreaElement>;
} & HTMLProps<HTMLTextAreaElement> &
  CommonInputProps;

export function Input(props: InputProps | TextAreaProps) {
  const {
    className,
    containerClassName,
    labelClassName,
    label,
    fullWidth,
    ...rest
  } = props;

  const inputClassName = cm(
    "bg-zinc-700 placeholder:text-zinc-400",
    "focus:outline-none rounded px-2 py-1",
    "border border-solid border-zinc-600",
    "hover:border-zinc-400 focus:border-zinc-400",
    {
      "w-full": fullWidth,
    },
    className
  );

  return (
    <div
      className={cm(
        "flex flex-col gap-1 text-zinc-100 items-start text-left",
        {
          "w-full": fullWidth,
        },
        containerClassName
      )}
    >
      {label && (
        <label
          className={cm(
            "text-[0.9rem] font-medium",
            { "w-full": fullWidth },
            labelClassName
          )}
        >
          {label}
        </label>
      )}

      {rest.multiLine ? (
        <textarea
          ref={rest.inputRef}
          className={cm("min-h-[2rem]", inputClassName)}
          {...omitMultiLineProp(rest)}
        />
      ) : (
        <input
          ref={rest.inputRef}
          className={inputClassName}
          {...omitMultiLineProp(rest)}
        />
      )}
    </div>
  );
}

function omitMultiLineProp<
  T extends { inputRef?: unknown; multiLine?: unknown }
>(p: T): Omit<T, "multiLine" | "inputRef"> {
  const { multiLine, inputRef, ...rest } = p;
  return rest;
}
