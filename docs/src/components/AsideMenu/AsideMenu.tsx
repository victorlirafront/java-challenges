import { StyledAsideMenu } from "./AsideMenu.styled";
import { AsideMenuProps } from "./AsideMenu.types"

function AsideMenu(props: AsideMenuProps) {
  const { className, onToggleAsideMenu} = props

  return (
    <StyledAsideMenu className={className}>
      <button className="close-menu" onClick={onToggleAsideMenu}>close</button>
    </StyledAsideMenu>
  );
}

export default AsideMenu
