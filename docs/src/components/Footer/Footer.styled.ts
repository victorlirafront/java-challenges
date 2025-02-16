import styled from 'styled-components';

export const StyledFooter = styled.footer`
  background: #253443;
  color: #fff;
  font-family: Martel, sans-serif;
  padding: 50px 0px;

  .container {
    max-width: 900px;
    margin: 0 auto;
    text-align: center;

    p {
      @media screen and (max-width: 768px) {
        font-size: 14px;
        line-height: 24px;
        margin-top: 10px;
      }
    }
  }
`;
