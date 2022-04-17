import type { UserInfo } from 'src/types';
import type { TeamData } from 'src/types/teamInfo';
import { derived, writable } from 'svelte/store';

export const userInfo = writable<UserInfo>({} as UserInfo);

export const IDToken = writable<string>('');

export const TeamID = writable<string>('');

export const TeamInfoData = writable<TeamData>({} as TeamData);

export const brand = derived(TeamInfoData, ($TeamInfoData) => {
	return $TeamInfoData.team
		? $TeamInfoData.team.division === 'Advanced'
			? 'alt'
			: 'brand'
		: 'brand';
});

// export const JWT = writable<string>('');
