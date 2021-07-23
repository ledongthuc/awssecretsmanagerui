import '@moai/core/dist/bundle.css';
import '@moai/core/dist/font/remote.css';
import 'styles/globals.css';

import type { AppProps } from 'next/app';

function MyApp({ Component, pageProps }: AppProps) {
  return <Component {...pageProps} />;
}

export default MyApp;
