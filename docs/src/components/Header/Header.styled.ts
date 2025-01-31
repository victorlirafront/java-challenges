import { GOLANG_AVATAR } from "@/constants/images";
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

    .title {
      margin-bottom: 20px;

      span {
        font-size: 40px;
      }
    }

    .golang {
      min-height: 100%;
      background-image: url(${GOLANG_AVATAR});
      width: 156px;
      background-repeat: no-repeat;
      background-position: top center;
      background-size: 150px;
    }
  }
`