import { derived, readable } from 'svelte/store';
import { startTime } from './contestData';

export const currentTime = readable(Date.now(), function start(set) {
	const interval = setInterval(() => {
		set(Date.now());
	}, 1000);

	return function stop() {
		clearInterval(interval);
	};
});

export const contestStarted = derived([currentTime, startTime], ([$currentTime, $startTime]) => {
	return Math.floor($currentTime / 1000) >= $startTime.valueOf() / 1000;
});
