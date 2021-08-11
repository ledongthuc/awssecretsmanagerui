import { config } from 'constants/config.const';
import { useRouter } from 'next/router';
import Head from 'next/head';
import { useEffect, useState } from 'react';
import Header from './header';
import styles from './layout.module.css';
import { pages } from 'constants/sites.const';
import { ProgressCircle } from '@moai/core';

interface Props {
  title?: string;
}

const Layout: React.FC<Props> = (props) => {
  const [isAuthorized, setAuthorization] = useState<null | boolean>(null);
  const router = useRouter();

  useEffect(() => {
    setAuthorization(true);
    if (isAuthorized === false) router.push(pages.login.path);
  }, [isAuthorized, router]);

  if (isAuthorized === null) {
    return (
      <div className={styles.container}>
        <ProgressCircle value="indeterminate" size={64} />
      </div>
    );
  }

  return (
    <div>
      <Head>
        <title>{props.title || config.title}</title>
      </Head>
      <Header />
      <main className={styles.main}>{props.children}</main>
    </div>
  );
};

export default Layout;
