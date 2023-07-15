import { useState } from "react";
import { cm } from "../utils/tailwindMerge";
import { Image } from "./Discord/Image";
import GifPicker, { Theme } from "gif-picker-react";

export type ImageSelectorTab = "built-in" | "tenor";

type ImageSelectorImplProps = {
  onImageSelected?: (imgUrl: string) => void;
};
function BuiltInSelector(props: ImageSelectorImplProps) {
  return (
    <div className="flex justify-between flex-wrap p-2 bg-zinc-800">
      {defaultImageList.map((img, i) => (
        <div
          key={i}
          onClick={() => props.onImageSelected?.(img)}
          className="hover:bg-zinc-500 p-1 cursor-pointer rounded-sm"
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
    <div>
      <ImageSelectorTabs tab={tab} setTab={setTab} />
      <div className="max-w-xs">{getSelector(tab, props.onImageSelected)}</div>
    </div>
  );
}

const defaultImageList: ReadonlyArray<string> = [
  "https://cdn.discordapp.com/attachments/591643710454890505/723846090247635084/kyaru01.gif",
  "https://cdn.discordapp.com/attachments/408969877580414976/714868499994116096/kyaru02.gif",
  "https://cdn.discordapp.com/attachments/591643710454890505/723846094353858580/kyaru03.gif",
  "https://cdn.discordapp.com/attachments/591643710454890505/723846093959594014/kyaru04.gif",
  "https://cdn.discordapp.com/attachments/591643710454890505/723846098371739728/kyaru08.gif",
  "https://cdn.discordapp.com/attachments/591643710454890505/723846100032946206/peko01.gif",
  "https://cdn.discordapp.com/attachments/591643710454890505/723846103006445618/ue01.gif",
] as const;
