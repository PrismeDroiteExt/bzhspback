# Project structure

This is the structure of the project:

```txt
app/
│
├── cmd/
│   └── web/
│       └── main.go         # Entrypoint for the application
│
├── config/
│   └── config.yaml         # Configuration file (e.g., DB connection details)
│
├── docs/                   # Documentation files (e.g., Swagger)
|
├── internal/               # Private application and library code (we do not want this exposed)
│   ├── api/              
│   │   └── routes.go       # API route definitions
│   ├── controllers/
│   ├── middleware/
│   ├── models/
│   └── repository/         # Database operations (a file per entity)
|
├── scripts/             
│   └── dev-entrypoint.sh   # Build automation script          
│
├── static/                 # Static files (e.g., images)
│
├── go.mod
└── go.sum
```