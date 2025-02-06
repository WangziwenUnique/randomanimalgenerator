#!/bin/bash

# 设置错误时退出
set -e

echo "开始构建项目..."

# 前端构建
echo "开始构建前端..."
cd frontend
npm install
npm run build
cd ..

# 后端构建
echo "开始构建后端..."
cd backend
# 设置 Go 环境为 linux
GOOS=linux GOARCH=amd64 go build -o app

echo "构建完成！"
echo "所有文件已打包到 dist 目录" 