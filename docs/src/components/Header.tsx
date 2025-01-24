import { StyledHeader } from "./Header.styled";

function Header(){
  return (
    <StyledHeader>
      <div className="container">
        <h1 className="title">API: Documentacão</h1>
        <div  className="golang">
        </div>
      </div>
    </StyledHeader>
  )
}

export default Header;