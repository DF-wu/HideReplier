import { useRef } from "react";
import { Content } from "../types";
import { Button } from "./Button";
import { Input } from "./Input";

export type FormProps = {
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
          ip: "1.1.1.1", // TODO: get ip using react query
        });
        if (!content.success) return console.error(content.error);
        props.onSubmit(content.data);
      }}
      className="flex flex-col gap-3"
    >
      <Input
        name="username"
        required
        label="機器人名稱"
        defaultValue="預設機器人:)"
      />
      <Input name="avatar_url" type="url" label="頭像連結" />
      {/* TODO: 小圖 */}
      <Input name="thumbnail" type="url" label="小圖" />
      <Input name="imgUrl" type="url" label="圖片連結" />
      <Input name="color" type="color" label="顏色（裝飾）" />

      <Input multiLine rows={5} required name="content" label="內容文字" />

      <div className="flex justify-between">
        <Button variant="success">預覽</Button>
        <Button type="submit" variant="primary">
          提交
        </Button>
      </div>
    </form>
  );
}
