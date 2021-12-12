import { Icon, ProgressCircle, Table, TableColumn, Tooltip } from '@moai/core';
import { HiCheckCircle, HiXCircle } from 'react-icons/hi';
import styles from './table.module.css';
import { useCallback, useEffect, useMemo, useState } from 'react';
import { useAtom } from 'jotai';
import { date, fetcher } from 'utils';
import { ENDPOINTS, pages } from 'consts';
import {
  dataSizeStore,
  detailStore,
  regionStore,
  searchValueStore
} from 'store';
import { useRouter } from 'next/router';
import { SecretDetail } from 'type';

const tableColumns = (
  changeRouter: (arn: string) => void
): TableColumn<SecretDetail>[] => {
  return [
    {
      title: 'Name',
      className: styles.borderCell,
      render: (r) => (
        <a onClick={() => changeRouter(r.ARN)}>
          <span className={styles.link}>{r.Name}</span>
        </a>
      )
    },
    {
      title: 'Accessed at',
      className: styles.borderCell,
      render: (r) =>
        r.LastAccessedDate ? date.formatDateTime(r.LastAccessedDate) : '-'
    },
    {
      title: 'Created at',
      className: styles.borderCell,
      render: (r) => (r.CreatedDate ? date.formatDateTime(r.CreatedDate) : '-')
    },
    {
      title: 'Updated at',
      className: styles.borderCell,
      render: (r) =>
        r.LastChangedDate ? date.formatDateTime(r.LastChangedDate) : '-'
    },
    {
      title: 'Rotate',
      className: [styles.borderCell, 'text-center'].join(' '),
      render: (r) =>
        r.RotationEnabled ? (
          <Tooltip content="Enabled">
            <span className={styles.success}>
              <Icon component={HiCheckCircle} />
            </span>
          </Tooltip>
        ) : (
          <Tooltip content="Disabled">
            <span className={styles.failure}>
              <Icon component={HiXCircle} />
            </span>
          </Tooltip>
        )
    },
    {
      title: 'Rorated at',
      className: styles.borderCell,
      render: (r) =>
        r.LastRotatedDate ? date.formatDateTime(r.LastRotatedDate) : '-'
    },
    {
      title: 'Deleted at',
      render: (r) => (r.DeletedDate ? date.formatDateTime(r.DeletedDate) : '-')
    }
  ];
};

const DataTable = () => {
  const [loading, setLoading] = useState(false);
  const [data, setData] = useState<SecretDetail[]>([]);

  const [region] = useAtom(regionStore);
  const [searchValue] = useAtom(searchValueStore);
  const [, setDetail] = useAtom(detailStore);
  const [, setDataSize] = useAtom(dataSizeStore);

  const router = useRouter();

  const getSecretList = useCallback(
    async (region: string) => {
      setLoading(true);
      const [data, error] = await fetcher.post<SecretDetail[]>(
        ENDPOINTS.getSecretList,
        { region }
      );
      setLoading(false);
      if (error) {
        console.error(error);
        return;
      }
      if (data) {
        setData(data);
        setDataSize(data.length);
      }
    },
    [setDataSize]
  );

  const rows = useMemo(() => {
    if (!searchValue) return data;
    return data.filter((r) => r.Name.includes(searchValue));
  }, [data, searchValue]);

  useEffect(() => {
    if (region) getSecretList(region);
  }, [region, getSecretList]);

  const changeRouter = (arn: string) => {
    if (!region) return;
    setDetail({ arn, region });
    router.push(pages.secret.path);
  };

  return (
    <div className={styles.container}>
      {loading ? (
        <div className={styles.loading}>
          <ProgressCircle value="indeterminate" size={24} />
        </div>
      ) : data.length ? (
        <Table<SecretDetail>
          rows={rows}
          rowKey={(r) => r.ARN}
          columns={tableColumns(changeRouter)}
          fill
        />
      ) : (
        <div className={styles.loading}>No data</div>
      )}
    </div>
  );
};

export default DataTable;
