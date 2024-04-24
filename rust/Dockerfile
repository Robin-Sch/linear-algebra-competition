FROM rust:1.77

WORKDIR /app

COPY . .

RUN cargo install --path .

CMD /app/target/release/one && /app/target/release/two && /app/target/release/three