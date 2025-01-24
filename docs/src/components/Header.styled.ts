import styled from "styled-components";

export const StyledHeader = styled.header`
  background: #027d9c;
  color: #fff;
  height: 315px;

  .container {
    max-width: 900px;
    margin: 0 auto;
    font-family: Martel,serif;
    font-size: 30px;
    height: 100%;
    display: flex;
    justify-content: space-between;
    align-items: flex-end;

    .golang {
      min-height: 100%;
      background-image: url(https://go.dev/images/gophers/ladder.svg);
      width: 300px;
      background-repeat: no-repeat;
      background-position: top center;
      background-size: 150px;
    }
  }
`