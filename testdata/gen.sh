#!/bin/bash

protoc \
-I=. \
-I=./testdata2 \
--descriptor_set_out=testdata.pb.desc \
--go_out=paths=source_relative:. \
*.proto testdata2/*.proto