import { Button } from '@moai/core';
import styles from '../styles/Home.module.css';

export default function Home() {
  return (
    <div className={styles.container}>
      <div className={styles.content}>
        <div className={styles.form}>
          <Button highlight fill href="/list">Login</Button>
        </div>
      </div>
    </div>
  )
}
