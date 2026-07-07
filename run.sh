#!/bin/bash
go run $(ls $(ls *.go | grep -v _test.go))
