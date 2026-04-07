#!/bin/bash

../bin/protoc --go_out=. --go_opt=paths=source_relative $1
