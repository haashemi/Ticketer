import { createQuery } from '@tanstack/svelte-query';
import { z } from 'zod';
import ky from 'ky';

const respSchema = z.object({
    name: z.string(),
    email: z.string().email(),
});

export const useProfile = () => createQuery({
    queryKey: ['profile'],
    queryFn: async () => {
        const resp = await ky.get('/api/profile').json();
        return respSchema.parse(resp);
    },
    retry: 0,
    refetchOnMount: true,
});