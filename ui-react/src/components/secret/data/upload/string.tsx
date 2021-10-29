import { Button, DivPx, toast } from '@moai/core';
import { ENDPOINTS } from 'consts';
import { Field, Form, Formik } from 'formik';
import { Fragment, useCallback, useState } from 'react';
import { regionStore, useAtom } from 'store';
import { SecretValue, SetState } from 'type';
import { fetcher } from 'utils';
import { DataType } from '..';
import styles from './string.module.css';

interface Props {
  arn: string;
  value: string;
  setType: SetState<DataType>;
  setData: SetState<SecretValue | null>;
}

export const UploadString = (props: Props) => {
  const { arn, value, setType, setData } = props;
  const [loading, setLoading] = useState(false);
  const [region] = useAtom(regionStore);

  const updateString = useCallback(
    async (value: string) => {
      if (!region) return;
      setLoading(true);
      const [data, error] = await fetcher.post<SecretValue>(
        ENDPOINTS.updateSecretValue(region),
        {
          SecretId: arn,
          SecretString: value
        }
      );
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
    },
    [arn, region, setData]
  );

  return (
    <Formik
      initialValues={{
        string: value
      }}
      onSubmit={(values) => updateString(values.string)}
    >
      {({ dirty, submitForm }) => (
        <Form>
          <Field as="textarea" name="string" className={styles.input} />
          <DivPx size={16} />
          <div className={styles.buttonGroup}>
            <Button
              color={Button.colors.highlight}
              busy={loading}
              onClick={submitForm}
              disabled={!dirty}
            >
              Save
            </Button>
            <DivPx size={16} />
            <span>or</span>
            <DivPx size={8} />
            <Button
              style={Button.styles.flat}
              onClick={() => setType('binary')}
            >
              <span className="text-blue">Upload a file</span>
            </Button>
          </div>
        </Form>
      )}
    </Formik>
  );
};
