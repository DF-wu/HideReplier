import { Form } from "./components/Form";
import { InfoCard } from "./components/InfoCard";
import { NavLink } from "./components/NavLink";
import { cm } from "./utils/tailwindMerge";

function App() {
  return (
    <div
      className={cm(
        "w-full h-[100svh] bg-zinc-800 text-zinc-50",
        "flex flex-col items-center justify-center gap-3"
      )}
    >
      <h1 className="text-[2rem] font-bold tracking-wide">匿名回覆機器人</h1>
      <nav className="flex flex-row gap-5">
        <NavLink text="使用說明" link="https://hackmd.io/@dfder/ByADyqspI" />
        <NavLink text="GitHub" link="https://github.com/DF-wu/HideReplier" />
      </nav>

      <InfoCard
        ip="1.1.1.1"
        country="Taiwan"
        city="New Taipei City"
        isp="Data Communication Business Group"
      />

      <Form onSubmit={() => undefined} />
    </div>
  );
}

export default App;
