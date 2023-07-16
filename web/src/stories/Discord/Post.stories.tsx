import { Meta, StoryObj } from "@storybook/react";
import { Post } from "../../components/Discord/Post";
import { mockArcitle } from "./MockData";

export default {
  component: Post,
} satisfies Meta<typeof Post>;

export const Default: StoryObj<typeof Post> = {
  render: (args) => <Post {...args} />,
  args: {
    avatar:
      "https://media.discordapp.net/attachments/680343695673131032/1049179889997185075/771c32d02f02a844f329e29b05674281.png",
    ...mockArcitle,
  },
};
