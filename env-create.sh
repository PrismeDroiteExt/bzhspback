#!/bin/bash

# Prompt for database details
read -p "Please enter your Database root password: " DB_ROOT_PASSWORD
read -p "Please enter your Database username: " DB_USER
read -p "Please enter your Database user password: " DB_USER_PASSWORD
read -p "Please enter your Database name: " DB_NAME
read -p "Please enter the name of your Database container name: " DB_CONTAINER_NAME

read -p "What Timezone to use? (default Europe/Paris) " TIMEZONE
if [ -z "${TIMEZONE}" ]
  then
    TIMEZONE="'Europe/Paris'"
fi

read -p "What DB SSLMODE to use? (default disable) " DB_SSLMODE
if [ -z "${DB_SSLMODE}" ]
    then
    DB_SSLMODE="'disable'"
    else
    if [ "${DB_SSLMODE}" == "disable" ] || [ "${DB_SSLMODE}" == "enable" ]
        then
        DB_SSLMODE="'${DB_SSLMODE}'"
        else
        echo "Invalid DB SSLMODE. Please enter either the word disable or enable."
        exit 1
    fi
fi

read -p "What is the Database exposed port? (default 5432) " DB_PORT
if [ -z "${DB_PORT}" ]
    then
    DB_PORT=5432
    else
    # Validate that DB_PORT is a number between 1024 and 65535
    if ! [[ "${DB_PORT}" =~ ^[0-9]+$ ]] || [ "${DB_PORT}" -lt 1024 ] || [ "${DB_PORT}" -gt 65535 ]; then
        echo "Invalid port number. Please enter a number between 1024 and 65535."
        exit 1
    fi
fi

# Create .env file
cat > .env <<EOL
DB_ROOT_PASSWORD=${DB_ROOT_PASSWORD}
DB_HOST=${DB_CONTAINER_NAME}
DB_USER=${DB_USER}
DB_PASSWORD=${DB_USER_PASSWORD}
HOSTNAME=${HOSTNAME}
DB_NAME=${DB_NAME}
TIMEZONE=${TIMEZONE}
DB_SSLMODE=${DB_SSLMODE}
DB_PORT=${DB_PORT}
EOL