import { ReactNode } from "react";
import reactStringReplace from "react-string-replace";
import { cm } from "../../utils/tailwindMerge";

export type DiscordMentionProps = {
  content?: string;
};
export function Mention(props: DiscordMentionProps) {
  if (!props.content) return null;
  return (
    <span
      className={cm(
        "text-[#C9CDFB] bg-[rgba(88,101,242,0.3)] font-medium",
        "rounded-[3px] px-0.5"
      )}
    >
      {props.content}
    </span>
  );
}

const mentionMatcher = /(?<=^|\s)(@here|@everyone)/g;
export function highlightMentions(
  content: string | null | undefined
): ReactNode {
  if (!content) return content;
  return reactStringReplace(content, mentionMatcher, (matched) => (
    <Mention content={matched} />
  ));
}
