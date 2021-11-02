import { DivPx, Icon, ProgressCircle, toast } from '@moai/core';
import { DataTab, InfoTab, Layout } from 'components';
import { ENDPOINTS } from 'consts';
import { useAtom } from 'jotai';
import { Fragment, useCallback, useEffect, useState } from 'react';
import { detailStore } from 'store';
import styles from 'styles/secret.module.css';
import { SecretDetail } from 'type';
import { fetcher } from 'utils';
import Link from 'next/link';
import { HiChevronLeft } from 'react-icons/hi';

type TabType = 'info' | 'data';

export default function SecretPage() {
  const [tab, setTab] = useState<TabType>('info');
  const [loading, setLoading] = useState(false);
  const [data, setData] = useState<SecretDetail | null>(null);
  const [detail] = useAtom(detailStore);

  const getDetail = useCallback(async (arn: string, region: string) => {
    setLoading(true);
    const [data, error] = await fetcher.post<SecretDetail>(
      ENDPOINTS.getSecretDetail,
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
    setData(data);
  }, []);

  useEffect(() => {
    if (detail?.arn && detail.region) getDetail(detail.arn, detail.region);
  }, [detail, getDetail]);

  return (
    <Layout>
      <div className={styles.container}>
        {loading ? (
          <div className={styles.notFound}>
            <ProgressCircle value="indeterminate" size={24} />
          </div>
        ) : data === null ? (
          <div className={styles.notFound}>Not Found</div>
        ) : (
          <Fragment>
            <div className={styles.header}>
              <Link href="/">
                <a className={styles.back}>
                  <Icon component={HiChevronLeft} />
                  <DivPx size={4} />
                  <span>Back</span>
                </a>
              </Link>
              <div className={styles.tabs}>
                {(['info', 'data'] as TabType[]).map((_tab) => (
                  <span
                    key={_tab}
                    className={[
                      styles.tab,
                      tab === _tab ? styles.tabActive : ''
                    ].join(' ')}
                    onClick={() => setTab(_tab)}
                  >
                    {_tab.charAt(0).toUpperCase() + _tab.slice(1)}
                  </span>
                ))}
              </div>
              <span className={styles.file}>
                <strong>Secrets</strong> {data.Name}
              </span>
            </div>
            <div className={styles.entry}>
              {tab === 'info' ? <InfoTab data={data} /> : <DataTab />}
            </div>
          </Fragment>
        )}
      </div>
    </Layout>
  );
}
