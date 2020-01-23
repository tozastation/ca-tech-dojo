# CA TECH DOJO (Golang)
## Description
|KEY|VALUE|
|---|---|
|パッケージ構成|golang-standards/project-layout|
|ソフトウェア設計|クリーンアーキテクチャ|
|ORM|GORM|
|HTTP-SERVER|gin-gonic|
## Endpoint
|URL|METHOD|DESCRIPTION|
|---|---|---|
|/v1/user/create|POST|ユーザ情報作成|
|/v1/user/get|GET|ユーザ情報取得|
|/v1/user/update|PUT|ユーザ情報更新|
## Requirement
- Docker
- Docker-Compose
## Install
### DEV MODE (Develop by Hot Reload)
- Run This Command: `make up-dev`