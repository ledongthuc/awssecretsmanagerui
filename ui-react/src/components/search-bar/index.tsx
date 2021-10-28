import { Button, ButtonGroup, Input, Select } from '@moai/core';
import { HiSearch } from 'react-icons/hi';
import styles from './search-bar.module.css';

const SearchBar = () => {
  return (
    <div className={styles.container}>
      <ButtonGroup>
        {[
          {
            fill: false,
            element: (
              <Select options={['eu-north-1'].map(Select.toStringOption)} />
            )
          },
          {
            fill: true,
            element: <Input placeholder="Search" icon={HiSearch} />
          },
          {
            fill: false,
            element: <Button color={Button.colors.highlight}>Search</Button>
          }
        ]}
      </ButtonGroup>
    </div>
  );
};

export default SearchBar;
