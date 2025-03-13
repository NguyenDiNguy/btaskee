#!/bin/bash

# Thư mục chứa các file .pb.go
MODEL_DIR="./model"

# Duyệt qua tất cả các file .pb.go trong thư mục
for file in "$MODEL_DIR"/*/*.pb.go; do
    echo "Injecting tags into $file"
    protoc-go-inject-tag -remove_tag_comment -input="$file"
done