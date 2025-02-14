import styled from 'styled-components';

const StyledProfileContainer = styled.div`
  min-height: calc(100vh - 315px);
  display: flex;
  align-items: center;
  flex-direction: column;
  justify-content: center;

  .container {
    border: 15px solid #e5e9f0;
    max-width: 900px;
    margin: 0 auto;
    margin-bottom: 3rem;
    border-radius: 6px;
    font-family: sans-serif;
    padding: 30px;
  }
`;

export default StyledProfileContainer;
