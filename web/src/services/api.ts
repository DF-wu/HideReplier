import { z } from "zod";
import { AdvancedInfo, Content } from "../types";

const IP_API = "https://ipapi.co/json";
const BACKEND_API =
  (import.meta.env.DEV ? "http://localhost:8082" : window.location.origin) +
  "/HideBot/discord";
const TARGETS_API = `${BACKEND_API}/targets`;
const DEFAULT_AVATAR =
  "https://cdn.discordapp.com/avatars/710112845567623238/f377b595ef4e0ea17826d7afbb20633f.webp";

export type DiscordTarget = {
  id: string;
  label: string;
  default?: boolean;
};

const discordTargetSchema = z.object({
  id: z.string(),
  label: z.string(),
  default: z.boolean().optional(),
});

const ispResSchema = z.object({
  ip: z.string().default(""),
  country_name: z.string().default(""),
  city: z.string().default(""),
  org: z.string().default(""),
  timezone: z.string().default(""),
});

export async function getAdvancedInfo(): Promise<AdvancedInfo> {
  const response = await fetch(IP_API);
  const result = (await response.json()) as unknown;
  const parsed = ispResSchema.safeParse(result);
  if (!parsed.success) return Promise.reject("cannot retrive IP information");
  return AdvancedInfo.parse({
    country: parsed.data.country_name,
    isp: parsed.data.org,
    ...parsed.data,
  });
}

export async function sendPost(post: Content) {
  const payload = {
    ...post,
    avatar_url: post.avatar_url || DEFAULT_AVATAR,
  };

  const response = await fetch(BACKEND_API, {
    body: JSON.stringify(payload),
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
  });

  if (!response.ok) {
    throw new Error(`post failed with status ${response.status}`);
  }
}

export async function getDiscordTargets(): Promise<DiscordTarget[]> {
  const response = await fetch(TARGETS_API);
  if (!response.ok) {
    throw new Error(`targets failed with status ${response.status}`);
  }

  const result = (await response.json()) as unknown;
  return z.array(discordTargetSchema).parse(result);
}
