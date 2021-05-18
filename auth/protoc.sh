#!/bin/bash
# PROTO_PATH=./crawl.proto
# DST_DIR=crawler
protoc --go_out=plugins=grpc:crawler crawl.proto 