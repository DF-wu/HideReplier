import { useRef, useState } from "react";
import { ImageSelectorPopover } from "./ImageSelectorPopover";
import { Input, InputProps } from "./Input";

export type ImageInputProps = {
  onImageUrlChanged?: (url: string) => void;
} & InputProps;
export function ImageInput(props: ImageInputProps) {
  const { onChange, onImageUrlChanged, ...rest } = props;
  const [focus, setFocus] = useState(false);
  const [containerFocus, setContainerFocus] = useState(false);
  const ref = useRef<HTMLInputElement>(null);

  return (
    <ImageSelectorPopover
      open={focus || containerFocus}
      setHasFocus={setContainerFocus}
      onImageSelected={(url) => {
        if (ref.current) ref.current.value = url;
        setContainerFocus(false);
        setFocus(false);
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
        onFocus={() => setFocus(true)}
        onBlur={() => setTimeout(() => setFocus(false), 0)}
      />
    </ImageSelectorPopover>
  );
}
