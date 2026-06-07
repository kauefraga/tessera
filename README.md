# Tessera

Téssera é um bot feito para o Discord com o objetivo de escolher um jogo aleatoriamente entre vários para um grupo.

## Começando

### Convidar o bot

Você pode adicionar o Tessera diretamente ao seu servidor clicando no link abaixo:

[Convidar o Tessera](https://discord.com/oauth2/authorize?client_id=1354275311600140338)

```txt
https://discord.com/oauth2/authorize?client_id=1354275311600140338
```

## Comandos

### `/pick-game`

Escolhe um jogo aleatoriamente entre os informados.

| Opção   | Descrição                              | Obrigatório |
|---------|----------------------------------------|-------------|
| `jogos` | Nomes dos jogos, separados por vírgula | Sim         |

**Exemplo:**

```txt
/pick-game jogos: Valorant,CS2,Minecraft
```

## Rodar na sua própria máquina

#### Pré-requisitos

- [Go 1.25+](https://go.dev/dl/) ou [Docker](https://www.docker.com/)
- As variáveis de ambiente configuradas (veja abaixo)

#### Variáveis de ambiente

Crie um arquivo `.env` na raiz do projeto com as seguintes variáveis:

```env
DISCORD_TOKEN=seu_token_aqui
APPLICATION_ID=seu_application_id_aqui
SERVER_ID=seu_server_id_aqui
```

| Variável         | Descrição                                              |
|------------------|--------------------------------------------------------|
| `DISCORD_TOKEN`  | Token do bot obtido no [Discord Developer Portal](https://discord.com/developers/applications) |
| `APPLICATION_ID` | ID da aplicação do bot no Discord Developer Portal     |
| `SERVER_ID`      | ID do servidor Discord onde o bot será utilizado       |

#### Com Go

```sh
go run cmd/main.go
```

#### Com Docker

```sh
# Build da imagem
docker build -t tessera .

# Rodar o container
docker run --rm --env-file .env tessera
```

## Licença

Este projeto está sob a licença MIT - Veja a [LICENÇA](LICENSE) para mais informações.
