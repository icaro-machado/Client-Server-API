# Client-Server-API
Desafio Pós Go Expert - Client-Server-API

Este projeto consiste em dois programas em Go: um servidor (`server.go`) que busca a cotação do dólar em uma API externa, armazena o valor em um banco de dados MySQL usando o GORM, e um cliente (`client.go`) que consulta o servidor, recebe a cotação e a salva em um arquivo de texto.

## Estrutura do Projeto

- **client.go**: Programa cliente que faz uma requisição ao servidor, recebe a cotação e salva em um arquivo.
- **server.go**: Webserver que consulta a API externa para obter a cotação do dólar, armazena o resultado no banco de dados MySQL, e retorna o valor para o cliente.

---

## Como Executar

### Pré-requisitos

- Go instalado no seu sistema.
- Um servidor MySQL em execução.
- GORM e o driver MySQL instalados:
  ```bash
  go get gorm.io/gorm
  go get gorm.io/driver/mysql
  ```

### Configurando o Banco de Dados MySQL

1. Crie um banco de dados MySQL.
2. Atualize a string de conexão no `server.go` com suas credenciais:
   ```go
   dsn := "user:password@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
   ```
   - **user**: Nome de usuário do MySQL.
   - **password**: Senha do usuário.
   - **127.0.0.1**: Endereço do servidor MySQL (pode ser localhost se estiver rodando na mesma máquina).
   - **3306**: Porta do MySQL (a porta padrão é `3306`).
   - **dbname**: Nome do banco de dados.

---

## Orientações para Executar o Projeto

### Passo 1: Iniciar o Servidor

1. Navegue até o diretório do projeto.
2. Execute o servidor:
   ```bash
   go run server.go
   ```
3. O servidor estará disponível na porta `8080`.

### Passo 2: Executar o Cliente

1. Em outra janela de terminal, execute o cliente:
   ```bash
   go run client.go
   ```
2. O cliente fará uma requisição ao servidor e salvará a cotação em um arquivo `cotacao.txt` no mesmo diretório.

---

## Estrutura do Arquivo `cotacao.txt`

O arquivo `cotacao.txt` será gerado com o seguinte formato:
```
Dólar: {valor}
```

---
