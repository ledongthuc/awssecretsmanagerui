import { atom } from 'jotai';
import { DetailPayload } from 'type';

export const detailStore = atom<DetailPayload | null>(null);
