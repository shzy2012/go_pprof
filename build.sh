#!/bin/bash

[ ! -d ./pprof ] && mkdir -p ./pprof

go tool pprof -png http://localhost:9090/debug/pprof/goroutine > ./pprof/goroutine.png
go tool pprof -png http://localhost:9090/debug/pprof/heap > ./pprof/heap.png
go tool pprof --seconds 25 -png http://localhost:9090/debug/pprof/profile > ./pprof/cpu.png

# go-torch -u http://localhost:9090 -t 30 