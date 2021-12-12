import { HiCheckCircle, HiXCircle } from 'react-icons/hi';
import { DivPx, Icon, Tooltip } from '@moai/core';
import styles from './info-tab.module.css';
import { SecretDetail } from 'type';
import { Fragment } from 'react';
import { date } from 'utils';

interface Props {
  data: SecretDetail;
}

const InfoTab = (props: Props): JSX.Element => {
  const { data } = props;

  return (
    <div>
      <label className={styles.label}>ARN</label>
      <p>{data.ARN}</p>
      <DivPx size={24} />
      {data.Description && (
        <Fragment>
          <label className={styles.label}>Description</label>
          <p>{data.Description}</p>
        </Fragment>
      )}
      <DivPx size={24} />
      <label className={styles.label}>Owning service</label>
      <div className={styles.labelGroup}>
        <div className={styles.labelItem}>
          <label className={styles.subLabel}>Accessed at</label>
          <p>
            {data.LastAccessedDate
              ? date.formatDateTime(data.LastAccessedDate)
              : '-'}
          </p>
        </div>
        <div className={styles.labelItem}>
          <label className={styles.subLabel}>Changed at</label>
          <p>
            {data.LastChangedDate
              ? date.formatDateTime(data.LastChangedDate)
              : '-'}
          </p>
        </div>
        <div className={styles.labelItem}>
          <label className={styles.subLabel}>
            <span>Rotation enabled:</span>
            <DivPx size={4} />
            {data.RotationEnabled ? (
              <Tooltip content="Enabled">
                <span className={styles.success}>
                  <Icon component={HiCheckCircle} />
                </span>
              </Tooltip>
            ) : (
              <Tooltip content="Disabled">
                <span className={styles.failure}>
                  <Icon component={HiXCircle} />
                </span>
              </Tooltip>
            )}
          </label>
        </div>
      </div>

      <DivPx size={24} />

      <label className={styles.label}>Tags</label>
      <div className={styles.table}>
        <div className={[styles.tableRow, styles.tableHeader].join(' ')}>
          <div className={styles.tableCell}>Key</div>
          <div className={styles.tableCell}>Value</div>
        </div>
        {data.Tags.map((tag) => (
          <div className={styles.tableRow} key={tag.Key}>
            <div className={styles.tableCell}>
              <strong>{tag.Key}</strong>
            </div>
            <div className={styles.tableCell}>{tag.Value}</div>
          </div>
        ))}
      </div>
    </div>
  );
};

export { InfoTab };
