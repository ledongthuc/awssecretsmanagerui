import { dataSizeStore, useAtom } from 'store';
import styles from './pagination.module.css';

const Pagination = () => {
  const [dataSize] = useAtom(dataSizeStore);

  return (
    <div className={styles.container}>
      <span className={styles.text}>{dataSize} item(s) total</span>
      {/* <div className={styles.pageControl}>
        <span className={styles.text}>Page</span>
        <DivPx size={8} />
        <Select options={['1', '2', '3'].map(Select.toStringOption)} />
        <DivPx size={8} />
        <span className={styles.text}>|</span>
        <DivPx size={8} />
        <Button icon={HiArrowLeft} iconLabel="Previous" />
        <DivPx size={8} />
        <Button icon={HiArrowRight} iconLabel="Next" />
      </div> */}
    </div>
  );
};

export default Pagination;
