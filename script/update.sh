#!/bin/bash
PROTO_FILE_DIR="./NewProto"
COMMON_PROTO_FILE_DIR="./NewCommonProto"
PARENT_DIR=$(basename "$(dirname "$(pwd)")")
GRPC_DIR="$PARENT_DIR/grpc"
PROJECT_GRPC_DIR="../../$GRPC_DIR"
COMMON_PROTO=("nodeoverview.proto" "response.proto")
if [ ! -d "$COMMON_PROTO_FILE_DIR" ]; then
  mkdir -p "$COMMON_PROTO_FILE_DIR"
fi

for common_file in "${COMMON_PROTO[@]}"
do
    mv -f "$PROTO_FILE_DIR"/"$common_file" "$COMMON_PROTO_FILE_DIR" >/dev/null 2>&1
done

for common_file in "$COMMON_PROTO_FILE_DIR"/*; do
    if [ -f "$common_file" ]; then
      go_package_line="option go_package = \"$GRPC_DIR/common/proto/;common\";"
      sed -i '' "/^syntax = \"proto3\";/a\\
$go_package_line\\
" "$common_file"
    fi
done

mv -f "$COMMON_PROTO_FILE_DIR"/* "$PROJECT_GRPC_DIR"/common/proto

protoc \
  --proto_path="$PROJECT_GRPC_DIR"/ \
  --go_out="$PROJECT_GRPC_DIR"/ \
  --go_opt=paths=source_relative \
  --go-grpc_out="$PROJECT_GRPC_DIR"/\
  --go-grpc_opt=paths=source_relative \
  "$PROJECT_GRPC_DIR"/common/proto/*.proto

if [ -d "$PROTO_FILE_DIR" ]; then
    proto_files=($(find "$PROTO_FILE_DIR" -maxdepth 1 -name "*.proto"))
    num_files=${#proto_files[@]}
    echo "Number of proto files: $num_files"

    if [ "$num_files" -gt 0 ]; then
        for file in "${proto_files[@]}"; do
            filename=$(basename "$file" .proto)
            go_package_line="option go_package = \"$GRPC_DIR/$filename/proto/;$filename\";"
            sed -i '' "/^syntax = \"proto3\";/a\\
$go_package_line\\
" "$file"
            if [ ! -d "$PROJECT_GRPC_DIR/$filename" ]; then
                mkdir -p "$PROJECT_GRPC_DIR/$filename/proto"
                echo "Directory '$GRPC_DIR/$filename/proto' created."
                mkdir -p "$PROJECT_GRPC_DIR/$filename/request"
                echo "Directory '$GRPC_DIR/$filename/request' created."
            fi

            mv -f "$file" "$PROJECT_GRPC_DIR/$filename/proto/"
            protoc \
              --proto_path="$PROJECT_GRPC_DIR" \
              --proto_path="$PROJECT_GRPC_DIR"/common/proto/ \
              --go_out="$PROJECT_GRPC_DIR"/ \
              --go_opt=paths=source_relative \
              --go-grpc_out="$PROJECT_GRPC_DIR"/\
              --go-grpc_opt=paths=source_relative \
              "$PROJECT_GRPC_DIR"/"$filename"/proto/*.proto
        done
    fi
else
    echo "Directory $PROTO_FILE_DIR does not exist."
fi