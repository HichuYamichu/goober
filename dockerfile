FROM rustlang/rust:nightly
WORKDIR /usr/src/uploader
COPY . .

RUN cargo install --path .

CMD ["uploader"]