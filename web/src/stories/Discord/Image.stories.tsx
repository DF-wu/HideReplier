import { Meta, StoryObj } from "@storybook/react";
import { Image } from "../../components/Discord/Image";

export default {
  component: Image,
} satisfies Meta<typeof Image>;

const images = [
  "https://cdn.discordapp.com/attachments/591643710454890505/723846090247635084/kyaru01.gif",
  "https://cdn.discordapp.com/attachments/591643710454890505/723846094353858580/kyaru03.gif",
  "https://cdn.discordapp.com/attachments/680343695673131032/1127512955999748126/20230620_013954.jpg",
  "https://media.discordapp.net/attachments/680343695673131032/1049179889997185075/771c32d02f02a844f329e29b05674281.png",
];

export const Default: StoryObj<typeof Image> = {
  render: (args) => <Image {...args} />,
  args: {
    size: "small",
    url: images[0],
  },
};

export const Large: StoryObj<typeof Image> = {
  render: (args) => <Image {...args} />,
  args: {
    size: "large",
    url: images[1],
  },
};

export const Avatar: StoryObj<typeof Image> = {
  render: (args) => <Image {...args} />,
  args: {
    size: "avatar",
    url: images[2],
  },
};
