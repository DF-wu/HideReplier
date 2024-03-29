import { Suspense, lazy, useCallback, useState } from "react";
import { NavLink } from "./components/NavLink";
import { useAdvancedInfo } from "./services/query";
import { cm } from "./utils/tailwindMerge";
import { Content, PartialContent } from "./types";
import { sendPost } from "./services/api";
import Swal from "./utils/sweetAlert2";

const Post = lazy(() => import("./components/Discord/Post"));
const Form = lazy(() => import("./components/Form"));
const InfoCard = lazy(() => import("./components/InfoCard"));

function App() {
  const advancedInfo = useAdvancedInfo();
  const [content, setContent] = useState<PartialContent>({});

  const Preview = (
    <Suspense fallback={<span>Loading preview...</span>}>
      <Post
        avatar={content.avatar_url || "/icon.png"}
        topTitle={`匿名機器人v${import.meta.env.VITE_BOT_VERSION} (點我去發文)`}
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
    </Suspense>
  );

  const onSubmit = useCallback(async (post: Content) => {
    try {
      await sendPost(post);
      Swal.fire({
        icon: "success",
        title: "發送成功！",
      });
    } catch (e) {
      console.error(e);
      Swal.fire({
        icon: "error",
        title: "發生未知錯誤！",
      });
    }
  }, []);

  const onPreSubmit = async (post: Content) => {
    const result = await Swal.fire({
      icon: "question",
      title: "確定送出？",
      html: (
        <div className="text-left text-[1rem] mx-auto w-fit">{Preview}</div>
      ),
      showConfirmButton: true,
      confirmButtonText: "送出",
      showCancelButton: true,
      cancelButtonText: "取消",
    });
    if (!result.isConfirmed) return;
    onSubmit(post);
  };

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
          <NavLink
            target="_blank"
            text="使用說明"
            link="https://hackmd.io/@dfder/ByADyqspI"
          />
          <NavLink
            target="_blank"
            text="GitHub"
            link="https://github.com/DF-wu/HideReplier"
          />
        </nav>

        <Suspense fallback={<span>Loading info...</span>}>
          <InfoCard {...advancedInfo} />
        </Suspense>

        <div className="flex flex-col gap-8 pt-2 justify-center sm:flex-row">
          <div className="flex-shrink-0">{Preview}</div>
          <div className="flex justify-center">
            <Suspense fallback={<span>Loading form...</span>}>
              <Form
                ip={advancedInfo.ip}
                onSubmit={onPreSubmit}
                onFormDataChanged={setContent}
              />
            </Suspense>
          </div>
        </div>
      </div>
    </main>
  );
}

export default App;
