import { Button, DivPx, Icon, Table } from '@moai/core';
import { CodeInput, Layout } from 'components';
import { useState } from 'react';
import { HiCheckCircle } from 'react-icons/hi';
import styles from 'styles/secret.module.css';

type TabType = 'info' | 'data';

const InfoTab = () => {
  return (
    <div>
      <label className={styles.label}>ARN</label>
      <p>
        arn:aws:secretsmanager:eu-north-1:554047623271:secret:dev/secret_variables/crypto
      </p>

      <DivPx size={24} />

      <label className={styles.label}>Description</label>
      <p>
        It was easy to spot her. All you needed to do was look at her socks.
        They were never a matching pair. One would be green while the other
        would be blue. One would reach her knee while the other barely touched
        her ankle. Every other part of her was perfect, but never the socks.
        They were her micro act of rebellion.
      </p>

      <DivPx size={24} />

      <label className={styles.label}>Owning service</label>
      <div style={{ display: 'flex', flexWrap: 'wrap' }}>
        <div style={{ width: '50%', marginBottom: 8 }}>
          <label className={styles.subLabel}>Accessed at</label>
          <p>2021/06/29 00:00:00</p>
        </div>
        <div style={{ width: '50%', marginBottom: 8 }}>
          <label className={styles.subLabel}>Changed at</label>
          <p>2021/06/29 00:00:00</p>
        </div>
        <div style={{ width: '50%', marginBottom: 8 }}>
          <label className={styles.subLabel}>
            <span>Rotation enabled:</span>
            <DivPx size={4} />
            <span className={styles.success}>
              <Icon component={HiCheckCircle} />
            </span>
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
        <div className={styles.tableRow}>
          <div className={styles.tableCell}>ENVIROMENT</div>
          <div className={styles.tableCell}>prod</div>
        </div>
      </div>
    </div>
  );
};

const DataTab = () => {
  const [value, setValue] = useState(
    '{ "password": "fc151f2a84772ca62bc0e07dc96f5ee9", "username": "Thức Thối Tha" }'
  );

  return (
    <div>
      <CodeInput value={value} setValue={setValue} />
      <DivPx size={16} />
      <div className={styles.flex}>
        <Button color={Button.colors.highlight}>Save</Button>
        <DivPx size={16} />
        <span className={styles.textGray}>or</span>
        <DivPx size={8} />
        <Button style={Button.styles.flat}><span className={styles.textBlue}>Upload a file</span></Button>
      </div>
    </div>
  );
};

export default function SecretPage() {
  const [tab, setTab] = useState<TabType>('info');

  return (
    <Layout>
      <div className={styles.container}>
        <div className={styles.header}>
          <div className={styles.tabs}>
            {(['info', 'data'] as TabType[]).map((_tab) => (
              <span
                className={[
                  styles.tab,
                  tab === _tab ? styles.tabActive : ''
                ].join(' ')}
                onClick={() => setTab(_tab)}
              >
                {_tab.charAt(0).toUpperCase() + _tab.slice(1)}
              </span>
            ))}
          </div>
          <span className={styles.file}>
            <strong>Secrets</strong> /dev/secret_variables/crypto
          </span>
        </div>
        <div className={styles.entry}>
          {tab === 'info' ? <InfoTab /> : <DataTab />}
        </div>
      </div>
    </Layout>
  );
}
