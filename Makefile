install:
	sudo chmod -R u+x ./scripts/bash
	@ ./scripts/bash/install.sh
	@ ./scripts/bash/prepare-githooks.sh

lint: 			# lint all the codebase
	golangci-lint run

lint-staged: 	# lint only the staged files
	golangci-lint run --new --new-from-rev=HEAD

debug:
	go run cmd/app/main.go

build:
	go build cmd/app/main.go

test:
	go test ./...
