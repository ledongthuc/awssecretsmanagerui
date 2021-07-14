import React from 'react';
import Head from 'next/head';
import Header from './Header';
import styles from './Layout.module.css';

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