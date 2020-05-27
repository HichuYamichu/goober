FROM golang:latest AS go-build

WORKDIR /build

ENV GO111MODULE=on

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o goober ./main.go

FROM node:latest AS node-build

WORKDIR /build

COPY ./web .

RUN npm i 

RUN npm run build

FROM alpine

WORKDIR /goober

COPY --from=go-build /build/goober /goober/
COPY --from=node-build /build/public/ /goober/web/public

EXPOSE 9000

CMD ["./goober", "start"]