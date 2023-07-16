import { Meta, StoryObj } from "@storybook/react";
import { NavLink, NavLinkProps } from "../components/NavLink";

export default {
  component: NavLink,
} satisfies Meta<typeof NavLink>;

function NavLinkWrapper(props: NavLinkProps) {
  const { text, link, ...rest } = props;

  return <NavLink text={text || "Link"} link={link || "#"} {...rest} />;
}

export const Default: StoryObj<typeof NavLink> = {
  render: (args) => <NavLinkWrapper {...args} />,
};
