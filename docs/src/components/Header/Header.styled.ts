import { GOLANG_AVATAR } from '@/constants/images';
import styled from 'styled-components';

export const StyledHeader = styled.header`
  background: #027d9c;
  color: #fff;
  height: 315px;
  padding: 0 20px;

  @media screen and (max-width: 510px) {
    padding: 0;
  }

  .container {
    max-width: 900px;
    margin: 0 auto;
    font-family: Martel, serif;
    font-size: 30px;
    height: 100%;
    display: flex;
    justify-content: space-between;
    align-items: flex-end;
    text-align: center;

    .box {
      margin-bottom: 20px;
      .title {
        @media screen and (max-width: 800px) {
          font-size: 30px;
          line-height: 30px;
        }

        span {
          font-size: 40px;

          @media screen and (max-width: 800px) {
            font-size: 20px;
          }
        }
      }

      small {
        @media screen and (max-width: 510px) {
          font-size: 16px;
        }
      }
    }

    .golang {
      min-height: 100%;
      background-image: url(${GOLANG_AVATAR});
      width: 156px;
      background-repeat: no-repeat;
      background-position: top center;
      background-size: 150px;

      @media screen and (max-width: 768px) {
        background-size: 110px;
        background-position: 44px 0px;
      }
    }
  }

  .harburguer {
    padding: 30px 40px;
    position: absolute;
    width: 100%;
    right: 50%;
    transform: translateX(50%);
    top: 0;
    text-align: end;
  }
`;
