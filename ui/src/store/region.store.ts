import { useStorage } from 'hooks';
import { atom } from 'jotai';
import { SetState } from 'type';

export const regionStore = atom<string | null>(null);
export const useRegionStorage = (): [string, SetState<string>] => {
  const [region, setRegion] = useStorage('region', '');
  return [region, setRegion];
};
