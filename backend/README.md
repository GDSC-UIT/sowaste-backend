# **Sowaste Backend**

## Folder structure

`backend` - root folder

- `ai`
  - Updating...
- `go`
  - `api`
  - `build`
  - `cmd`
  - `configs`
  - `deployments`
  - `docs`
  - `internal`
    - `app` - The point where all our dependencies and logic are collected and run the app. The run method that is called from /cmd.
    - `config` - Initialization of the general app configurations that we wrote in the root of the project.
    - `database` - The files contain methods for interacting with databases.
    - `models` - The structures of database tables.
    - `services` - The entire business logic of the application.
    - `transport` - Here we store http-server settings, handlers, ports, etc.
  - `migrations` - This contains all migrations related to databases, e.g. SQL files.
  - `pkg`
  - `go.mod`
  - `README.md`
