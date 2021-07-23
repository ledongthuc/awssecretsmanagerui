import Head from 'next/head';
import Header from './header';
import styles from './layout.module.css';

interface LayoutProps {
  children: React.ReactNode;
  title: string;
}

const Layout = (props: LayoutProps) => {
  return (
    <div>
      <Head>
        <title>{props.title}</title>
        <link rel="icon" href="/favicon.ico" />
      </Head>
      <Header />
      <main className={styles.main}>{props.children}</main>
    </div>
  )
}

export default Layout;