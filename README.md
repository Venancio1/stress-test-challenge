# stress-test-challenge

Esta é uma ferramenta de teste de estresse para URLs, escrita em Go. Ela realiza múltiplas requisições HTTP simultâneas para avaliar o desempenho e a resposta de um servidor.

## Pré-requisitos

- Docker instalado no sistema.

## Construindo a Imagem Docker

Para construir a imagem Docker, execute o seguinte comando no diretório raiz do projeto:

```bash
docker build -t stress-test .
```

Este comando criará uma imagem chamada `stress-test` baseada no `Dockerfile` fornecido.

## Executando o Teste

Após construir a imagem, você pode executar o teste de estresse usando o comando abaixo:

```bash
docker run stress-test --url=https://httpbin.org  --requests=1000 --concurrency=10
```

### Parâmetros

- `--url`: A URL a ser testada.
- `--requests`: Número total de requisições a serem realizadas (padrão: 1).
- `--concurrency`: Número de requisições simultâneas (padrão: 1).

### Exemplo de Saída

A ferramenta exibirá o tempo total de execução, o número de ocorrências por código de status HTTP e o total de requisições realizadas.

Exemplo:

```
Tempo total de execução: 2.345s
Status 200: 950 ocorrências
Status 500: 50 ocorrências
Total de requisições realizadas: 1000
```