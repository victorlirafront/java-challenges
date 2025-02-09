import { useAsideMenu } from '@/context/AsideMenuContext';
import Header from '@/components/Header/Header';
import { LayoutProps } from './Layout.types';
import Footer from '@/components/Footer/Footer';

const Layout = ({ children }: LayoutProps) => {
  const { toggleAsideMenu } = useAsideMenu();

  return (
    <>
      <Header onToggleAsideMenu={toggleAsideMenu} />
      <main>{children}</main>
      <Footer />
    </>
  );
};

export default Layout;
