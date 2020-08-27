FROM golang:1.15

WORKDIR /api
# server以下をコピーする
COPY . .
ENV GO111MODULE=on

# APIをホットリロードで起動する
RUN go get github.com/pilu/fresh
CMD ["fresh"]