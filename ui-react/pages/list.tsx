import Layout from "../components/Layout"
import styles from "../styles/List.module.css";

export default function Home() {
  return (
    <Layout title="AWS Secrets Manager">
      <div className={styles.container}>
        Secret List
      </div>
    </Layout>
  )
}
