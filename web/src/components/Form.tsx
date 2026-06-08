import { useCallback, useEffect, useRef, useState } from "react";
import { Content, PartialContent } from "../types";
import { Button } from "./Button";
import { Input } from "./Input";
import { str2rgb } from "../utils/string2rgb";
import { ImageInput } from "./ImageInput";
import { DiscordTarget } from "../services/api";

export type FormProps = {
  ip: string;
  discordTargets?: DiscordTarget[];
  onSubmit?: (content: Content) => void;
  onFormDataChanged?: (content: PartialContent) => void;
};
export function Form(props: FormProps) {
  const {
    ip,
    discordTargets = [],
    onFormDataChanged: emitFormDataChanged,
    onSubmit,
  } = props;
  const formRef = useRef<HTMLFormElement>(null);
  const getFormData = useCallback(() => {
    if (!formRef.current) return;
    return {
      ...Object.fromEntries(new FormData(formRef.current)),
      ip,
    };
  }, [ip]);

  const onFormDataChanged = useCallback(() => {
    const partialContent = PartialContent.safeParse(getFormData());
    if (!partialContent.success) return console.error(partialContent.error);
    emitFormDataChanged?.(partialContent.data);
  }, [emitFormDataChanged, getFormData]);

  const [color, setColor] = useState<string | undefined>(
    ip ? str2rgb(ip) : ""
  );
  useEffect(() => {
    if (ip && !color) setColor(str2rgb(ip));
    const timeout = window.setTimeout(() => onFormDataChanged(), 0);
    return () => window.clearTimeout(timeout);
  }, [color, ip, onFormDataChanged]);

  return (
    <form
      ref={formRef}
      onSubmit={(e) => {
        e.preventDefault();
        if (!formRef.current) return;
        const content = Content.safeParse(getFormData());
        if (!content.success) return console.error(content.error);
        onSubmit?.(content.data);
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
      {discordTargets.length > 1 && (
        <label className="flex flex-col gap-1 text-left text-zinc-100">
          <span className="text-[0.9rem] font-medium">發送頻道</span>
          <select
            name="targetId"
            defaultValue={
              discordTargets.find((target) => target.default)?.id ||
              discordTargets[0]?.id
            }
            onChange={onFormDataChanged}
            className="w-full rounded border border-solid border-zinc-600 bg-zinc-700 px-2 py-1 text-zinc-100 hover:border-zinc-400 focus:border-zinc-400 focus:outline-none"
          >
            {discordTargets.map((target) => (
              <option key={target.id} value={target.id}>
                {target.label}
              </option>
            ))}
          </select>
        </label>
      )}
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
