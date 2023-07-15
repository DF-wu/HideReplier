import { HTMLProps, ReactNode, useRef, useState } from "react";
import { cm } from "../../utils/tailwindMerge";
import { Embeded } from "../../types";
import { Image } from "./Image";
import { highlightMentions } from "./Mention";
import { toHTML } from "discord-markdown";

export type DiscordArticleProps = {
  trimColor?: string;
  topTitle?: string;
  mainTitle?: string;
  content?: string;
  embeded?: Embeded[];
  smallImage?: string;
  largeImage?: string;
} & HTMLProps<HTMLDivElement>;
export function Article(props: DiscordArticleProps) {
  const {
    trimColor = "#000",
    topTitle,
    mainTitle,
    content,
    embeded,
    smallImage,
    largeImage,
    className,
    ...rest
  } = props;

  const contentRef = useRef<HTMLDivElement>(null);
  const [preferedWidth, setPreferedWidth] = useState<
    { width: number } | undefined
  >(undefined);

  const highlighted = highlightMentions(content);
  const elements = mapMarkdownToElement(highlighted);
  console.log(elements);

  return (
    <div
      style={{
        backgroundColor: trimColor,
      }}
      className={cm(
        "text-[#F2F3F5] rounded",
        "w-fit max-w-[min(516px,100%-1rem)] min-w-[min(318px,100%-1rem)]",
        className
      )}
      {...rest}
    >
      <div
        style={largeImage ? preferedWidth : undefined}
        className={cm(
          "bg-[#2B2D31] text-[14px] pl-3 pr-4 py-4 flex flex-col gap-4",
          "h-full w-[100%-4px] relative left-1 rounded-r",
          "break-all"
        )}
      >
        <div className="flex flex-row gap-4 justify-between">
          <div className="flex flex-col gap-2">
            {topTitle && <div className="font-semibold">{topTitle}</div>}
            {mainTitle && (
              <div className="font-semibold text-[16px] text-[#09F]">
                {mainTitle}
              </div>
            )}
            {content && (
              <div ref={contentRef} className="text-[#ccc]">
                {mapMarkdownToElement(highlighted)}
              </div>
            )}
            {embeded && embeded.length > 0 && (
              <div className="flex flex-row gap-2">
                {embeded.map(({ field, value }, i) => (
                  <div
                    key={`${field}-${i}`}
                    className="flex flex-col gap-0.5 flex-grow"
                  >
                    <div className="font-semibold">{field}</div>
                    <div className="text-[#ccc]">{value}</div>
                  </div>
                ))}
              </div>
            )}
          </div>
          <Image url={smallImage} size="small" />
        </div>
        {largeImage && (
          <div>
            <Image
              onLoad={(e) => {
                const img = e.target as HTMLImageElement;
                const { width, height } = img.getBoundingClientRect();
                if (width < height) return setPreferedWidth(undefined);
                setPreferedWidth({ width: width + 28 });
              }}
              url={largeImage}
              size="large"
            />
          </div>
        )}
      </div>
    </div>
  );
}

function mapMarkdownToElement(content: ReactNode[]): ReactNode[] {
  return content.map((c, i) => {
    if (typeof c !== "string") return c;
    console.log(toHTML(c));
    return (
      <span key={`text-${i}`} dangerouslySetInnerHTML={{ __html: toHTML(c) }} />
    );
  });
}
