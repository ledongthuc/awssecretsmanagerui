import { DivPx, Icon } from '@moai/core';
import { HiDocument, HiDownload } from 'react-icons/hi';
import styles from './file-download.module.css';

interface Props {
  name: string;
  data: string;
}

const getFileSize = (data: string) => {
  const sizeInBytes = 4 * Math.ceil(data.length / 3) * 0.5624896334383812;
  const sizeInKb = sizeInBytes / 1000;
  return Math.round(sizeInKb * 100) / 100;
};

const getFileName = (name: string) => {
  if (/\//.test(name)) {
    const fileName = name.split('/');
    return fileName[fileName.length - 1];
  }
  return name;
};

export const FileDownload = (props: Props) => {
  return (
    <a
      className={styles.container}
      href={`data:application/octet-stream;base64,${props.data}`}
      download={getFileName(props.name)}
    >
      <div className={styles.icon}>
        <Icon component={HiDocument} size={24} />
      </div>
      <DivPx size={10} />
      <div className={styles.text}>
        <p className={styles.fileName}>{getFileName(props.name)}</p>
        <span className={styles.fileSize}>
          {getFileSize(props.data || '')}kb
        </span>
      </div>
      <DivPx size={48} />
      <div className={styles.download}>
        <Icon component={HiDownload} size={24} />
      </div>
    </a>
  );
};
