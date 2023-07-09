import { z } from "zod";

export const Content = z.object({
  avatar_url: z.string(),
  color: z.string(),
  content: z.string(),
  imgUrl: z.string(),
  ip: z.string(),
  username: z.string(),
  thumbnail: z.string(),
});

export type Content = z.infer<typeof Content>;
