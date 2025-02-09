import { AsideMenuProvider } from '@/context/AsideMenuContext';
import Layout from '@/layout/Layout';
import '@/styles/globals.css';
import type { AppProps } from 'next/app';

export default function App({ Component, pageProps }: AppProps) {
  return (
    <AsideMenuProvider> 
      <Layout>
        <Component {...pageProps} />
      </Layout>
    </AsideMenuProvider>
  );
}
