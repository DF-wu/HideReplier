import { Meta, StoryObj } from "@storybook/react";
import { Input, TextAreaProps } from "../components/Input";

export default {
  component: Input,
} satisfies Meta<typeof Input>;

export const Default: StoryObj<typeof Input> = {
  render: (args) => <Input {...args} containerClassName="max-w-[240px]" />,
};
export const TextArea: StoryObj<TextAreaProps> = {
  render: (args) => (
    <Input {...args} multiLine containerClassName="max-w-[240px]" />
  ),
};
