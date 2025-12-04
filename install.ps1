# HajiLang Windows Kurulum Scripti
$ErrorActionPreference = "Stop"

Write-Host "HajiLang Kurulumu Baslatiliyor..." -ForegroundColor Green

# Mimari tespiti
$arch = if ([Environment]::Is64BitOperatingSystem) { "amd64" } else { "386" }
Write-Host "Mimari: windows-$arch" -ForegroundColor Yellow

# GitHub'dan indir
$repo = "slymanmrcan/hajilang"
$downloadUrl = "https://github.com/$repo/releases/latest/download/hajilang-windows-$arch.exe"
$tempFile = "$env:TEMP\hajilang.exe"

Write-Host "Indiriliyor..." -ForegroundColor Yellow
try {
    Invoke-WebRequest -Uri $downloadUrl -OutFile $tempFile
} catch {
    Write-Host "Indirme hatasi: $_" -ForegroundColor Red
    exit 1
}

# Kurulum dizini
$installDir = "$env:LOCALAPPDATA\hajilang"
New-Item -ItemType Directory -Force -Path $installDir | Out-Null

# Kopyala
Move-Item -Path $tempFile -Destination "$installDir\hajilang.exe" -Force

# PATH'e ekle
$currentPath = [Environment]::GetEnvironmentVariable("Path", "User")
if ($currentPath -notlike "*$installDir*") {
    Write-Host "PATH'e ekleniyor..." -ForegroundColor Yellow
    [Environment]::SetEnvironmentVariable(
        "Path",
        "$currentPath;$installDir",
        "User"
    )
    Write-Host "PATH guncellendi. Yeni terminalde calisir." -ForegroundColor Yellow
}

Write-Host "Kurulum tamamlandi!" -ForegroundColor Green
Write-Host "Kullanim: hajilang [dosya.haji]" -ForegroundColor Green
Write-Host "REPL icin: hajilang" -ForegroundColor Green
Write-Host "" 
Write-Host "NOT: Yeni bir PowerShell/CMD penceresi acin." -ForegroundColor Yellow