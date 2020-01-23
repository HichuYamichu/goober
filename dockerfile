FROM rustlang/rust:nightly AS builder
WORKDIR /usr/src/uploader
COPY . .
RUN cargo install --path .

FROM debian:buster-slim
COPY --from=builder /usr/local/cargo/bin/uploader /usr/local/bin/uploader

CMD ["uploader"]