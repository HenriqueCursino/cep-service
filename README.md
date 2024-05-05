# cep-service

## Descrição

O objetivo deste projeto é desenvolver um serviço dedicado à busca rápida de CEPs, utilizando as principais APIs de CEP disponíveis no Brasil. Para a implementação, foi escolhida a linguagem Go para realizar buscas de forma paralela e concorrente, aproveitando as goroutines.

Em situações em que um CEP válido não retorna um endereço correspondente, o sistema realizará iterações, substituindo gradualmente os dígitos da direita para a esquerda por zero, até que o endereço seja encontrado. Por exemplo, ao fornecer o CEP 22333999 e não obter resultado, o sistema tentará com variações como 22333990, 22333900, e assim por diante, até obter sucesso.

## Arquitetura

Para este projeto, escolhi uma arquitetura simples, pois estou lidando apenas com o contexto de CEPs. Me baseei nos padrões arquiteturais para APIs REST, com os quais muitos desenvolvedores estão familiarizados. O repositório `config` é responsável pela configuração inicial do projeto. Já o repositório `api` engloba todos os principais pacotes da API (`controller`, `service`).

# 1 - Como rodar o projeto

## Variáveis de Ambiente

Crie o arquivo `.env` seguindo o exemplo do `env.example` e configure as variáveis de ambiente necessárias.
O token usado na API de busca de endereço está protegido por variável de ambiente.

## Iniciando o Projeto

Para executar a aplicação, siga os seguintes passos:

1. `go mod tidy`
2. `go run main.go`

Ou se preferir pode executar utilizando Makefile:

1. `make install`
2. `make run`

## Documentação

Se desejar usar o Swagger para a documentação da API, execute o seguinte comando:

- `swag init`

Acesse a documentação do Swagger em:

- [http://localhost:8080/swagger/index.html](http://localhost:8080/swagger/index.html)

# 2 - Como funciona as requisições no navegador

1. Digitar a URL no navegador: O cliente (seu navegador) solicita o site digitado.
2. Resolução de DNS: O navegador traduz o nome do site em um endereço IP usando o DNS.
3. Estabelecimento de conexão: O navegador abre uma conexão TCP com o servidor.
4. Envio da solicitação HTTP: O navegador envia uma solicitação HTTP, especificando o tipo de requisição (GET, POST, etc.).
5. Processamento da solicitação: O servidor recebe a solicitação, processa e responde.
6. Recebimento da resposta HTTP: O navegador recebe a resposta do servidor.
7. Renderização da página: O navegador renderiza a página HTML recebida.
8. Exibição da página: A página do site é exibida para o usuário.
