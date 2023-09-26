# gohexagonal

Este é um exemplo de um projeto Go que segue a arquitetura hexagonal (também conhecida como arquitetura de portas e adaptadores).

## Visão Geral

Este projeto demonstra uma estrutura básica de aplicação em Go, implementada usando a arquitetura hexagonal.

## Estrutura do Projeto

A estrutura do projeto é organizada da seguinte forma:

- `cmd`: Este diretório contém os pontos de entrada da aplicação (como a função main).
- `internal`: Este diretório contém a lógica do negócio e os detalhes da implementação.
- `pkg`: Este diretório contém arquivos de infraestrutura do projeto.
- `ent`: Este diretório contém arquivos gerados pelo ORM ENT.

## Como usar

Para usar este projeto, você pode clonar o repositório e executar o comando `go run` no diretório `cmd`.

## O que contém neste projeto
### banco de dados usados:
- `mongodb`: Este diretório contém os pontos de entrada da aplicação (como a função main).
- `redis`: Este diretório contém a lógica do negócio e os detalhes da implementação.
- `postgres`: Este diretório contém arquivos de infraestrutura do projeto.
### serviços de email:
- `smtp`: Este diretório contém arquivos gerados pelo ORM ENT.

## Contribuições

Contribuições para este projeto são bem-vindas. Por favor, abra uma issue ou um pull request se você quiser contribuir.

## Licença

Este projeto está licenciado sob a licença MIT. Veja o arquivo LICENSE para mais detalhes.

