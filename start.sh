#!/bin/bash

REPO_URL="https://github.com/Diverstt/Yandex_Go.git"  # URL репозитория
PROJECT_DIR="Yandex_Go"  # Имя директории проекта

git clone $REPO_URL $PROJECT_DIR

cd $PROJECT_DIR

go mod download

go run cmd/main.go