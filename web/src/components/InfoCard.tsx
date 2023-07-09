import { useRef } from "react";
import { Button } from "./Button";
import { cm } from "../utils/tailwind-merge";

export type InfoCardProps = {
  ip: string;
  country: string;
  city: string;
  isp: string;
};
export function InfoCard(props: InfoCardProps) {
  const ref = useRef<HTMLDialogElement>(null);

  return (
    <div>
      <Button
        onClick={() => {
          if (!ref.current) return;
          ref.current.showModal();
        }}
      >
        顯示更多資訊
      </Button>
      <dialog
        ref={ref}
        className={cm(
          "m-auto bg-zinc-700 text-zinc-100 rounded-md",
          "shadow-md shadow-zinc-900"
        )}
      >
        <table className="max-w-[200px] border-collapse">
          <tbody>
            <tr>
              <td className="w-10">IP</td>
              <td title={props.ip}>{props.ip}</td>
            </tr>
            <tr>
              <td>國家</td>
              <td title={props.country}>{props.country}</td>
            </tr>
            <tr>
              <td>城市</td>
              <td title={props.city}>{props.city}</td>
            </tr>
            <tr>
              <td>ISP</td>
              <td title={props.isp}>{props.isp}</td>
            </tr>
          </tbody>
        </table>
        <form method="dialog" className="flex justify-end mt-2">
          <Button variant="success">OK</Button>
        </form>
      </dialog>
    </div>
  );
}
