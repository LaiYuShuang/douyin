# douyin
## Introduction
The project implements a simple tiktok service, divided into four main sections.

| Service Name    |  Usage    | Framework    | protocol    | Path                   | IDL                                      |
| --------------- | ------------ | ---------- | -------- | ---------------------- | ----------------------------------------- |
| demoapi         | http interface | kitex/gin  | http     | douyin/cmd/api  |                                           |
| demouser | user data management | kitex/gorm | thrift | douyin/cmd/user |        douyin/idl/user.thrift  |
| demovideo | video data management | kitex/gorm | thrift   | douyin/cmd/video |         doutin/idl/video.thrift |
| demofavorite | favorite data management | kitex/gorm | thrift   | douyin/cmd/favorite |         doutin/idl/favorite.thrift |
## catalog introduce
| catalog       | introduce      |
| ---------- | ---------------- |
| pkg/constants   | constant        |
| pkg/bound    |  customized bound handler    |
| pkg/errno      | customized error number |
| pkg/middleware | RPC middleware     |
| pkg/tracer  | init jaeger     |
| sql         | create tables | 
| dal   | db operation              |
| pack       | data pack         |
| service    | business logic   |
## Quick Start
### 1.Setup Basic Dependence
```shell
docker-compose up
```
### 2.Run User RPC Server
```shell
cd cmd/user
sh build.sh
sh output/bootstrap.sh
```
### 3.Run Video RPC Server
```shell
cd cmd/video
sh build.sh
sh output/bootstrap.sh
```
### 4.Run Favorite RPC Server
```shell
cd cmd/favorite
sh build.sh
sh output/bootstrap.sh
```
### 5.Run API Server
```shell
cd cmd/api
sh build.sh
```
