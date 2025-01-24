import { StyledHeader } from "./Header.styled";

function Header(){
  return (
    <StyledHeader>
      <div className="container">
        <h1 className="title">Blog API: <span>Documentation</span></h1>
        <div  className="golang">
        </div>
      </div>
    </StyledHeader>
  )
}

export default Header;