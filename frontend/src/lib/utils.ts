import moment from 'moment';

export function msToTime(s: number): string {
	const d = new Date(s / 1000);
	return moment(d).format('HH:mm');
}
