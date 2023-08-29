#!/bin/sh
cd /app
go mod tidy 
go run . &

air