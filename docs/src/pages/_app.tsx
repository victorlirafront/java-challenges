import { AsideMenuProvider } from '@/context/AsideMenuContext';
import { AuthProvider } from '@/context//Auth';
import Layout from '@/layout/Layout';
import '@/styles/globals.css';
import type { AppProps } from 'next/app';
import Aos from 'aos';
import { useEffect } from 'react';
import 'aos/dist/aos.css';

export default function App({ Component, pageProps }: AppProps) {

  useEffect(() => {
    Aos.init();
  }, []);

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