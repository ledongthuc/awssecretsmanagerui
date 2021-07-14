import React from 'react';
import Link from 'next/link';
import styles from './Header.module.css';

interface HeaderProps {
}

const Logo = () => {
  return <div>Logo</div>
}

const Header = (props: HeaderProps) => {
  return (
    <header>
      <div className={styles.headerContent}>
        <Link href="/" passHref>
          <a>
            <Logo />
          </a>
        </Link>
        <div>Action</div>
      </div>
    </header>
  )
}

export default Header;