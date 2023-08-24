#!/bin/bash
test-service:
	@go test  -cover ./internal/service

test-cover:
	@go test $$(go list ./internal/service  | grep -v /vendor/) -coverprofile=cover.out && go tool cover -html=cover.out ; rm -f cover.out

start:
	@go run main.go

