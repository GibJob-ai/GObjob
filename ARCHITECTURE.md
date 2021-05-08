# Architecture
**This document describes the high-level architecture of this project**

If you want to familiarize yourself with the code base and *generally* how it works, this is a good place to be.

## High Level TLDR
All the main code is loaded into main.go and ran from there.

## Code Map

#### Code Map Legend

`<file name>` for a file name

`<folder name>/` for a folder

`<folder name>/<file name>` for a file within a folder

### `<folder name>/<folder name>.go`

This is typically the main file of the package `<folder name>`.

### `config/`

The config files, currently WIP (needs to be integrated with secret management stuff for CI purposes)

### `config/config.example.hcl`

The example config that should be copied and name changed to `config.hcl`

### `config/config.hcl`

The config file to configure the server settings

### `frontend/`

This is the place that we put the entire frontend. This has its own `README.md` and other documentation. Check those if needed.

### `backend/`

All the backend code lives here.

### `backend/main.go`

The main server file. All the code is loaded and run from here.

### `backend/db/`

The connection to the database. 

### `backend/model/`

This is the place where the models are placed.

### `backend/migrations/`

Migrations directory, one go script for each new database revision.

### `backend/resolvers/`

This is the place where the graphql resolvers live.

### `backend/schema/`

This is the place where the graphql schema lives.

### `backend/utils/`

This is the place where utility functions live. 

### `backend/handler/`

This is the place for route handlers.

### `backend/config/`

This is the place that reads the config
