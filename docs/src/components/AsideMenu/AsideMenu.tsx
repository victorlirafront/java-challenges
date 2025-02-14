import { CLOSE_MENU_HAMBURGUER } from '@/constants/images';
import { StyledAsideMenu } from './AsideMenu.styled';
import { AsideMenuProps } from './AsideMenu.types';
import Image from 'next/image';
import Link from 'next/link';
import { useAuth } from '@/context/Auth';
import router from 'next/router';

function AsideMenu(props: AsideMenuProps) {
  const { className, onToggleAsideMenu } = props;
  const { isAuthenticated, logout } = useAuth();

  const handleLogout = () => {
    logout();
    router.push('/auth/login');
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
      <ul className="menu-options">
        <li className="option">
          <Link href="/">Home</Link>
        </li>
        {isAuthenticated && (
          <li className="option">
            <Link href="/auth/profile">Profile</Link>
          </li>
        )}
        {!isAuthenticated && (
          <li className="option">
            <Link href="/auth/login">Login</Link>
          </li>
        )}
        {!isAuthenticated && (
          <li className="option">
            <Link href="/auth/signup">SignUp</Link>
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
