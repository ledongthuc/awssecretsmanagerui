import { Icon, Table, TableColumn } from '@moai/core';
import { HiCheckCircle, HiXCircle } from 'react-icons/hi';
import Link from 'next/link';
import styles from './table.module.css';

interface TableData {
  id: string;
  name: string;
  accessed_at?: string;
  created_at: string;
  updated_at?: string;
  rotate: boolean;
  rotated_at?: string;
  deleted_at?: string;
}

const tableColumns = (): TableColumn<TableData>[] => {
  return [
    {
      title: 'Name',
      render: (r) => (
        <Link href="/secret">
          <a>{r.name}</a>
        </Link>
      )
    },
    { title: 'Accessed at', render: (r) => r.accessed_at },
    { title: 'Created at', render: (r) => r.created_at },
    { title: 'Updated at', render: (r) => r.updated_at },
    {
      title: 'Rotate',
      render: (r) =>
        r.rotated_at ? (
          <span className={styles.success}>
            <Icon component={HiCheckCircle} />
          </span>
        ) : (
          <span className={styles.failure}>
            <Icon component={HiXCircle} />
          </span>
        )
    },
    { title: 'Rorated at', render: (r) => r.rotated_at },
    { title: 'Deleted at', render: (r) => r.deleted_at }
  ];
};

const DataTable = () => {
  return (
    <div className={styles.container}>
      <Table<TableData>
        rows={[
          {
            id: '1',
            name: '/dev/secrets/capacity_binary',
            accessed_at: '2020-01-01T00:00:00Z',
            created_at: '2020-01-01T00:00:00Z',
            updated_at: '2020-01-01T00:00:00Z',
            rotate: true,
            rotated_at: '2020-01-01T00:00:00Z',
            deleted_at: undefined
          }
        ]}
        rowKey={(r) => r.id.toString()}
        columns={tableColumns()}
        fill
      />
    </div>
  );
};

export default DataTable;
