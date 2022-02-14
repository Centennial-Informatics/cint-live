import type { UserInfo } from 'src/types';
import { writable } from 'svelte/store';

export const userInfo = writable<UserInfo>({} as UserInfo);

export const IDToken = writable<string>('');

// export const JWT = writable<string>('');
