import { useAsideMenu } from '@/context/AsideMenuContext';
import Header from '@/components/Header/Header';
import { LayoutProps } from './Layout.types';
import Footer from '@/components/Footer/Footer';
import AsideMenu from '@/components/AsideMenu/AsideMenu';

const Layout = ({ children }: LayoutProps) => {
  const { displayAsideMenu, toggleAsideMenu } = useAsideMenu();

  return (
    <>
      <Header data-aos="fade-up" data-aos-delay="100" data-aos-offset="0" onToggleAsideMenu={toggleAsideMenu} />
      <main data-aos="fade-down" data-aos-delay="100" data-aos-offset="0">{children}</main>
      <AsideMenu
        onToggleAsideMenu={toggleAsideMenu}
        className={`${displayAsideMenu ? 'active' : ''}`}
      />
      <Footer data-aos="fade-up" data-aos-delay="100" data-aos-offset="0" />
    </>
  );
};

export default Layout;
