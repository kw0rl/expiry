# Expiry

A learning project for tracking what expires and when to take action.

## Current foundation

- `frontend/`: Vue 3, TypeScript, and Vite.
- `backend/cmd/api/main.go`: Go + Gin server with `GET /api/health`.
- `scripts/go.ps1`: runs Go from the backend directory, using the workspace-local toolchain if present or Go from PATH otherwise.

The database and authentication will be added in later learning phases.

## Run locally (Windows PowerShell)

Open two terminals at the repository root.

Backend:

```powershell
powershell -NoProfile -ExecutionPolicy Bypass -File .\scripts\go.ps1 run ./cmd/api
```

Frontend:

```powershell
cd frontend
npm.cmd install
npm.cmd run dev
```

Open http://127.0.0.1:5173. The connection check should succeed while both servers are running. Stop each server with Ctrl+C.

The API listens on `127.0.0.1:8081`. Its health response is:

```json
{"status":"ok","service":"expiry-api"}
```

Vue calls `/api/health`; Vite forwards it to Go using the proxy in `frontend/vite.config.ts`. The Gin handler returns JSON, and Vue displays the result. There is no database operation yet. This health check only confirms that the API responds.

`API_ADDR` can override the backend listening address. If you change its port, update the Vite proxy target too. The development proxy is not included in production builds: deployment will need to route `/api` requests to Go separately.

## Checks

From `frontend/`:

```powershell
npm.cmd run typecheck
npm.cmd run build
```

From the repository root:

```powershell
powershell -NoProfile -ExecutionPolicy Bypass -File .\scripts\go.ps1 fmt ./...
powershell -NoProfile -ExecutionPolicy Bypass -File .\scripts\go.ps1 vet ./...
powershell -NoProfile -ExecutionPolicy Bypass -File .\scripts\go.ps1 test ./...
powershell -NoProfile -ExecutionPolicy Bypass -File .\scripts\go.ps1 build -o bin/expiry-api.exe ./cmd/api
```

No lint tool or automated frontend tests are configured yet.

## Tooling and dependencies

Use a Node.js version supported by Vite (this setup was created with Node 22.14.0). Install frontend dependencies with `npm.cmd ci` when reproducing the lockfile.

Go is downloaded locally into `.tools/go/` for this workspace and is not committed. On another machine, install the Go version required by `backend/go.mod` or newer. `go.mod` and `go.sum` record backend dependencies; Go downloads them on the first run.

`.tools/`, `.cache/`, build output, and local environment files are ignored by Git. Never put secrets into frontend environment variables: frontend code is delivered to the browser.
