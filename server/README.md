# Guia de Instalação

## Pre-Requisitos

Para executar o projeto, por padrão, utilizasse:

- Docker
- Go

## Inicializando o Bando de dados

1 - Rodar Docker compose

```bash
sudo docker compose up
```

2 - Sugestões de ferramentas para fazer o gerenciamento do banco de dados
    - PGAdmin
    - Dbeaver

## Executando o projeto

Para executar o projeto basta utilizar o comando

```bash
go run cmd/main.go
```

# Entidades

1. **Usuarios**: a entidade irá auxiliar na construção da autenticação do projeto
2. **Admins**: responsáveis por fazer o controle de adoções (ONGs, Canis, Secretarias, Encarregados, ...)
3. **Tutores**: usuários que irão demonstrar o interesse em adotar um animal.
4. **Endereço**: endereços que serão incorporados aos tutores
5. **Especie**: a especie do animal que poderá ser adotado
6. **Raça**: raça do animal que poderá ser adotado
7. **Animal**: animal que está disponivel para adoção

![modelagem](docs/modelagem_gaart.png)

# Rotas

1. **METADATA**

    - `ping`: retorna pong
    - `version`: retorna versão do recurso

2. **Usuário**

    - `usuario`: criar um usuario

3. **Especies**

    - [GET] `especies`: retorna uma listagem de todas as especies cadastradas
    - [POST] `especie`: cria uma nova instancia de especie
    - [PATCH] `especie/:id`: atualiza os dados de uma especie

4. **Raças**

    - [GET] `racas`: retorna uma listagem com todas as raças cadastradas
    - [POST] `raca`: cria uma nova instancia de raça
    - [PATCH] `raca/:id`: atualiza os dados de uma raça

5. **Animais**

    - [GET] `animais`: retorna uma listagem de todos os animais cadastrados
    - [GET] `animais/:id`: realiza uma busca de um animal pelo id
    - [POST] `animal`: cria uma nova instancia de animal
    - [PATCH] `animal/:id`: atualiza os dados de um animal

6. **Endereços**

    - [GET] `enderecos`: retorna uma listagem de todos os endereços cadastrados
    - [POST] `endereco`: cria uma nova instancia de endereço
    - [PATCH] `endereco/:id`: atualiza os dados de um endereço

7. **Tutor**

    - [GET] `tutores`: retorna uma listagem de todos os tutores
    - [POST] `tutor`: cria uma nova instancia de tutor