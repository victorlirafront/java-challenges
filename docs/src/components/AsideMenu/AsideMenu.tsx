import { CLOSE_MENU_HAMBURGUER } from '@/constants/images';
import { StyledAsideMenu } from './AsideMenu.styled';
import { AsideMenuProps } from './AsideMenu.types';
import Image from 'next/image';
import Link from 'next/link';

function AsideMenu(props: AsideMenuProps) {
  const { className, onToggleAsideMenu } = props;

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
        <li className="option">About</li>
        <li className="option">Documentation</li>
        <li className="option">
          <Link href="/login">Login</Link>
        </li>
        <li className="option">
          <Link href="/signup">SignUp</Link>
        </li>
      </ul>
    </StyledAsideMenu>
  );
}

export default AsideMenu;
