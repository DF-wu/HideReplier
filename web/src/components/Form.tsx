import { useRef } from "react";
import { Content } from "../types";
import { Button } from "./Button";
import { Input } from "./Input";

export type FormProps = {
  ip: string;
  onSubmit: (content: Content) => void;
};
export function Form(props: FormProps) {
  const formRef = useRef<HTMLFormElement>(null);
  return (
    <form
      ref={formRef}
      onSubmit={(e) => {
        e.preventDefault();
        if (!formRef.current) return;
        const content = Content.safeParse({
          ...Object.fromEntries(new FormData(formRef.current)),
          ip: props.ip,
        });
        if (!content.success) return console.error(content.error);
        console.log(content.data);
        props.onSubmit(content.data);
      }}
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
      {/* TODO: 小圖 */}
      <Input fullWidth name="thumbnail" type="url" label="小圖" />
      <Input fullWidth name="imgUrl" type="url" label="圖片連結" />
      <Input fullWidth name="color" type="color" label="顏色（裝飾）" />

      <Input
        multiLine
        fullWidth
        rows={5}
        required
        name="content"
        label="內容文字*"
      />

      <div className="flex justify-between">
        <Button variant="success">預覽</Button>
        <Button type="submit" variant="primary">
          提交
        </Button>
      </div>
    </form>
  );
}
