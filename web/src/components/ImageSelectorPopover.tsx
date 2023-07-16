import * as Popover from "@radix-ui/react-popover";
import { ReactNode } from "react";
import { ImageSelector } from "./ImageSelector";

export type ImageSelectorPopoverProps = {
  onImageSelected?: (imgUrl: string) => void;
  open: boolean;
  setHasFocus?: (focus: boolean) => void;
  children?: ReactNode;
};

export function ImageSelectorPopover(props: ImageSelectorPopoverProps) {
  return (
    <Popover.Root open={props.open}>
      <Popover.Trigger asChild>
        <button tabIndex={-1} className="cursor-default">
          {props.children}
        </button>
      </Popover.Trigger>
      <Popover.Portal>
        <Popover.Content
          onFocus={() => props.setHasFocus?.(true)}
          onInteractOutside={() => props.setHasFocus?.(false)}
          onOpenAutoFocus={(e) => e.preventDefault()}
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
