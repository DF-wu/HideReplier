import { z } from "zod";
import { AdvancedInfo } from "../types";

const IP_API = "https://ipapi.co/json";

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
