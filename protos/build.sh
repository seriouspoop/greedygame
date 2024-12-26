#!/bin/bash

REPO_DIR="$PWD"

go install google.golang.org/protobuf/cmd/protoc-gen-go
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

protoc_version=$(protoc --version | grep -oE '[0-9]+\.[0-9]+' | head -n1)
required_version="25.3"

if [[ "$(printf '%s\n' "$required_version" "$protoc_version" | sort -V | head -n1)" != "$required_version" ]]; then
    echo "Error: protoc version $required_version or higher is required"
    #brew install protobuf
    exit 1
fi

for service_dir in proto/*pb; do
  service=${service_dir%*/}  # remove trailing "/"
  service=${service/proto\//} # remove the word proto

  # Generate Go directories
  mkdir -p go/$service

  # Check if the proto file contains the line
  if grep -q 'import "google/api/annotations.proto";' "$service_dir"/*.proto; then
    
      # Generate go, grpc, gw files
    protoc -I=proto $service_dir/*.proto --go_out=paths=source_relative:go \
      --go-grpc_out=paths=source_relative:go \
      --grpc-gateway_out=paths=source_relative:go

  else

    # Generate grpc and go files
    protoc -I=proto $service_dir/*.proto --go_out=paths=source_relative:go \
      --go-grpc_out=paths=source_relative:go
      protos/build.sh
  fi
done