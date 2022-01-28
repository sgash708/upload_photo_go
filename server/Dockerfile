FROM golang:1.15

WORKDIR /api
# server以下をコピーする
COPY . .
ENV GO111MODULE=on

# DO: APIをホットリロードで起動
RUN go get github.com/pilu/fresh
CMD ["fresh"]