import { useState } from "react";
import { cm } from "../utils/tailwindMerge";
import { Image } from "./Discord/Image";
import GifPicker, { Theme } from "gif-picker-react";
import imageList from "../images.json";

const imgList: ReadonlyArray<string> = imageList;

export type ImageSelectorTab = "built-in" | "tenor";

type ImageSelectorImplProps = {
  onImageSelected?: (imgUrl: string) => void;
};
function BuiltInSelector(props: ImageSelectorImplProps) {
  return (
    <div className="flex justify-evenly flex-wrap p-2 max-h-80 overflow-auto">
      {imgList.map((img, i) => (
        <div
          key={i}
          onClick={() => props.onImageSelected?.(img)}
          className="hover:bg-zinc-500 p-1 cursor-pointer rounded-sm min-w-[88px] min-h-[88px]"
        >
          <Image url={img} size="small" />
        </div>
      ))}
    </div>
  );
}

function TenorSelector(props: ImageSelectorImplProps) {
  return (
    <div>
      <GifPicker
        tenorApiKey={import.meta.env.VITE_TENOR_API_KEY}
        theme={Theme.DARK}
        clientKey="hide-bot"
        width="auto"
        onGifClick={(gif) => props.onImageSelected?.(gif.url)}
      />
    </div>
  );
}

function getSelector(
  type: ImageSelectorTab,
  onSelected: ImageSelectorImplProps["onImageSelected"]
) {
  switch (type) {
    case "built-in":
      return <BuiltInSelector onImageSelected={onSelected} />;
    case "tenor":
      return <TenorSelector onImageSelected={onSelected} />;
  }
}

const tabs: ReadonlyArray<{
  id: ImageSelectorTab;
  text: string;
}> = [
  {
    id: "built-in",
    text: "Default",
  },
  {
    id: "tenor",
    text: "Tenor",
  },
] as const;
type ImageSelectorTabProps = {
  tab: ImageSelectorTab;
  setTab: (tab: ImageSelectorTab) => void;
};
function ImageSelectorTabs(props: ImageSelectorTabProps) {
  return (
    <div>
      {tabs.map((t) => (
        <button
          key={t.id}
          onClick={() => props.setTab(t.id)}
          className={cm("text-zinc-100 py-0.5 px-2 rounded-t", {
            "bg-zinc-500 text-zinc-50": t.id === props.tab,
            "hover:bg-zinc-600": t.id !== props.tab,
          })}
        >
          {t.text}
        </button>
      ))}
    </div>
  );
}

export type ImageSelectorProps = {
  onImageSelected?: (imgUrl: string) => void;
};
export function ImageSelector(props: ImageSelectorProps) {
  const [tab, setTab] = useState<ImageSelectorTab>("built-in");
  return (
    <div className="bg-zinc-700 rounded border-solid border border-zinc-600 shadow-lg">
      <ImageSelectorTabs tab={tab} setTab={setTab} />
      <div className="max-w-xs">{getSelector(tab, props.onImageSelected)}</div>
    </div>
  );
}
