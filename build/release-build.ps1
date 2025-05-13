param(
    [Parameter(Mandatory=$true)]
    [string]$version
)

if ([string]::IsNullOrWhiteSpace($version)) {
    Write-Error "Version number cannot be empty. Usage: ./release-build.ps1 <version>"
    exit 1
}

$projectName = "go-git-prompt"
$outputDir = "build"


if (-not (Test-Path -Path $outputDir -PathType Container)) {
    New-Item -Path $outputDir -ItemType Directory | Out-Null
}

# Windows build
Write-Host "Building for Windows x64 (version: $version)..."
$env:GOOS = "windows"
$env:GOARCH = "amd64"
$outputFileNameWindows = "$projectName-v$version-windows-amd64.exe"
go build -o "$outputDir/$outputFileNameWindows"
if ($LASTEXITCODE -ne 0) {
    Write-Error "Error building for Windows."
    exit $LASTEXITCODE
}
Write-Host "Windows build complete: $outputFileNameWindows"

# Linux build
Write-Host "Building for Linux x64 (version: $version)..."
$env:GOOS = "linux"
$env:GOARCH = "amd64"
$outputFileNameLinux = "$projectName-v$version-linux-amd64"
go build -o "$outputDir/$outputFileNameLinux"
if ($LASTEXITCODE -ne 0) {
    Write-Error "Error building for Linux."
    exit $LASTEXITCODE
}
Write-Host "Linux build complete: $outputFileNameLinux"

# Clean up GOOS and GOARCH
Remove-Item Env:\GOOS -ErrorAction SilentlyContinue
Remove-Item Env:\GOARCH -ErrorAction SilentlyContinue

Write-Host "All builds complete."
