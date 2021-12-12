import { ProgressCircle, toast } from '@moai/core';
import { ENDPOINTS } from 'consts';
import { useAtom } from 'jotai';
import { Fragment, useCallback, useEffect, useState } from 'react';
import { detailStore } from 'store';
import { SecretValue } from 'type';
import { fetcher } from 'utils';
import styles from './data-tab.module.css';
import { UploadBinary } from './upload/binary';
import { UploadString } from './upload/string';

export type DataType = 'string' | 'binary';

const DataTab = (): JSX.Element => {
  const [type, setType] = useState<DataType>('string');
  const [loading, setLoading] = useState(false);
  const [data, setData] = useState<SecretValue | null>(null);
  const [detail] = useAtom(detailStore);

  const getData = useCallback(async (arn: string, region: string) => {
    setLoading(true);
    const [data, error] = await fetcher.post<SecretValue>(
      ENDPOINTS.getSecretValue,
      {
        arn,
        region
      }
    );
    setLoading(false);
    if (error) {
      toast(toast.types.failure, 'Something went wrong');
      console.error(error);
      return;
    }
    if (data) {
      setData(data);
      if (data.SecretString) setType('string');
      if (data.SecretBinary) setType('binary');
    }
  }, []);

  useEffect(() => {
    if (detail?.arn && detail.region) getData(detail.arn, detail.region);
  }, [detail, getData]);

  if (!data) return <Fragment />;
  if (loading)
    return (
      <div className={styles.loading}>
        <ProgressCircle value="indeterminate" size={24} />
      </div>
    );

  return type === 'string' ? (
    <UploadString
      arn={data.ARN}
      value={data.SecretString || ''}
      setType={setType}
      setData={setData}
    />
  ) : (
    <UploadBinary
      arn={data.ARN}
      data={data}
      setType={setType}
      setData={setData}
    />
  );
};

export { DataTab };
