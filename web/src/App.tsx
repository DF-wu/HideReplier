import { Form } from "./components/Form";
import { InfoCard } from "./components/InfoCard";
import { NavLink } from "./components/NavLink";
import { useAdvancedInfo } from "./services/query";
import { cm } from "./utils/tailwindMerge";

function App() {
  const advancedInfo = useAdvancedInfo();

  return (
    <main
      className={cm(
        "w-full h-[100svh] bg-zinc-800 text-zinc-50",
        "flex overflow-auto flex-col items-center"
      )}
    >
      <div
        className={cm(
          "relative py-4 top-[5svh]",
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

        <Form ip={advancedInfo.ip} onSubmit={() => undefined} />
      </div>
    </main>
  );
}

export default App;
