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

### `main.go`

The main server file. All the code is loaded and run from here.

### `<folder name>/<folder name>.go`

This is typically the main file of the package `<folder name>`.

### `db/`

The connection to the database. 

### `frontend/`

This is the place that we put the entire frontend. This has its own `README.md` and other documentation. Check those if needed.

### `model/`

This is the place where the models are placed.

### `resolvers/`

This is the place where the graphql resolvers live.

### `schema/`

This is the place where the graphql schema lives.

### `utils/`

This is the place where utility functions live. We currently use this just for JWT stuff.
