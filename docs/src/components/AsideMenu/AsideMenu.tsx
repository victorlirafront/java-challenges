import { CLOSE_MENU_HAMBURGUER } from '@/constants/images';
import { StyledAsideMenu } from './AsideMenu.styled';
import { AsideMenuProps } from './AsideMenu.types';
import Image from 'next/image';

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
        <li className="option">Home</li>
        <li className="option">About</li>
        <li className="option">Documentation</li>
        <li className="option">Account</li>
        <li className="option">Logout</li>
      </ul>
    </StyledAsideMenu>
  );
}

export default AsideMenu;
