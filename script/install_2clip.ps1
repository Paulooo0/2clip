# Define a versão
$version = "1.3.0"

# Detectar arquitetura
$arch = $env:PROCESSOR_ARCHITECTURE
switch ($arch) {
    "AMD64" { $arch = "amd64" }
    "ARM64" { $arch = "arm64" }
    "x86"   { $arch = "386" }
    default {
        Write-Error "Unsupported architecture: $arch"
        exit 1
    }
}

# Definir nome do binário
$binary = "2clip-$version-windows-$arch.exe"
$url = "https://github.com/Paulooo0/2clip/releases/download/v$version/$binary"

# Download do binário
$tempPath = "$env:TEMP\$binary"
Write-Host "Downloading $binary from $url"
Invoke-WebRequest -Uri $url -OutFile $tempPath

# Caminho de destino (pasta no PATH, por ex: C:\Program Files\2clip)
$destinationDir = "$env:ProgramFiles\2clip"
if (!(Test-Path $destinationDir)) {
    New-Item -ItemType Directory -Path $destinationDir | Out-Null
}

# Mover o binário
Move-Item -Path $tempPath -Destination "$destinationDir\2clip.exe" -Force

# Adicionar ao PATH (se ainda não estiver)
$path = [System.Environment]::GetEnvironmentVariable("Path", [System.EnvironmentVariableTarget]::Machine)
if ($path -notlike "*$destinationDir*") {
    [System.Environment]::SetEnvironmentVariable("Path", "$path;$destinationDir", [System.EnvironmentVariableTarget]::Machine)
    Write-Host "Added $destinationDir to system PATH. You may need to restart your shell."
}

Write-Host "2clip v$version installed successfully!"
