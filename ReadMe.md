## CRUD APP WITH GO (Guide on how to start)

### Dependencies

- [Go](https://go.dev/)
- [PostgreSQL](https://www.postgresql.org/)
- [GORM](https://gorm.io/)
- [Gin](https://gin-gonic.com/)
- [gqlgen](https://gqlgen.com/)

Install required packages

```bash
go mod download
```

Run server with compile daemon package for hot-reload

```bash
CompileDaemon -command="./allergymanagement-api"
```

> Note this all requires database and `.env` file to work so

- Rename .env.example --> .env and include your db credentials

### Folder structure

```
.
└── ALLERGYMANAGEMENT-API-GIN/
    ├── configs/
    │   ├── env.go
    ├── db/
    │   ├── db.go
    ├── graph/
    │   ├── model (auto generated model by gqlgen)
    │   ├── *.schema.graphqls (schemas of graphqls)
    │   ├── *.schema.resolvers.go (resolvers for graphqls schemas)
```
