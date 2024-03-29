import { useEffect, useRef, useState } from "react";
import { Content, PartialContent } from "../types";
import { Button } from "./Button";
import { Input } from "./Input";
import { str2rgb } from "../utils/string2rgb";
import { ImageInput } from "./ImageInput";

export type FormProps = {
  ip: string;
  onSubmit?: (content: Content) => void;
  onFormDataChanged?: (content: PartialContent) => void;
};
export function Form(props: FormProps) {
  const formRef = useRef<HTMLFormElement>(null);
  const getFormData = () => {
    if (!formRef.current) return;
    return {
      ...Object.fromEntries(new FormData(formRef.current)),
      ip: props.ip,
    };
  };

  const onFormDataChanged = () => {
    const partialContent = PartialContent.safeParse(getFormData());
    if (!partialContent.success) return console.error(partialContent.error);
    props.onFormDataChanged?.(partialContent.data);
  };

  const [color, setColor] = useState<string | undefined>(
    props.ip ? str2rgb(props.ip) : ""
  );
  useEffect(() => {
    if (props.ip && !color) setColor(str2rgb(props.ip));
    setTimeout(() => onFormDataChanged(), 0);
  }, [props.ip]);

  return (
    <form
      ref={formRef}
      onSubmit={(e) => {
        e.preventDefault();
        if (!formRef.current) return;
        const content = Content.safeParse(getFormData());
        if (!content.success) return console.error(content.error);
        props.onSubmit?.(content.data);
      }}
      onChange={onFormDataChanged}
      className="flex flex-col gap-3 w-[min(100%-1rem,300px)]"
    >
      <Input
        fullWidth
        name="username"
        required
        label="機器人名稱*"
        defaultValue="預設機器人:)"
      />
      <ImageInput
        onImageUrlChanged={onFormDataChanged}
        fullWidth
        name="avatar_url"
        label="頭像連結"
      />
      <ImageInput
        onImageUrlChanged={onFormDataChanged}
        fullWidth
        name="thumbnail"
        label="小圖連結"
      />
      <ImageInput
        onImageUrlChanged={onFormDataChanged}
        fullWidth
        name="imgUrl"
        label="圖片連結"
      />
      <Input
        fullWidth
        name="color"
        type="color"
        label="顏色（裝飾）"
        value={color}
        onChange={(e) => setColor(e.currentTarget.value)}
      />

      <Input
        multiLine
        fullWidth
        rows={5}
        required
        name="content"
        label="內容文字*"
      />

      <div className="flex justify-end">
        <Button type="submit" variant="primary">
          提交
        </Button>
      </div>
    </form>
  );
}
export default Form;
