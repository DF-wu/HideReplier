import { z } from "zod";
import { AdvancedInfo } from "../types";

const API_IP = "https://httpbin.org/ip";
const API_COUNTRY = "https://ipwhois.app/json";

const ipResSchema = z.object({ origin: z.string() });
const ispResSchema = z.object({
  country: z.string().default(""),
  city: z.string().default(""),
  isp: z.string().default(""),
});

export async function getAdvancedInfo(): Promise<AdvancedInfo> {
  const result = await fetch(API_IP).then((res) => res.json());
  const parsedIp = ipResSchema.safeParse(result);
  if (!parsedIp.success) return Promise.reject("cannot retrive IP");
  const ip = parsedIp.data.origin.split(",")[0];
  const ispInfo = await fetch(`${API_COUNTRY}/${ip}`).then((res) => res.json());
  const parsedIspInfo = ispResSchema.parse(ispInfo);

  return AdvancedInfo.parse({
    ip,
    ...parsedIspInfo,
  });
}
