import { derived, readable } from 'svelte/store';
import { startTime, stopTime } from './contestData';

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

export const contestEnded = derived([currentTime, stopTime], ([$currentTime, $stopTime]) => {
	return Math.floor($currentTime / 1000) >= $stopTime.valueOf() / 1000;
});
