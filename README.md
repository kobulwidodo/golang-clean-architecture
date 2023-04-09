# Golang Clean Architecture RestAPI

## Prerequsites

- Go v1.18
- MySQL
- Docker & Docker Compose

## How to Prepare the Environment

Run this command line to prepare the environment, it will create `.env` file.

```shell
echo "export JWT_KEY=typeyourjwtsecrethere" >> .env
```

the `.env` file will looks like this:

```shell
export JWT_KEY=typeyourjwtsecrethere
```

Run this command line to export to your environment :

```shell
source .env
```

Run this command to duplicate config file then you must fill `config.json` :

```shell
cd etc/cfg
cp config.json.template config.json
```

Run this command line to create database using docker compose :

```shell
cd env
docker-compose up -d
```

Run this command line to install swagger :

```shell
make swag-install
```

## How to Run the Application

Start the application by running:

```shell
make run-app
```
