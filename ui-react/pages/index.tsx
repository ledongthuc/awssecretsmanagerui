import { Button } from '@moai/core';
import styles from '../styles/Home.module.css';

export default function Home() {
  return (
    <div className={styles.container}>
      <div className={styles.content} style={{ background: "yellow" }}>
        <Button highlight style={Button.styles.flat}>Log In</Button>
      </div>
    </div>
  )
}
