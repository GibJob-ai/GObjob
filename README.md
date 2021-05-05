# GObjob

**Assume everything <MAIN> or root directory**

## Setup

### Creating the config
Create a copy of `config.example.hcl` named `config.hcl`

``` sh
$ cp config.example.hcl config.hcl
```

Then change the variables to whats needed

## Run

run `go run main.go`

### Run in prod

To run for prod, set `Mode="prod"` inside `config.hcl`
then run `sudo go run main.go`

## Graphql 

On dev mode you can GET at `/graphql` to get graphql playground

## Architecture

This can be read at `ARCHITECTURE.md`

## Migrate
The server migrates automatically on boot

If you wish to only migrate and not serve then run `go run main.go --migrate-only`
