import { DiscordArticleProps } from "../../components/Discord/Article";

export const mockArcitle = {
  trimColor: "#000",
  topTitle: "匿名機器人v0.4 (點我去發文)",
  mainTitle: "預設機器人:)",
  content: "Test!!",
  embeded: [
    { field: "流水號", value: "3553" },
    { field: "來自：", value: "1.1.1.1" },
  ],
  smallImage:
    "https://cdn.discordapp.com/attachments/591643710454890505/723846094353858580/kyaru03.gif",
  largeImage:
    "https://cdn.discordapp.com/attachments/591643710454890505/723846090247635084/kyaru01.gif",
} satisfies DiscordArticleProps;
