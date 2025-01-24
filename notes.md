# API Endpoints para o Blog

## 1. Listar todos os posts
- **Método**: `GET`
- **Endpoint**: `/posts`
- **Descrição**: Retorna uma lista de todos os posts com informações como título, autor, data de publicação e resumo.

---

## 2. Pesquisar posts
- **Método**: `GET`
- **Endpoint**: `/posts/search`
- **Parâmetros**: 
  - `q`: Query string para busca por título, autor ou conteúdo parcial.
- **Exemplo**: `/posts/search?q=javascript`
- **Descrição**: Retorna os posts que correspondem ao termo de busca.

---

## 3. Visualizar um post específico
- **Método**: `GET`
- **Endpoint**: `/posts/{id}`
- **Descrição**: Retorna os detalhes de um post único com base no ID.
- **Exemplo de Retorno**:
  - Título
  - Corpo do post
  - Autor
  - Data de publicação
  - Tags

---

## 4. Listar posts por categoria
- **Método**: `GET`
- **Endpoint**: `/posts/category/{categoria}`
- **Exemplo**: `/posts/category/tecnologia`
- **Descrição**: Retorna todos os posts pertencentes a uma categoria específica.

---

## 5. Filtrar posts por autor
- **Método**: `GET`
- **Endpoint**: `/posts/author/{nomeAutor}`
- **Exemplo**: `/posts/author/Victor`
- **Descrição**: Mostra todos os posts escritos por um autor específico.

---

## 6. Ordenar posts
- **Método**: `GET`
- **Endpoint**: `/posts?orderBy={field}&direction={asc|desc}`
- **Parâmetros**:
  - `orderBy`: Campo para ordenar (ex.: `data` ou `titulo`).
  - `direction`: Direção (`asc` para crescente, `desc` para decrescente).
- **Exemplo**: `/posts?orderBy=data&direction=desc`
- **Descrição**: Retorna os posts ordenados de acordo com o campo e direção especificados.

---

## 7. Filtrar por data de publicação
- **Método**: `GET`
- **Endpoint**: `/posts?startDate={dataInicial}&endDate={dataFinal}`
- **Exemplo**: `/posts?startDate=2025-01-01&endDate=2025-01-31`
- **Descrição**: Retorna posts publicados dentro de um intervalo de datas.

---

## 8. Obter tags disponíveis
- **Método**: `GET`
- **Endpoint**: `/tags`
- **Descrição**: Lista todas as tags usadas nos posts, permitindo que os usuários filtrem os posts por tags.

---

## 9. Listar posts por tag
- **Método**: `GET`
- **Endpoint**: `/posts/tag/{tag}`
- **Exemplo**: `/posts/tag/javascript`
- **Descrição**: Retorna todos os posts associados a uma tag específica.

---

## 10. Obter estatísticas básicas
- **Método**: `GET`
- **Endpoint**: `/stats`
- **Descrição**: Retorna estatísticas básicas do blog, como:
  - Número total de posts.
  - Número de autores únicos.
  - Número de categorias e tags usadas.
