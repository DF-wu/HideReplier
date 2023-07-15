import { useEffect, useRef } from "react";
import { Content, PartialContent } from "../types";
import { Button } from "./Button";
import { Input } from "./Input";
import { str2rgb } from "../utils/string2rgb";

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

  useEffect(() => {
    onFormDataChanged();
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
      <Input fullWidth name="avatar_url" type="url" label="頭像連結" />
      <Input fullWidth name="thumbnail" type="url" label="小圖連結" />
      <Input fullWidth name="imgUrl" type="url" label="圖片連結" />
      <Input
        fullWidth
        name="color"
        type="color"
        label="顏色（裝飾）"
        defaultValue={str2rgb(props.ip)}
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
