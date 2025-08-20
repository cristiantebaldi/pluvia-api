# Pluvia - Back-end

![Status](https://img.shields.io/badge/status-em%20desenvolvimento-yellow)
![Go](https://img.shields.io/badge/Go-1.24%2B-blue.svg)
![PostgreSQL](https://img.shields.io/badge/PostgreSQL-17-blue)

Este é o repositório do back-end do projeto "Pluvia", uma aplicação web destinada a notificar os cidadãos sobre riscos de enchentes em sua localidade. A API é responsável por gerenciar os dados, processar informações meteorológicas e enviar os alertas.

Este projeto é o final do curso de Desenvolvimento de Sistemas.

## 🎯 Objetivo

Fornecer uma API robusta e escalável para dar suporte à aplicação Pluvia, garantindo que os dados sobre alertas de enchentes sejam coletados, processados e distribuídos de forma eficiente e confiável.

## ✨ Funcionalidades

- API REST para gerenciar dados de alertas, usuários e localidades.
- Integração com fontes de dados meteorológicos (a ser definido).
- Lógica para processamento de dados e identificação de riscos.
- Sistema para disparo de notificações.

## 🛠️ Tecnologias Utilizadas

- **Linguagem:** [Go (Golang)](https://golang.org/)
- **Banco de Dados:** [PostgreSQL](https://www.postgresql.org/)
- **Roteador HTTP:** `Gin`

## 🚀 Como Executar o Projeto

### Pré-requisitos

- [Go](https://go.dev/doc/install) (versão 1.24 ou superior)
- [PostgreSQL](https://www.postgresql.org/download/)
- [Git](https://git-scm.com/)

### Instalação

1.  Clone o repositório:
    ```sh
    git clone https://github.com/seu-usuario/pluvia-backend.git
    ```
2.  Navegue até o diretório do projeto:
    ```sh
    cd pluvia-backend
    ```
3.  **Configure o banco de dados:**
    - Crie um banco de dados no PostgreSQL chamado `pluvia_db`.
    - Renomeie o arquivo `.env.example` para `.env`.
    - Preencha o arquivo `.env` com as suas credenciais de acesso ao PostgreSQL.
4.  Instale as dependências:
    ```sh
    go mod tidy
    ```

### Executando a Aplicação

1.  Inicie o servidor:
    ```sh
    docker-compose up
    ```
2.  O servidor estará disponível em `http://localhost:3000`.

## 🧑‍💻 Equipe

- **Leonardo Antunes Benedetti** - [GitHub](https://github.com/Leonardo-Benedetti-Antunes)
- **Lorenzo de Maman Nied** - [GitHub](https://github.com/LorenzoNied)
- **Gian Roso** - [GitHub](https://github.com/Gian-Roso)
- **Cristian Luís Tebaldi** - [GitHub](https://github.com/cristiantebaldi)
