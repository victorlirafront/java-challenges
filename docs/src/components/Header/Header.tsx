import { StyledHeader } from "./Header.styled";

function Header(){

  const smallText = "(Developed with Golang)"

  return (
    <StyledHeader>
      <div className="container">
        <div className="box">
          <h1 className="title">Blog API: <span>Documentation</span></h1>
          <small>{smallText}</small>
        </div>
        <div  className="golang">
        </div>
      </div>
    </StyledHeader>
  )
}

export default Header;