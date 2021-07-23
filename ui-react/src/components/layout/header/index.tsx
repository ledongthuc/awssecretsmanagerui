import { Button } from '@moai/core';
import Link from 'next/link';
import Image from 'next/image';
import styles from './header.module.css';

interface HeaderProps {
}

const Brand = () => {
  return (
    <Link href="/" passHref >
      <a className={styles.brand}>
        <Image alt="Logo"
          src="/logo.svg"
          width={36}
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
