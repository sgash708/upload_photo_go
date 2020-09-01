# 開発環境
FROM node:14-alpine as build-stage
WORKDIR /app
# client以下をコピーする
COPY . .
RUN yarn install
