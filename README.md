# Desafio Clean Architecture Full Cycle - Gerenciamento de Pedidos

Este projeto implementa uma abordagem de Clean Architecture para um sistema de Gerenciamento de Pedidos com múltiplas interfaces de serviço: REST, gRPC e GraphQL.

## Visão Geral do Projeto

O sistema fornece funcionalidades de gerenciamento de pedidos através de três diferentes interfaces de serviço:
- API REST (HTTP)
- Serviço gRPC
- Query GraphQL

## Pré-requisitos

- Docker e Docker Compose
- Go 1.24 (para desenvolvimento local)
- Make (opcional, para usar comandos do Makefile)

## Portas dos Serviços

A aplicação expõe os seguintes serviços:
- API REST: `:8000` - GET /orders/list - Listar todos os pedidos
- API REST: `:8000` - POST /orders/create - Criar um novo pedido
- Servidor gRPC: `:50051` - Serviço ListOrders
- Servidor GraphQL: `:8080` - Query ListOrders

## Como Começar

1. Clone o repositório:
```bash
git clone [repository-url]
cd fc-clean-arch-challenge
```

2. Inicie o banco de dados e os serviços da aplicação:
```bash
docker compose up -d
```

Este comando irá:
- Iniciar um banco de dados MySQL 8.0 na porta 3307
- Criar o banco de dados 'orders'
- Aplicar todas as migrações necessárias

3. A aplicação estará pronta para aceitar requisições através das três interfaces.

## Uso da API

### Endpoints REST
```http
# Criar um novo pedido
POST http://localhost:8000/order
Content-Type: application/json

{
    "id": "abc",
    "price": 100.0,
    "tax": 10.0
}

# Listar todos os pedidos
GET http://localhost:8000/order
```

### Serviço gRPC
O serviço gRPC está disponível em `localhost:50051` e fornece os seguintes métodos:
- `CreateOrder`: Cria um novo pedido
- `ListOrders`: Lista todos os pedidos

Para testar, você pode usar um cliente gRPC como o [Evans](https://github.com/ktr0731/evans) ou [grpcurl](https://github.com/fullstorydev/grpcurl).

### Query GraphQL
Acesse o playground GraphQL em: http://localhost:8080

```graphql
# Criar um novo pedido
mutation CreateOrder {
  createOrder(input: {
    id: "abc",
    price: 100.0,
    tax: 10.0
  }) {
    id
    price
    tax
    finalPrice
  }
}

# Listar todos os pedidos
query ListOrders {
  orders {
    id
    price
    tax
    finalPrice
  }
}
```

## Detalhes de Implementação

### REST (Clean Architecture)
1. Interface do repositório definida em `internal/entity/interface.go`
2. Implementação do repositório em `internal/infra/database/order_repository.go`
3. Casos de uso em `internal/usecase/`:
   - `create_order_usecase.go`
   - `list_orders_usecase.go`
4. Handlers HTTP em `internal/infra/web/order_handle.go`

### gRPC
1. Definição do protobuf em `internal/infra/grpc/pb/order.proto`
2. Código gerado em:
   - `internal/infra/grpc/pb/order.pb.go`
   - `internal/infra/grpc/pb/order_grpc.pb.go`
3. Implementação do serviço em `internal/infra/grpc/service/order_service.go`
4. Utiliza os mesmos casos de uso da aplicação

### GraphQL
1. Schema definido em `internal/infra/graph/schema.graphqls`
2. Resolvers implementados em `internal/infra/graph/schema.resolvers.go`
3. Modelos gerados em `internal/infra/graph/model/`
4. Utiliza os mesmos casos de uso da aplicação

Todos os três endpoints (REST, gRPC e GraphQL) compartilham a mesma camada de domínio e casos de uso, seguindo os princípios da Clean Architecture.

## Configuração do Banco de Dados

O banco de dados MySQL está configurado com as seguintes configurações:
- Host: localhost
- Porta: 3307
- Banco de Dados: orders
- Senha Root: root

## Desenvolvimento

Para executar a aplicação localmente sem Docker:

1. Certifique-se que o banco de dados está rodando:
```bash
docker compose up mysql -d
```

2. Execute a aplicação:
```bash
go run cmd/servers/main.go
```

## Testes

Para executar os testes:
```bash
go test ./...
```
