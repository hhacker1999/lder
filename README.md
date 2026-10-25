# lder

Movie and TV discovery with a Go API and embedded React UI.

## Development

1. Create `config.json` with `port`, `tmdb_key`, `tvdb_key`, and `database`.
2. Start the API: `go run .`
3. Start the UI (proxies `/api` to port 6869): `make ui-dev`

Open http://localhost:5173

## Production

```bash
make build
./bin/lder
```

The binary embeds `web/dist` and serves the SPA on the configured port.
