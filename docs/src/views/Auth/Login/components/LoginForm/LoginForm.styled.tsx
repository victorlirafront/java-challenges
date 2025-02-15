import styled from 'styled-components';

export const StyledLoginForm = styled.div`
  text-align: center;
  display: flex;
  flex: 1;
  min-height: calc(100vh - 342px);

  form {
    width: 300px;
    margin: 0 auto;
    font-family: Arial, Helvetica, sans-serif;
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: center;

    h1 {
      margin-bottom: 20px;
    }

    .form-control {
      display: flex;
      flex-direction: column;
      align-items: flex-start;
      width: 100%;
      position: relative;

      &:not(:first-of-type) {
        margin-top: 40px;
      }

      input {
        width: 100%;
        padding: 10px;
        border-radius: 5px;
        border: none;
        margin-top: 5px;
      }

      .error-message {
        position: absolute;
        font-size: 14px;
        bottom: -20px;
        color: red;
      }
    }

    button {
      width: 100%;
      padding: 10px;
      border: none;
      background: #027d9c;
      color: #fff;
      font-weight: bolder;
      margin-top: 40px;
      border-radius: 4px;
      cursor: pointer;
      transition: 0.3s;

      &:hover {
        background: #016d88;
      }
    }


  }
`;
