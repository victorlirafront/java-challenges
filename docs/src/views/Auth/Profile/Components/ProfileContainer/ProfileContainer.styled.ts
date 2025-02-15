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
    margin: 100px 0 100px 0;

    .token-container {
      .btn-copy {
        background: red;
        display: block;
        justify-self: end;
        background: #027d9c;
        color: #fff;
        padding: 10px;
        border: none;
        border-radius: 4px;
      }

      code {
        word-break: break-word;
        white-space: pre-wrap;
        overflow-wrap: break-word;
        display: block;
        padding: 8px;
        border-radius: 4px;
        overflow-x: auto;
        background: #1e1e1e;
        color: #dbdba9;
        padding: 50px 20px;
        margin-top: 20px;
      }
    }

    .code-example {
      margin-top: 50px;
      code {
        word-break: break-word;
        white-space: pre-wrap;
        overflow-wrap: break-word;
        display: block;
        padding: 8px;
        border-radius: 4px;
        overflow-x: auto;
        background: #1e1e1e;
        color:rgb(6, 223, 169);
        padding: 50px 20px;
        margin-top: 20px;
      }
    }
  }
`;

export default StyledProfileContainer;
