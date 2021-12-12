import { DataTable, Pagination, SearchBar, Layout } from 'components';
import { DivPx } from '@moai/core';
import styles from 'styles/home.module.css';

export default function Home() {
  return (
    <Layout>
      <div className={styles.container}>
        <SearchBar />
        <DivPx size={24} />
        <DataTable />
        <DivPx size={24} />
        <Pagination />
      </div>
    </Layout>
  );
}
