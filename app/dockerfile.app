FROM golang:1.24.7 AS builder

WORKDIR /app

COPY main/ ./main
COPY server/ ./server
COPY src/ ./src
COPY db/ ./db

WORKDIR /app/main

RUN go mod download

RUN mkdir -p /app/bin

RUN go build -o /app/bin/main .

WORKDIR /app

# Base: Debian slim
FROM debian:bookworm-slim

# Evita interações durante a instalação
ENV DEBIAN_FRONTEND=noninteractive
ENV FIRESIMPLE_DEFAULT_USER_DB=test
ENV FIRESIMPLE_DEFAULT_PASSWORD_DB=test

# Atualiza o sistema e instala dependências básicas
RUN apt-get update && apt-get install -y \
    sudo wget gnupg lsb-release curl ca-certificates \
    && apt-get clean && rm -rf /var/lib/apt/lists/*

# Adiciona chave e repositório oficial do PostgreSQL 18
RUN wget -qO - https://www.postgresql.org/media/keys/ACCC4CF8.asc | apt-key add - \
    && echo "deb http://apt.postgresql.org/pub/repos/apt $(lsb_release -cs)-pgdg main" > /etc/apt/sources.list.d/pgdg.list

# Atualiza os repositórios e instala PostgreSQL 18
RUN apt-get update && apt-get install -y \
    postgresql-18 postgresql-client-18 libpq-dev \
    && apt-get clean && rm -rf /var/lib/apt/lists/*

# Cria diretório para Unix socket do PostgreSQL e dá permissões
RUN mkdir -p /var/run/postgresql && chown -R postgres:postgres /var/run/postgresql

# Copia o entrypoint e libera execução
COPY entrypoint.sh /entrypoint.sh
RUN chmod +x /entrypoint.sh

COPY --from=builder /app/bin/main /usr/local/bin/fire_simple
COPY --from=builder /app/db/tables /usr/local/fire_simple/tables

# Define entrypoint
ENTRYPOINT ["/entrypoint.sh"]


