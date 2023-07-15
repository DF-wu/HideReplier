import { z } from "zod";
import { AdvancedInfo, Content } from "../types";

const IP_API = "https://ipapi.co/json";
const BACKEND_API =
  (import.meta.env.DEV ? "http://localhost:8080" : window.location.origin) +
  "/HideBot/discord";
const DEFAULT_AVATAR =
  "https://cdn.discordapp.com/avatars/710112845567623238/f377b595ef4e0ea17826d7afbb20633f.webp";

const ispResSchema = z.object({
  ip: z.string().default(""),
  country_name: z.string().default(""),
  city: z.string().default(""),
  org: z.string().default(""),
  timezone: z.string().default(""),
});

export async function getAdvancedInfo(): Promise<AdvancedInfo> {
  const result = await fetch(IP_API).then((res) => res.json());
  const parsed = ispResSchema.safeParse(result);
  if (!parsed.success) return Promise.reject("cannot retrive IP information");
  return AdvancedInfo.parse({
    country: parsed.data.country_name,
    isp: parsed.data.org,
    ...parsed.data,
  });
}

export async function sendPost(post: Content) {
  if (!post.avatar_url) post.avatar_url = DEFAULT_AVATAR;
  await fetch(BACKEND_API, {
    body: JSON.stringify(post),
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
  });
}
