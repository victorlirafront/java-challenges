import styled from "styled-components";


export const StyledBlocksWrapper = styled.div`
  border: 15px solid #e5e9f0;
  max-width: 900px;
  margin: 0 auto;
  margin-bottom: 3rem;
  border-radius: 6px;
  font-family: sans-serif;
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
    line-height: 30px;

    a {
      color: #007bff;
    }
  }

  .api-table {
  width: 100%;
  border-collapse: collapse;
  margin: 1rem 0;
  font-size: 1rem;
  text-align: left;
  border: 1px solid #ddd;

  thead {
    background-color: #f5f5f5;

    th {
      padding: 0.75rem;
      border: 1px solid #ddd;
      font-weight: bold;
      color: #333;
      text-transform: uppercase;
    }
  }

  tbody {
    tr {
      &:nth-child(even) {
        background-color: #f9f9f9; // Alterna as cores das linhas
      }

      &:hover {
        background-color: #e6f7ff; // Destaque ao passar o mouse
      }

      td {
        padding: 0.75rem;
        border: 1px solid #ddd;

        &.method {
          font-weight: bold;
          font-size: 14px;
          color: #007bff;
          text-transform: uppercase;
        }

        &.route {
          font-family: monospace;
          color: #555;
          line-height: 24px;
        }
      }
    }
  }

  // Responsividade
  @media (max-width: 768px) {
    font-size: 0.9rem;

    thead {
      display: none; // Esconde o cabeçalho em telas pequenas
    }

    tbody {
      display: block;

      tr {
        display: flex;
        flex-direction: column;
        margin-bottom: 1rem;
        border: 1px solid #ddd;
        border-radius: 8px;
        overflow: hidden;
        box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);

        td {
          display: flex;
          justify-content: space-between;
          padding: 0.5rem;

          &:before {
            content: attr(data-label); // Insere rótulos semânticos
            font-weight: bold;
            text-transform: uppercase;
            color: #333;
            margin-right: 0.5rem;
            display: none;
          }

          &.method {
            color: #007bff;
          }
        }
      }
    }
  }
}

`