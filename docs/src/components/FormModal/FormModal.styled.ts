import styled from 'styled-components';

export const StyledFormModal = styled.div`
  position: absolute;
  right: 50%;
  width: 300px;
  top: -100px;
  transform: translateX(50%);
  display: flex;
  align-items: center;
  justify-content: center;
  background: #fff;
  padding: 15px 10px;
  border-radius: 4px;
  box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);
  opacity: 0;
  transition: transform 0.3s, opacity 0.3s;

  &.active {
    transform: translateX(50%) translateY(140px);  /* Ajuste para mover no eixo Y */
    opacity: 1;
  }

  p {
    font-family: sans-serif;
    margin-left: 10px;
    font-size: 13px;
  }
`;
