import { Suspense, lazy } from "react";
import { cm } from "../../utils/tailwindMerge";
import type { DiscordArticleProps } from "./Article";
import { Image } from "./Image";

const Article = lazy(() => import("./Article"));

export type DiscordPostProps = {
  avatar: string;
} & DiscordArticleProps;
export function Post(props: DiscordPostProps) {
  const { avatar, ...rest } = props;

  return (
    <div className="flex flex-row gap-4">
      <div className="xs:block hidden">
        <Image url={avatar} size="avatar" />
      </div>
      <div className="flex flex-col gap-0.5">
        <div className="flex flex-row gap-2 items-center">
          <span className="inline-flex gap-1 items-start">
            <span className="text-[#F2F3F5] font-medium">{rest.mainTitle}</span>
            <span
              className={cm(
                "px-[4.4px] rounded-[3px] text-[10px] text-white bg-[#5865F2]"
              )}
            >
              機器人
            </span>
          </span>
          <span className="text-[12px] font-medium text-[#949BA4]">
            {new Date().toLocaleString()}
          </span>
        </div>
        <Suspense fallback={<span>Loading...</span>}>
          <Article {...rest} />
        </Suspense>
      </div>
    </div>
  );
}

export default Post;
