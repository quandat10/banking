# BANKING GOLANG

## Description

Banking money from first account to another account

### Main Feature

- [🍀] Connect DB (Postgres, Redis)
- [🍀] Unit test + TDD
- [🍀] Mock DB (gomock)
- [🍀] CI github action
- [🦊] Http request + gRPC
- [💥] Verify token + Authorization (PASETO & JWT)
- [💥] Deployment (Docker & Kubernetes)
- [💥] Document API (Swagger)

### Environment

1. Golang (go1.18.1)
2. Docker (20.10.22)
3. Sqlc (v1.16.0) [ <a href="https://sqlc.dev/">Doc</a> ]
4. Make (4.3)
5. Ubuntu (22.04LTS)

## Run Application

From your download repository location

```bash
$ cd bank
```

Start postgres database with port 5432

```bash
$ make postgres
```

Create simple_bank database

```bash
$ make create-db
```

Migration DB

```bash
$ make migrate-up
```

### Test Entry

```bash
$ make test
```

### Start Server

```bash
$ make server
```

After server is running, It will listen on port 3000

### Database Tables

1. Accounts
2. Entries
3. Transfers

🌈<i>You should reference to Makefile and app.env files to easy setup project