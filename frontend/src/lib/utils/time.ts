import type { TimeRemaining } from 'src/types';

const MS_IN_DAY = 86400000;
const MS_IN_HOUR = 3600000;
const MS_IN_MINUTE = 60000;
const MS_IN_SECOND = 1000;

export function getTime(startTime: number, stopTime: number, currentTime: number): TimeRemaining {
	let timeLeft = stopTime - currentTime;
	const totalTime = stopTime - startTime;
	if (timeLeft <= 0) {
		return { days: 0, hours: 0, minutes: 0, seconds: 0 };
	} else if (timeLeft > totalTime) {
		timeLeft = timeLeft - totalTime;
	}

	if (timeLeft <= 3 * MS_IN_DAY) {
		const hours = Math.floor(timeLeft / MS_IN_HOUR);
		const minutes = Math.floor((timeLeft % MS_IN_HOUR) / MS_IN_MINUTE);
		const seconds = Math.floor((timeLeft % MS_IN_MINUTE) / MS_IN_SECOND);

		return {
			days: 0,
			hours,
			minutes,
			seconds
		};
	} else {
		const days = Math.floor(timeLeft / MS_IN_DAY);
		const hours = Math.floor((timeLeft % MS_IN_DAY) / MS_IN_HOUR);
		const minutes = Math.floor((timeLeft % MS_IN_HOUR) / MS_IN_MINUTE);
		const seconds = Math.floor((timeLeft % MS_IN_MINUTE) / MS_IN_SECOND);

		return {
			days,
			hours,
			minutes,
			seconds
		};
	}
}

export function fmtLabel(startTime: number, stopTime: number, currentTime: number) {
	const timeLeft = stopTime - currentTime;
	const totalTime = stopTime - startTime;
	if (timeLeft < 0) {
		return 'COMPETITION HAS CONCLUDED';
	} else if (timeLeft < totalTime) {
		return 'TIME REMAINING';
	} else {
		return 'COMPETITION BEGINS IN';
	}
}

export default function fmtTime(startTime: number, stopTime: number, currentTime: number): string {
	const time = getTime(startTime, stopTime, currentTime);

	let str = '';
	if (time.days > 0) str += time.days + ' days ';
	if (time.days > 0 || time.hours > 0) str += time.hours + 'h ';
	if (time.days > 0 || time.hours > 0 || time.minutes > 0) str += time.minutes + 'm ';
	str += time.seconds + 's';

	return str;
}

// format seconds into hh:mm
export function fmtTimeHS(seconds: number): string {
	const absSeconds = Math.abs(seconds); // before contest is negative
	const hours = Math.floor(absSeconds / 3600);
	const minutes = Math.floor((absSeconds / 60) % 60);

	return `${hours < 10 ? '0' + hours : hours}:${minutes < 10 ? '0' + minutes : minutes}`;
}
