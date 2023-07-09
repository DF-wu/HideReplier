import { Meta, StoryObj } from "@storybook/react";
import { Article, DiscordArticleProps } from "../../components/Discord/Article";

export default {
  component: Article,
} satisfies Meta<typeof Article>;

function ArticleWrapper(props: DiscordArticleProps) {
  return (
    <Article
      topTitle="匿名機器人v0.4 (點我去發文)"
      mainTitle="預設機器人:)"
      content="Test!!"
      embeded={[
        { field: "流水號", value: "3553" },
        { field: "來自：", value: "1.1.1.1" },
      ]}
      smallImage="https://cdn.discordapp.com/attachments/591643710454890505/723846094353858580/kyaru03.gif"
      largeImage="https://cdn.discordapp.com/attachments/591643710454890505/723846090247635084/kyaru01.gif"
      {...props}
    />
  );
}

export const Default: StoryObj<typeof Article> = {
  render: (args) => <ArticleWrapper {...args} />,
  argTypes: {
    trimColor: {
      type: "string",
      control: {
        type: "color",
      },
    },
    topTitle: {
      type: "string",
    },
    mainTitle: {
      type: "string",
    },
    content: {
      type: "string",
    },
    embeded: {
      control: {
        type: "object",
      },
    },
    smallImage: {
      control: {
        type: "select",
      },
      options: [
        null,
        "https://cdn.discordapp.com/attachments/591643710454890505/723846090247635084/kyaru01.gif",
        "https://cdn.discordapp.com/attachments/591643710454890505/723846094353858580/kyaru03.gif",
        "https://cdn.discordapp.com/attachments/680343695673131032/1127512955999748126/20230620_013954.jpg",
        "https://media.discordapp.net/attachments/680343695673131032/1049179889997185075/771c32d02f02a844f329e29b05674281.png",
        "https://cdn.discordapp.com/attachments/680343695673131032/1127512955664220241/black-21x9.png",
      ],
    },
    largeImage: {
      control: {
        type: "select",
      },
      options: [
        null,
        "https://cdn.discordapp.com/attachments/591643710454890505/723846090247635084/kyaru01.gif",
        "https://cdn.discordapp.com/attachments/591643710454890505/723846094353858580/kyaru03.gif",
        "https://cdn.discordapp.com/attachments/680343695673131032/1127512955999748126/20230620_013954.jpg",
        "https://media.discordapp.net/attachments/680343695673131032/1049179889997185075/771c32d02f02a844f329e29b05674281.png",
        "https://cdn.discordapp.com/attachments/680343695673131032/1127512955664220241/black-21x9.png",
      ],
    },
  },
};
