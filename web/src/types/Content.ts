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

export const PartialContent = z.object({
  avatar_url: z.string().optional(),
  color: z.string().optional(),
  content: z.string().optional(),
  imgUrl: z.string().optional(),
  ip: z.string().optional(),
  username: z.string().optional(),
  thumbnail: z.string().optional(),
});

export type PartialContent = z.infer<typeof PartialContent>;
