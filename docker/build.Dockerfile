FROM rust:latest

RUN sudo apt-get update && \
 sudo apt-get install -y protobuf-compiler cmake && \
 sudo apt-get clean && \
 sudo rm -r /var/lib/apt/lists/*
