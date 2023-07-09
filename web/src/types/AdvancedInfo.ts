import { z } from "zod";

export const AdvancedInfo = z.object({
  ip: z.string().default(""),
  isp: z.string().default(""),
  city: z.string().default(""),
  country: z.string().default(""),
});

export type AdvancedInfo = z.infer<typeof AdvancedInfo>;
