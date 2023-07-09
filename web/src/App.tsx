import { useState } from "react";
import { Form } from "./components/Form";
import { InfoCard } from "./components/InfoCard";
import { NavLink } from "./components/NavLink";
import { useAdvancedInfo } from "./services/query";
import { cm } from "./utils/tailwindMerge";
import { PartialContent } from "./types";
import { Post } from "./components/Discord/Post";

function App() {
  const advancedInfo = useAdvancedInfo();
  const [content, setContent] = useState<PartialContent>({});

  return (
    <main
      className={cm(
        "w-full h-[100svh] bg-zinc-800 text-zinc-50",
        "flex overflow-auto flex-row justify-center"
      )}
    >
      <div
        className={cm(
          "relative py-4 top-[5svh] h-fit",
          "flex flex-col items-center justify-center gap-3"
        )}
      >
        <h1 className="text-[2rem] font-bold tracking-wide text-center">
          匿名回覆機器人
        </h1>
        <nav className="flex flex-row gap-5">
          <NavLink text="使用說明" link="https://hackmd.io/@dfder/ByADyqspI" />
          <NavLink text="GitHub" link="https://github.com/DF-wu/HideReplier" />
        </nav>

        <InfoCard {...advancedInfo} />

        <div className="flex flex-col gap-8 pt-2 justify-center sm:flex-row">
          <div className="flex-shrink-0">
            <Post
              avatar={content.avatar_url || "/icon.png"}
              topTitle="匿名機器人v0.4 (點我去發文)"
              mainTitle={content.username || ""}
              trimColor={content.color}
              content={content.content}
              smallImage={content.thumbnail}
              largeImage={content.imgUrl}
              embeded={[
                { field: "流水號", value: "XXXX" },
                { field: "來自：", value: content.ip || "" },
              ]}
            />
          </div>
          <div className="flex justify-center">
            <Form
              ip={advancedInfo.ip}
              onSubmit={() => undefined}
              onFormDataChanged={setContent}
            />
          </div>
        </div>
      </div>
    </main>
  );
}

export default App;
