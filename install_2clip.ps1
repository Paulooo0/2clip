# Ensure script exits if any command fails
$ErrorActionPreference = "Stop"

# Compile the Go tool
Write-Host "Compiling 2clip..."
go build -o 2clip.exe cmd/2clip/main.go

# Define the installation path
$installPath = "C:\Program Files\2clip"

# Create the installation directory if it doesn't exist
if (-Not (Test-Path $installPath)) {
    Write-Host "Creating installation directory..."
    New-Item -ItemType Directory -Path $installPath
}

# Move the compiled binary to the installation path
Write-Host "Moving 2clip to $installPath..."
Move-Item -Path "2clip.exe" -Destination "$installPath\2clip.exe"

# Add the installation path to the system PATH if not already present
$path = [System.Environment]::GetEnvironmentVariable("Path", [System.EnvironmentVariableTarget]::Machine)
if (-Not ($path -contains $installPath)) {
    Write-Host "Adding $installPath to the system PATH..."
    [System.Environment]::SetEnvironmentVariable("Path", "$path;$installPath", [System.EnvironmentVariableTarget]::Machine)
}

Write-Host "2clip installation complete. You can now use '2clip' from anywhere in your command prompt."