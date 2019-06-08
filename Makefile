JOB ?= build

.PHONY: build
build: test
	@ go build

.PHONY: test
test:
	@ go vet ./...
	@ richgo test -v -cover ./...

.PHONY: coverage
coverage:
	@ richgo test -v -coverprofile=/tmp/profile -covermode=atomic ./...
	@ go tool cover -html=/tmp/profile

.PHONY: validate-ci-config
validate-ci-config:
	@ circleci config validate -c .circleci/config.yml

.PHONY: local-ci
local-ci: validate-ci-config
	@ circleci local execute --job $(JOB)

