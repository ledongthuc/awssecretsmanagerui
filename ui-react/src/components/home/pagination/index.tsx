import { Button, DivPx, Select } from '@moai/core';
import { HiArrowLeft, HiArrowRight } from 'react-icons/hi';
import styles from './pagination.module.css';

const PAGE_SIZE = 10;

const Pagination = () => {
  return (
    <div className={styles.container}>
      <span className={styles.text}>1 - {PAGE_SIZE} of 38</span>
      <div className={styles.pageControl}>
        <span className={styles.text}>Page</span>
        <DivPx size={8} />
        <Select options={['1', '2', '3'].map(Select.toStringOption)} />
        <DivPx size={8} />
        <span className={styles.text}>|</span>
        <DivPx size={8} />
        <Button icon={HiArrowLeft} iconLabel="Previous" />
        <DivPx size={8} />
        <Button icon={HiArrowRight} iconLabel="Next" />
      </div>
    </div>
  );
};

export default Pagination;
