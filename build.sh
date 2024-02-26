#!/bin/bash

# Check OS and architecture
os=$(uname -s | tr '[:upper:]' '[:lower:]')
arch=$(uname -m)

# Map architecture to GOARCH
case $arch in
    "x86_64")
        goarch="amd64"
        ;;
    "i386" | "i686")
        goarch="386"
        ;;
    "aarch64")
        goarch="arm64"
        ;;
    "arm64")
        goarch="arm64"
        ;;
    *)
        echo "Unsupported architecture: $arch"
        exit 1
        ;;
esac

# Check if Go is installed
if ! command -v go &> /dev/null

then
    echo "Go could not be found, installing..."
    wget https://dl.google.com/go/go1.18.5.$os-$arch.tar.gz -O go.tar.gz
    sudo tar -C /usr/local -xzf go.tar.gz
    rm go.tar.gz
    echo "export PATH=$PATH:/usr/local/go/bin" >> ~/.bashrc
    source ~/.bashrc
fi


# Check if the bin/OS_ARCH directory exists, and if not, create it
if [ ! -d "bin/$os"_"$goarch" ]; then
    mkdir -p bin/$os"_"$goarch
fi
# Build main.go
GOOS=$os GOARCH=$goarch go build -o bin/$os"_"$goarch/eps_emulator main.go