import { Button, DivPx, Icon, toast } from '@moai/core';
import { Fragment, useCallback, useState } from 'react';
import { HiUpload } from 'react-icons/hi';
import { SecretValue, SetState } from 'type';
import { DataType } from '../index';
import styles from './binary.module.css';
import { FileDownload } from '../download';
import { fetcher } from 'utils';
import { ENDPOINTS } from 'consts';
import { regionStore, useAtom } from 'store';

interface Props {
  arn: string;
  setType: SetState<DataType>;
  data: SecretValue;
  setData: SetState<SecretValue | null>;
}

export const UploadBinary = (props: Props) => {
  const { arn, setType, data, setData } = props;
  const [loading, setLoading] = useState(false);
  const [file, setFile] = useState<File | null>(null);
  const [region] = useAtom(regionStore);

  const uploadFile = useCallback(async () => {
    if (!file || !region) return;
    setLoading(true);

    const payload = new FormData();
    payload.append('file', file);
    const [data, error] = await fetcher.post<SecretValue>(
      ENDPOINTS.updateSecretBinary(arn, region),
      payload,
      {},
      { header: false, stringify: false }
    );

    setFile(null);
    setLoading(false);

    if (error) {
      toast(
        toast.types.failure,
        'Something went wrong, check the console for more information'
      );
      console.error(error);
    } else {
      toast(toast.types.success, 'Updated successful');
      if (data) setData(data);
    }
  }, [arn, file, region, setData]);

  return (
    <Fragment>
      <label className={styles.label}>
        <input
          className={styles.input}
          type="file"
          onChange={(e) => e.target.files && setFile(e.target.files[0])}
        />
        <Icon component={HiUpload} />
        <DivPx size={8} />
        <p>
          {file ? (
            <Fragment>
              {file.name} (
              {parseFloat((file.size / 1000).toString()).toFixed(2)}
              kb)
            </Fragment>
          ) : (
            <Fragment>
              Drop file here or <span className={styles.browse}>browse</span>
            </Fragment>
          )}
        </p>
      </label>
      <DivPx size={8} />
      <div className={styles.buttonGroup}>
        <Button
          color={Button.colors.highlight}
          busy={loading}
          disabled={file === null}
          onClick={() => uploadFile()}
        >
          Upload
        </Button>
        <DivPx size={16} />
        <span>or</span>
        <DivPx size={8} />
        <Button style={Button.styles.flat} onClick={() => setType('string')}>
          <span className="text-blue">Write a string</span>
        </Button>
      </div>
      <DivPx size={16} />
      <FileDownload name={data.Name} data={data.SecretBinary || ''} />
    </Fragment>
  );
};
