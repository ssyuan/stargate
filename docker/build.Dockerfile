FROM rust:latest

RUN  apt-get update && \
  apt-get install -y protobuf-compiler cmake && \
  apt-get install --no-install-recommends -y nano net-tools tcpdump iproute2 netcat ngrep atop gdb strace && \
  apt-get clean && \
  rm -r /var/lib/apt/lists/*
  
WORKDIR /starcoin
COPY rust-toolchain /starcoin/rust-toolchain
RUN rustup install $(cat rust-toolchain)
