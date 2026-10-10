# Forward arguments unchanged, including Go flags such as -o.
$ErrorActionPreference = 'Stop'
$projectRoot = Split-Path -Parent $PSScriptRoot
$goExecutable = Join-Path $projectRoot '.tools/go/bin/go.exe'
if (-not (Test-Path -LiteralPath $goExecutable)) {
    $goExecutable = (Get-Command go -ErrorAction Stop).Source
}

# Keep downloaded modules and build caches inside the workspace.
$env:GOPATH = Join-Path $projectRoot '.cache/go'
$env:GOCACHE = Join-Path $projectRoot '.cache/go-build'
Push-Location (Join-Path $projectRoot 'backend')
try {
    & $goExecutable @args
    $goExitCode = $LASTEXITCODE
} finally {
    Pop-Location
}
exit $goExitCode
