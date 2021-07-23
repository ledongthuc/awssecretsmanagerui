import { DivPx, Input, Button } from '@moai/core';
import { Formik, Field, Form } from 'formik';
import styles from 'styles/login.module.css';
import Image from 'next/image';
import Head from 'next/head';
import { config } from 'constants/config.const';
import { pages } from 'constants/sites.const';

const LoginPage = () => {
  return (
    <div className={styles.container}>
      <Head>
        <title>{pages.login.title}</title>
      </Head>
      <div className={styles.entry}>
        <Image src="/logo.svg" alt="Logo" width={64} height={64} />
        <DivPx size={48} />
        <h2 className={styles.heading}>Welcome back</h2>
        <span className={styles.description}>Just secure everything!</span>
        <DivPx size={48} />
        <Formik
          initialValues={{ username: '', password: '' }}
          onSubmit={async (values, { setSubmitting }) => {
            console.log(JSON.stringify(values, undefined, 2));
            setSubmitting(false);
          }}
        >
          {({ isSubmitting: busy }) => (
            <Form>
              <div className={styles.formGroup}>
                <label htmlFor="username" className={styles.formLabel}>
                  Username
                </label>
                <Field
                  type="text"
                  name="username"
                  id="username"
                  size={Input.sizes.large}
                  as={Input}
                />
              </div>
              <div className={styles.formGroup}>
                <label htmlFor="password" className={styles.formLabel}>
                  Password
                </label>
                <Field
                  type="password"
                  name="password"
                  id="password"
                  size={Input.sizes.large}
                  as={Input}
                />
              </div>
              <Button
                type="submit"
                highlight
                fill
                size={Button.sizes.large}
                busy={busy}
              >
                Login
              </Button>
            </Form>
          )}
        </Formik>
      </div>
      <div className={styles.copy}>{config.title}</div>
    </div>
  );
};

export default LoginPage;
