import { CSSProperties, ImgHTMLAttributes, SyntheticEvent } from "react";
import { cm } from "../../utils/tailwindMerge";
import { LazyLoadImage } from "react-lazy-load-image-component";

export type ImageSize = "large" | "small" | "avatar";

export type ImageProps = {
  url?: string;
  size: ImageSize;
  onLoad?: (e: SyntheticEvent<HTMLImageElement, Event>) => void;
  loading?: ImgHTMLAttributes<HTMLImageElement>["loading"];
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
  avatar: {
    width: 40,
    height: 40,
    objectFit: "cover",
  },
};

export function Image(props: ImageProps) {
  if (!props.url) return null;

  return (
    <LazyLoadImage
      className={cm({
        "rounded-[3px]": props.size !== "avatar",
        "rounded-full": props.size === "avatar",
      })}
      style={{ ...sizeMapping[props.size] }}
      src={props.url}
      onLoadedData={props.onLoad}
      loading={props.loading || "lazy"}
      alt={props.url}
    />
  );
}
