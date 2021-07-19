import Layout from "../components/Layout"
import SearchBar from "../components/SearchBar";
import Table from "../components/Table";
import Pagination from "../components/Pagination";
import styles from "../styles/List.module.css";

export default function Home() {
  return (
    <Layout title="AWS Secrets Manager">
      <div className={styles.container}>
        <SearchBar />
        <Table />
        <Pagination />
      </div>
    </Layout>
  )
}
