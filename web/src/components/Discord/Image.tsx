import { CSSProperties, SyntheticEvent } from "react";

export type ImageSize = "large" | "small";

export type ImageProps = {
  url?: string;
  size: ImageSize;
  onLoad?: (e: SyntheticEvent<HTMLImageElement, Event>) => void;
};

const sizeMapping: Record<ImageSize, CSSProperties> = {
  small: {
    maxWidth: 80,
    maxHeight: 80,
    height: "fit-content",
  },
  large: {
    maxWidth: 400,
    maxHeight: 300,
  },
};

export function Image(props: ImageProps) {
  if (!props.url) return null;

  return (
    <img
      className="rounded-[3px]"
      style={{ ...sizeMapping[props.size] }}
      src={props.url}
      onLoad={props.onLoad}
    />
  );
}
