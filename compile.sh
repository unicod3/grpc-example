#bin/bash

protoc --go_out=plugins=grpc:. pcalc/calc.proto
