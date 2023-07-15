import { useRef, useState } from "react";
import { ImageSelectorPopover } from "./ImageSelectorPopover";
import { Input, InputProps } from "./Input";

export type ImageInputProps = {
  onImageUrlChanged?: (url: string) => void;
} & InputProps;
export function ImageInput(props: ImageInputProps) {
  const { onChange, onImageUrlChanged, ...rest } = props;
  const [open, setOpen] = useState(false);
  const ref = useRef<HTMLInputElement>(null);

  return (
    <ImageSelectorPopover
      open={open}
      setOpen={setOpen}
      onImageSelected={(url) => {
        if (ref.current) ref.current.value = url;
        setOpen(false);
        onImageUrlChanged?.(url);
      }}
    >
      <Input
        {...rest}
        inputRef={ref}
        type="url"
        onChange={(e) => {
          if (ref.current) onImageUrlChanged?.(ref.current.value);
          onChange?.(e);
        }}
        onFocus={() => setOpen(true)}
      />
    </ImageSelectorPopover>
  );
}
