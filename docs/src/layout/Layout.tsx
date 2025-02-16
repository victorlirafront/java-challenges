import { useAsideMenu } from '@/context/AsideMenuContext';
import Header from '@/components/Header/Header';
import { LayoutProps } from './Layout.types';
import Footer from '@/components/Footer/Footer';
import AsideMenu from '@/components/AsideMenu/AsideMenu';

const Layout = ({ children }: LayoutProps) => {
  const { displayAsideMenu, toggleAsideMenu, closeMenu } = useAsideMenu();

  return (
    <>
      <Header onToggleAsideMenu={toggleAsideMenu} />
      <main data-aos="fade-down" data-aos-delay="100" data-aos-offset="0">{children}</main>
      <AsideMenu
        onToggleAsideMenu={toggleAsideMenu}
         onCloseMenu={closeMenu}
        className={`${displayAsideMenu ? 'active' : ''}`}
      />
      <Footer data-aos="fade-up" data-aos-delay="100" data-aos-offset="0" />
    </>
  );
};

export default Layout;
