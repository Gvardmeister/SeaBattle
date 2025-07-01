@echo off
echo === Building SeaBattle for all platforms ===

:: Создание папки build, если не существует
if not exist build mkdir build

:: Сборка под Windows
echo --- Building for Windows (amd64)...
set GOOS=windows
set GOARCH=amd64
go build -o build\seabattle_windows.exe cmd\main.go

:: Сборка под Linux
echo --- Building for Linux (amd64)...
set GOOS=linux
set GOARCH=amd64
go build -o build\seabattle_linux cmd\main.go

:: Сборка под macOS
echo --- Building for macOS (amd64)...
set GOOS=darwin
set GOARCH=amd64
go build -o build\seabattle_macos cmd\main.go

:: Очистка переменных окружения
set GOOS=
set GOARCH=

echo === Build completed! Files saved in /build ===
pause
