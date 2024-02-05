@echo off
cd /d "%~dp0"

cd .\nats-streaming-server && (
    start cmd /k "go run nats-streaming-server.go"
    timeout /t 15 /nobreak
)
