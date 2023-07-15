import { Meta, StoryObj } from "@storybook/react";
import { ImageInput } from "../components/ImageInput";

export default {
  component: ImageInput,
} satisfies Meta<typeof ImageInput>;

export const Default: StoryObj<typeof ImageInput> = {
  render: (args) => <ImageInput {...args} />,
  args: {
    label: "Image",
  },
};
