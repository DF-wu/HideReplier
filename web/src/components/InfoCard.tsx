import { Button } from "./Button";
import Swal from "../utils/sweetAlert2";
import { AdvancedInfo } from "../types";

export type InfoCardProps = AdvancedInfo;
export function InfoCard(props: InfoCardProps) {
  return (
    <div>
      <Button
        disabled={!props.ip}
        onClick={() => {
          Swal.fire({
            title: <h1 className="text-[22px] font-semibold">發文資訊</h1>,
            html: (
              <table className="w-full border-collapse text-left">
                <tbody>
                  <tr>
                    <td className="w-16">IP</td>
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
                  <tr>
                    <td>時區</td>
                    <td title={props.timezone}>{props.timezone}</td>
                  </tr>
                </tbody>
              </table>
            ),
          });
        }}
      >
        顯示更多資訊
      </Button>
    </div>
  );
}
export default InfoCard;
