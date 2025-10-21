#!/bin/bash
set -e

PGDATA=/var/lib/postgresql/data/pgdata
PG_BIN=/usr/lib/postgresql/18/bin

# Inicializa o banco caso ainda não exista
if [ ! -d "$PGDATA" ]; then
    mkdir -p "$PGDATA"
    chown -R postgres:postgres /var/lib/postgresql/data
    su - postgres -c "$PG_BIN/initdb -D $PGDATA"
fi

# Inicia o servidor em background
echo "🚀 Iniciando servidor temporário..."
su - postgres -c "$PG_BIN/pg_ctl -D $PGDATA -o '-c listen_addresses=' -w start"

NEW_USER_DB=${FIRESIMPLE_DEFAULT_USER_DB}
NEW_PASSWORD_DB=${FIRESIMPLE_DEFAULT_PASSWORD_DB}

if [ -z "${NEW_USER_DB+x}" ]; then
    echo "❌ FIRESIMPLE_DEFAULT_USER_DB não existe"
    exit 1
elif [ -z "$NEW_USER_DB" ]; then
    echo "⚠️ FIRESIMPLE_DEFAULT_USER_DB existe mas está vazia"
    NEW_USER_DB="firesimples_default"
else
    echo "✅ FIRESIMPLE_DEFAULT_USER_DB = $NEW_USER_DB"
fi

if [ -z "${NEW_PASSWORD_DB+x}" ]; then
    echo "❌ FIRESIMPLE_DEFAULT_PASSWORD_DB não existe"
    exit 1
elif [ -z "$NEW_PASSWORD_DB" ]; then
    echo "⚠️ FIRESIMPLE_DEFAULT_PASSWORD_DB existe mas está vazia"
    NEW_PASSWORD_DB="firesimples_default"
else
    echo "✅ FIRESIMPLE_DEFAULT_PASSWORD_DB = $NEW_PASSWORD_DB"
fi

echo "Criando novo usuário $NEW_USER_DB..."
su - postgres -c "$PG_BIN/createuser --superuser $NEW_USER_DB"
su - postgres -c "$PG_BIN/psql -d postgres -c \"ALTER USER $NEW_USER_DB WITH PASSWORD '$NEW_PASSWORD_DB';\""

su - postgres -c "psql -c 'ALTER ROLE $NEW_USER_DB WITH REPLICATION BYPASSRLS;'"

echo "Parando servidor temporário..."
su - postgres -c "$PG_BIN/pg_ctl -D $PGDATA -m fast stop"

echo "Iniciando PostgreSQL em modo foreground..."
su - postgres -c "$PG_BIN/postgres -D $PGDATA"



