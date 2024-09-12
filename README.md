# GoTH-Starter

This repo should be a starter project for a GoTH stack application.

for `Go` version - 1.23.0

### Dependencies
- `go install github.com/air-verse/air@latest`
- `go get github.com/a-h/templ/cmd/templ@latest`
- `go get -u github.com/gin-gonic/gin`



# Development Mode & Tips.

 - Open 3 Terminals, run the following commands.
    - `air` - Air will watch for changes overall, and reload server on changes you define in `.air.toml`

    - `templ generate --watch` - this will watch for changes in `.templ` files, and will regenerate them on demand.

    - `npm run watch` - This will make tailwind to be watched, and will create / remove css on demand.

