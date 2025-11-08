# Clean Architecture + Spring Boot + DDD Sample

Este projeto demonstra uma API REST construída com Spring Boot seguindo princípios de Clean Architecture e Domain-Driven Design (DDD). A API expõe recursos para cadastrar e listar artistas em uma implementação simples usando um repositório em memória.

## Executando o projeto

1. **Pré-requisitos**: JDK 11+ e Gradle (ou utilize o wrapper incluso).
2. **Subir a aplicação**:
   ```bash
   ./gradlew bootRun
   ```
   A API será exposta em `http://localhost:8080`.

## Autenticação

Não há autenticação; todas as rotas são públicas.

## Requisições disponíveis

### Listar artistas

- **Endpoint**: `GET /artists`
- **Descrição**: Retorna todos os artistas cadastrados, ordenados pela data de nascimento.
- **Dados iniciais**
  - James Alan Hetfield (alias "James Hetfield")
  - Mike Portnoy
  - Mat Heaffy
- **Resposta 200**:

```json
[
  {
    "id": "80b33f54-7d6a-4fb3-a4ed-8a1d2a7c3b83",
    "name": "James Alan Hetfield",
    "birthday": "1963-08-03",
    "alias": "James Hetfield"
  },
  {
    "id": "5c6eee56-0501-4e3e-97b7-4ccf3b0248db",
    "name": "Mike Portnoy",
    "birthday": "1967-04-20",
    "alias": null
  },
  {
    "id": "2d4b5a30-4931-4b3c-83a9-9d8cb2da0021",
    "name": "Mat Heaffy",
    "birthday": "1986-01-26",
    "alias": null
  }
]
```

> Os registros acima são carregados automaticamente ao iniciar o serviço.
- Os valores de `id` mudarão a cada execução, pois são gerados como UUID.

**Exemplo usando cURL**

```bash
curl -X GET http://localhost:8080/artists
```

### Criar artista

- **Endpoint**: `POST /artists`
- **Descrição**: Cria um novo artista. O campo `alias` é opcional.
- **Headers**: `Content-Type: application/json`
- **Body**:

```json
{
  "name": "Geddy Lee",
  "birthday": "1953-07-29",
  "alias": "Geddy"
}
```

- **Resposta 200**:

```json
{
  "id": "9dca8718-87d0-4b2a-b283-01cdfea3b939",
  "name": "Geddy Lee",
  "birthday": "1953-07-29",
  "alias": "Geddy"
}
```

**Exemplo usando cURL**

```bash
curl -X POST http://localhost:8080/artists \
  -H "Content-Type: application/json" \
  -d '{
        "name": "Geddy Lee",
        "birthday": "1953-07-29",
        "alias": "Geddy"
      }'
```

### Erros comuns

- **400 Bad Request**: ocorre quando `birthday` está em branco ou em um formato inválido (`yyyy-MM-dd` é obrigatório).

## Estrutura da API

```java
@RestController
public class ArtistController {

    @GetMapping("/artists")
    public List<ArtistResponse> artists() {
        return artistApplicationService.findAll();
    }

    @PostMapping("/artists")
    public ArtistResponse create(@RequestBody CreateArtistCommand command) {
        return artistApplicationService.create(command);
    }
}
```

- `ArtistApplicationService` contém a lógica de uso de caso.
- `SamplePopulationService` adiciona dados iniciais ao repositório em memória.
- `ArtistResponse` padroniza o payload de saída com `id`, `name`, `birthday` e `alias`.

## Próximos passos sugeridos

- Persistir os dados em um banco real (por exemplo, PostgreSQL).
- Adicionar validação avançada e tratamento de erros com mensagens mais detalhadas.
- Criar testes de integração para os endpoints.
