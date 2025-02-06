#!/bin/bash
echo "[#] Generating code for gRPC..."
protoc --go_out=../ --go-grpc_out=../ TokenServiceClient.proto
protoc --go_out=../ --go-grpc_out=../ TokenServiceServer.proto
echo "[#] Done!"