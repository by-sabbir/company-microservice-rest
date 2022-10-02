.PHONY:

deps:
	go mod tidy

test:
	go test -coverprofile cover.out ./... -v -cover

coverage-report-html:
	go tool cover -html=cover.out
