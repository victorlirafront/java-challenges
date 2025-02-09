import styled from "styled-components";


export const StyledAsideMenu = styled.aside`
  background: #edeff3;
  width: 320px;
  position: fixed;
  z-index: 10;
  top: 0;
  height: 100vh;
  padding: 20px;
  text-align: end;
  right: -320px;
  transition: 0.3s;

  &.active {
    right: 0;
  }
  
  .close-menu {
    padding: 10px;
    color: #000;
    cursor: pointer;
  }
`