import type { Meta, StoryObj } from "@storybook/react";
import { Button } from "../components/Button";

export default {
  component: Button,
} satisfies Meta<typeof Button>;

export const Default: StoryObj<typeof Button> = {
  render: (args) => (
    <Button variant="default" {...args}>
      Button
    </Button>
  ),
};

export const Primary: StoryObj<typeof Button> = {
  render: (args) => (
    <Button variant="primary" {...args}>
      Button
    </Button>
  ),
};

export const Secondary: StoryObj<typeof Button> = {
  render: (args) => (
    <Button variant="secondary" {...args}>
      Button
    </Button>
  ),
};

export const Success: StoryObj<typeof Button> = {
  render: (args) => (
    <Button variant="success" {...args}>
      Button
    </Button>
  ),
};

export const Warning: StoryObj<typeof Button> = {
  render: (args) => (
    <Button variant="warning" {...args}>
      Button
    </Button>
  ),
};

export const Danger: StoryObj<typeof Button> = {
  render: (args) => (
    <Button variant="danger" {...args}>
      Button
    </Button>
  ),
};

export const Disabled: StoryObj<typeof Button> = {
  render: (args) => (
    <Button disabled {...args}>
      Button
    </Button>
  ),
};
