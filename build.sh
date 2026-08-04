#!/bin/bash

# Compilar para Windows
GOOS=windows GOARCH=amd64 go build -ldflags "-s -w" -o MeuStealer.exe

# Compilar para Linux
GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o MeuStealer

echo "[+] Compilação concluída!"
