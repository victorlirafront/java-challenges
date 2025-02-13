import { AsideMenuProvider } from '@/context/AsideMenuContext';
import { AuthProvider } from '@/context//Auth';
import Layout from '@/layout/Layout';
import '@/styles/globals.css';
import type { AppProps } from 'next/app';

export default function App({ Component, pageProps }: AppProps) {
  return (
    <AuthProvider>
      <AsideMenuProvider> 
        <Layout>
          <Component {...pageProps} />
        </Layout>
      </AsideMenuProvider>
    </AuthProvider>
  );
}