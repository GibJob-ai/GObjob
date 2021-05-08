# GObjob

**Assume everything <MAIN> or root directory**

## Setup

### Creating the config
Create a copy of `config.example.hcl` named `config.hcl`

``` sh
$ cp config.example.hcl config.hcl
```

Then change the variables to whats needed

## Run app locally
* TODO: Once docker-compose is working for local development / testing, put command here

## Run backend (must have built frontend locally, see frontend/README.md)
```bash
cd backend
go run main.go
```

### Deploying to Prod
We currently deploy through CI only. This is triggered by merging a PR into the `prod` branch.

## Graphql 

On dev mode you can GET at `/graphql` to get graphql playground

## Architecture

This can be read at `ARCHITECTURE.md`

## Migrate
The server migrates automatically on boot

If you wish to only migrate and not serve then run `go run main.go --migrate-only`
