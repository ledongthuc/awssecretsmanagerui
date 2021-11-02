import { Button, ButtonGroup, Input, Select } from '@moai/core';
import { fetcher } from 'utils';
import { ENDPOINTS } from 'consts';
import { useEffect, useState } from 'react';
import { HiSearch } from 'react-icons/hi';
import styles from './search-bar.module.css';
import {
  useAtom,
  regionStore,
  useRegionStorage,
  searchValueStore
} from 'store';

type ServerList = { [id: string]: string };

const getSevers = async () => {
  const [data, error] = await fetcher.get<ServerList>(ENDPOINTS.getRegions);
  if (error) {
    console.error(error);
    return null;
  }
  return data;
};

const SearchBar = () => {
  const [loading, setLoading] = useState(false);

  const [regions, setRegions] = useState<string[]>([]);
  const [regionStorage, setRegionStorage] = useRegionStorage();
  const [region, setRegion] = useAtom(regionStore);

  const [search, setSearch] = useState('');
  const [, setSearchValue] = useAtom(searchValueStore);

  useEffect(() => {
    setLoading(true);
    getSevers().then((res) => {
      if (res) {
        const data = Object.values(res).sort();
        setRegions(data);
        if (!regionStorage) {
          setRegionStorage(data[0]);
          setRegion(data[0]);
        } else {
          setRegion(regionStorage);
        }
      }
      setLoading(false);
    });
  }, [regionStorage, setRegionStorage, setRegion]);

  useEffect(() => setSearchValue(''), [setSearchValue]);

  const handleSearch = () => setSearchValue(search);
  const handleKeyPress = (e: React.KeyboardEvent<HTMLInputElement>) => {
    if (e.key === 'Enter') handleSearch();
  };

  return (
    <div className={styles.container}>
      <ButtonGroup>
        {[
          {
            fill: false,
            element: (
              <Select<string>
                options={regions.map(Select.toStringOption)}
                disabled={loading}
                defaultValue={region || undefined}
                value={region || undefined}
                setValue={(value) => {
                  setRegion(value);
                  setRegionStorage(value);
                }}
              />
            )
          },
          {
            fill: true,
            element: (
              <Input
                placeholder="Search"
                icon={HiSearch}
                value={search}
                setValue={setSearch}
                onKeyPress={handleKeyPress}
              />
            )
          },
          {
            fill: false,
            element: (
              <Button
                color={Button.colors.highlight}
                onClick={() => handleSearch()}
              >
                Search
              </Button>
            )
          }
        ]}
      </ButtonGroup>
    </div>
  );
};

export default SearchBar;
