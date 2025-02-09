import { useAsideMenu } from '@/context/AsideMenuContext';
import Header from '@/components/Header/Header';
import { LayoutProps } from './Layout.types';
import Footer from '@/components/Footer/Footer';
import AsideMenu from '@/components/AsideMenu/AsideMenu';

const Layout = ({ children }: LayoutProps) => {
  const { displayAsideMenu, toggleAsideMenu } = useAsideMenu();

  return (
    <>
      <Header onToggleAsideMenu={toggleAsideMenu} />
      <main>{children}</main>
      <AsideMenu
        onToggleAsideMenu={toggleAsideMenu}
        className={`${displayAsideMenu ? 'active' : ''}`}
      />
      <Footer />
    </>
  );
};

export default Layout;
