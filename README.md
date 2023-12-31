# Simple Bank

## Requirements
Have docker installed

Have [Go Migrate](https://github.com/golang-migrate/migrate) installed

Have [SQLC](https://sqlc.dev/) installed

## First Steps

### Create database
- `make postgres`
- `make createdb`
- `make migrateup`


## Database

Database schema created with: [dbdiagram](https://dbdiagram.io/home)

Database migrations managed with: [Go Migrate](https://github.com/golang-migrate/migrate)

### Structure
![Database Schema](simpl-bank-db-schema.svg)
