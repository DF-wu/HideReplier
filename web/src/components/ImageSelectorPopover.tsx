import * as Popover from "@radix-ui/react-popover";
import { ReactNode } from "react";
import { ImageSelector } from "./ImageSelector";

export type ImageSelectorPopoverProps = {
  onImageSelected?: (imgUrl: string) => void;
  open: boolean;
  setOpen: (open: boolean) => void;
  children?: ReactNode;
};

export function ImageSelectorPopover(props: ImageSelectorPopoverProps) {
  return (
    <Popover.Root open={props.open}>
      <Popover.Trigger>{props.children}</Popover.Trigger>
      <Popover.Portal>
        <Popover.Content
          onOpenAutoFocus={(e) => e.preventDefault()}
          onPointerDownOutside={() => props.setOpen(false)}
          side="bottom"
          align="start"
          sideOffset={5}
        >
          <ImageSelector onImageSelected={props.onImageSelected} />
        </Popover.Content>
      </Popover.Portal>
    </Popover.Root>
  );
}
