import GifPicker, { Theme } from "gif-picker-react";

export type TenorPickerProps = {
  tenorApiKey: string;
  onImageSelected?: (imgUrl: string) => void;
};

export function TenorPicker(props: TenorPickerProps) {
  return (
    <GifPicker
      tenorApiKey={props.tenorApiKey}
      theme={Theme.DARK}
      clientKey="hide-bot"
      width="auto"
      onGifClick={(gif) => props.onImageSelected?.(gif.url)}
    />
  );
}

export default TenorPicker;
