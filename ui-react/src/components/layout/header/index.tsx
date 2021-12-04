import Link from 'next/link';
import styles from './header.module.css';

const Brand = () => {
  return (
    <Link href="/" passHref>
      <a className={styles.brand}>
        <img alt="Logo" src="/logo.svg" width={36} height={36} />
        <h3>Secrets Manager</h3>
      </a>
    </Link>
  );
};

const Header = () => {
  return (
    <header className={styles.container}>
      <div className={styles.content}>
        <Brand />
        <button className={styles.logout}>Log out</button>
      </div>
    </header>
  );
};

export default Header;
