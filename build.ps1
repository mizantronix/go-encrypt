$ErrorActionPreference = "Stop"

$targets = @(
    @{ GOOS = "windows"; GOARCH = "amd64"; OUT = "./build/go-encrypt-windows-amd64.exe" },
    @{ GOOS = "windows"; GOARCH = "386";   OUT = "./build/go-encrypt-windows-386.exe" },
    @{ GOOS = "windows"; GOARCH = "arm64"; OUT = "./build/go-encrypt-windows-arm64.exe" },

    @{ GOOS = "linux";   GOARCH = "amd64"; OUT = "./build/go-encrypt-linux-amd64" },
    @{ GOOS = "linux";   GOARCH = "386";   OUT = "./build/go-encrypt-linux-386" },
    @{ GOOS = "linux";   GOARCH = "arm64"; OUT = "./build/go-encrypt-linux-arm64" },

    @{ GOOS = "darwin";  GOARCH = "amd64"; OUT = "./build/go-encrypt-macos-amd64" },
    @{ GOOS = "darwin";  GOARCH = "arm64"; OUT = "./build/go-encrypt-macos-arm64" }
)

foreach ($target in $targets) {
    $env:GOOS = $target.GOOS
    $env:GOARCH = $target.GOARCH
    $outFile = $target.OUT

    Write-Host "Building for $($env:GOOS)/$($env:GOARCH)... -> $outFile"
    go build -o $outFile
}

Write-Host "Done"
