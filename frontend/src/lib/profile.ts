import { createQuery } from "@tanstack/svelte-query";
import ky from "ky";
import { z } from "zod";

const respSchema = z.object({
  name: z.string(),
  email: z.string().email(),
});

export const useProfile = () =>
  createQuery({
    queryKey: ["profile"],
    queryFn: async () => {
      const resp = await ky.get("/api/profile").json();
      return respSchema.parse(resp);
    },
    retry: 0,
    refetchOnMount: true,
  });
