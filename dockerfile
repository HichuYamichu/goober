FROM golang:latest AS go-build

WORKDIR /build

ENV GO111MODULE=on

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o uploader ./main.go

FROM node:latest AS node-build

WORKDIR /build

COPY ./web/package.json ./web/yarn.lock ./

COPY ./web .

RUN yarn && yarn build

FROM scratch

WORKDIR /uploder

COPY --from=go-build /build/uploader /uploder/
COPY --from=node-build /build/public /uploder/web/public

EXPOSE 9000

CMD ["./uploader", "start"]