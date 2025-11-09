# Pact Broker - Docker Compose

Este projeto utiliza o Docker Compose para subir um servidor Pact Broker, que é usado para gerenciar contratos (pactos) entre serviços (Consumer e Provider). Além disso, o projeto inclui duas aplicações: uma `Consumer` e uma `Provider`.

## Subindo o servidor Pact Broker

1. Certifique-se de ter o Docker e o Docker Compose instalados na sua máquina.
2. No diretório do projeto, execute o comando abaixo para subir o servidor:
   ```bash
   docker-compose up -d
   ```
3. O servidor Pact Broker estará disponível em: http://localhost:9292.

## Configuração do servidor

O servidor Pact Broker utiliza as seguintes configurações padrão:

- Porta: 9292
- Banco de dados: SQLite (armazenado no volume do Docker)

> Se necessário, você pode alterar as configurações no arquivo docker-compose.yml.

# Aplicações

## Consumer

A aplicação `Consumer` é responsável por gerar os contratos (pactos) que descrevem como ela espera que o Provider se comporte.

## Provider

A aplicação `Provider` é responsável por implementar os endpoints descritos nos contratos gerados pelo Consumer.

## Publicando payloads e pactos no servidor

1. Gere o contrato (pacto) no `Consumer` executando o seguinte comando:

```bash
npm run pact:generate
```

Este comando criará um arquivo .json contendo o contrato.

2. Publique o contrato no servidor Pact Broker:

```bash
pact broker publish <caminho-do-arquivo-pacto.json> --broker-base-url=http://localhost:9292 --consumer-app-version=1.0.0
```

Substitua <caminho-do-arquivo-pacto.json> pelo caminho do arquivo gerado no passo anterior.

## Validando o pacto

### No Consumer

1. Execute os testes no Consumer para garantir que ele está gerando os contratos corretamente:

```go
go test ./...
```

### No Provider

1. Baixe o contrato do servidor Pact Broker:

```bash
pact broker fetch --broker-base-url=http://localhost:9292 --provider=Provider --consumer=Consumer --latest
```

2. Execute os testes no Provider para validar o contrato:

```bash
go test -tags=pact ./...
```

> Se os testes passarem, significa que o Provider está em conformidade com o contrato gerado pelo Consumer.

## Conclusão

Com este fluxo, você pode garantir que o Consumer e o Provider estão alinhados em relação aos contratos, reduzindo problemas de integração entre serviços.