import { CLOSE_MENU_HAMBURGUER } from '@/constants/images';
import { StyledAsideMenu } from './AsideMenu.styled';
import { AsideMenuProps } from './AsideMenu.types';
import Image from 'next/image';
import { useAuth } from '@/context/Auth';
import router from 'next/router';

function AsideMenu(props: AsideMenuProps) {
  const { className, onToggleAsideMenu, onCloseMenu } = props;
  const { isAuthenticated, logout } = useAuth();

  const handleLogout = () => {
    onCloseMenu();
    logout();
    router.push('/auth/login');
  };

  const menuHandler = function (route: string) {
    router.push(route);
  };

  return (
    <StyledAsideMenu className={className}>
      <Image
        src={CLOSE_MENU_HAMBURGUER}
        alt="close menu"
        width={40}
        height={40}
        className="close-menu"
        onClick={onToggleAsideMenu}
      />
      <ul className="menu-options" onClick={onCloseMenu}>
        <li className="option" onClick={() => menuHandler('/')}>
          Home
        </li>
        {isAuthenticated && (
          <li className="option" onClick={() => menuHandler('/auth/profile')}>
            Profile
          </li>
        )}
        {!isAuthenticated && (
          <li className="option" onClick={() => menuHandler('/auth/login')}>
            Login
          </li>
        )}
        {!isAuthenticated && (
          <li className="option" onClick={() => menuHandler('/auth/signup')}>
            SignUp
          </li>
        )}
        {isAuthenticated && (
          <li className="option" onClick={handleLogout}>
            Logout
          </li>
        )}
      </ul>
    </StyledAsideMenu>
  );
}

export default AsideMenu;
