# Momentum
Smashing your goals through consistent actions



### Database setup

- install postgres and psql  locally on your device `brew install postgres`
- start the postgres database `brew services start postgres`
- create a database, a user and a password,
```
    psql postgres
    CREATE USER root WITH SUPERUSER PASSWORD=root;
    CREATE DATABASE task_manager;
```

### Vault Setup
- Install Hashicorp Vault locally on your device
- start a development instance of vault  using the command ` vault server -dev -dev-root-token-id="root" `
- Add the application secret to vault using the command `vault kv put -address=http://127.0.0.1:8200 -mount=secret task-manager @secret.json`
- The vault secret path should be `secret/data/task-manager`
- The vault root token should be `root`
- The vault address should be `http://127.0.0.1:8200`
- export the vault configurations in your terminal
```
    export VAULT_ADDR=http://127.0.0.1:8200
    export VAULT_TOKEN=root
    export VAULT_PATH=secret/data/task-manager

```

### Project setup


- Install Migrate - brew install golang-migrate
- Install Go - brew install go
- install all project dependencies ` go mod tidy`
- run `make start`


## Contribution  Guide
- for adding new feature, create a new  feature branch i.e `feat/add-oauth-signup` from the develop branch, after completing the feature, create a Pull Request to the develop branch
- for bug fixes, create a new  fix branch i.e `fix/add-oauth-signup` from the develop branch, after completing the fixes, create a Pull Request to the develop branch
