import { Meta, StoryObj } from "@storybook/react";
import { Article } from "../../components/Discord/Article";
import { mockArcitle } from "./MockData";

export default {
  component: Article,
} satisfies Meta<typeof Article>;

export const Default: StoryObj<typeof Article> = {
  render: (args) => <Article {...args} />,
  args: mockArcitle,
};
