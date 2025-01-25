import styled from "styled-components";


export const StyledBlocksWrapper = styled.div`
  border: 15px solid #e5e9f0;
  width: 900px;
  margin: 0 auto;
  margin-bottom: 3rem;
  border-radius: 6px;
  font-family: Martel, serif;
  padding: 30px;

  &:first-child{
    margin-top: 100px;
  }

  .title {
    font-size: 30px;
    line-height: 4rem;
    color: #2e3440;
  }

  .paragraph {
    font-size: 18px;
    line-height: 2.3rem;
  }
`