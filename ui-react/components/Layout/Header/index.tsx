import { Button } from '@moai/core';
import React from 'react';
import Link from 'next/link';
import Image from 'next/image';
import styles from './Header.module.css';

interface HeaderProps {
}

const Brand = () => {
  return (
    <Link href="/" passHref >
      <a className={styles.brand}>
        <Image alt="Logo"
          src="/logo.png"
          width={31.37}
          height={36}
        />
        <h3>Secrets Manager</h3>
      </a>
    </Link>
  )
}

const Header = (props: HeaderProps) => {
  return (
    <header className={styles.container}>
      <div className={styles.content}>
        <Brand />
        <Button style={Button.styles.flat}>Log out</Button >
      </div>
    </header>
  )
}

export default Header;