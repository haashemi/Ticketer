import moment from 'moment';

export const msToTime = (s: number): string => moment(new Date(s)).format('HH:mm');
