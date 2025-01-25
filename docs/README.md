
```plaintext
backend/
│
├── cmd/                     # Ponto de entrada para a aplicação
│   └── main.go              # Arquivo principal para iniciar o servidor
│
├── config/                  # Configurações gerais da aplicação
│   └── config.go            # Gerenciamento de variáveis de ambiente e configurações
│
├── controllers/             # Lógica de controle para rotas
│   ├── postController.go    # Controlador de posts
│   ├── userController.go    # Controlador de usuários
│   └── commentController.go # Controlador de comentários
│
├── routes/                  # Definição de rotas da API
│   ├── postRoutes.go        # Rotas relacionadas a posts
│   ├── userRoutes.go        # Rotas relacionadas a usuários
│   └── commentRoutes.go     # Rotas relacionadas a comentários
│
├── models/                  # Definição de estruturas e interação com o banco de dados
│   ├── post.go              # Modelo de post
│   ├── user.go              # Modelo de usuário
│   └── comment.go           # Modelo de comentário
│
├── services/                # Lógica de negócio e regras complexas
│   ├── postService.go       # Serviço para posts
│   ├── userService.go       # Serviço para usuários
│   └── commentService.go    # Serviço para comentários
│
├── repositories/            # Interação direta com o banco de dados
│   ├── postRepository.go    # Repositório de posts
│   ├── userRepository.go    # Repositório de usuários
│   └── commentRepository.go # Repositório de comentários
│
├── middleware/              # Middleware para autenticação, logs, etc.
│   ├── authMiddleware.go    # Middleware de autenticação
│   ├── loggerMiddleware.go  # Middleware de logs
│   └── corsMiddleware.go    # Middleware para CORS
│
├── utils/                   # Funções auxiliares e utilitárias
│   ├── jwtUtils.go          # Funções para manipular tokens JWT
│   ├── hashUtils.go         # Funções para hashing (ex: senhas)
│   └── validationUtils.go   # Funções para validações
│
├── database/                # Conexão e migrações de banco de dados
│   ├── connection.go        # Configuração de conexão com o banco
│   └── migrations.go        # Scripts para migrações
│
├── docs/                    # Documentação da API
│   └── swagger.yaml         # Arquivo para documentação Swagger
│
└── go.mod                   # Gerenciamento de dependências
```
