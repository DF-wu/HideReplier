import { Meta, StoryObj } from "@storybook/react";
import { ImageSelector } from "../components/ImageSelector";

export default {
  component: ImageSelector,
} satisfies Meta<typeof ImageSelector>;

export const Default: StoryObj<typeof ImageSelector> = {
  render: (args) => <ImageSelector {...args} />,
};
