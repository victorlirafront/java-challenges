import { StyledHeader } from './Header.styled';
import Image from 'next/image';
import { HeaderProps } from './Header.types';

function Header(props: HeaderProps) {
  const { onToggleAsideMenu } = props;
  const smallText = '(Developed with Golang)';

  return (
    <StyledHeader>
      <div className="container">
        <div className="box">
          <h1 className="title">
            Blog API: <span>Documentation</span>
          </h1>
          <small>{smallText}</small>
        </div>
        <div className="golang"></div>
      </div>
      <div className="harburguer" onClick={onToggleAsideMenu}>
        <Image
          src="https://go.dev/images/menu-24px-white.svg"
          width={45}
          height={45}
          alt="hamburguer"
        />
      </div>
    </StyledHeader>
  );
}

export default Header;
