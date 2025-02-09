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
  display: flex;
  align-items: center;
  justify-content: center;
  border-left: 1rem solid #e5e9f0;

  &.active {
    right: 0;
  }
  
  .close-menu {
    padding: 10px;
    color: #000;
    cursor: pointer;
    position: absolute;
    top: 20px;
    right: 20px;
  }

  .menu-options {
    list-style: none;
    text-align: center;

    .option {
      padding: 20px 0;
      font-family: Arial, Helvetica, sans-serif;
      font-weight: 600;
      color: #555;
      cursor: pointer;
    }
  }
`