
alias ut := unit-test
[group("ci")]
unit-test path="./..." *args="":
    @go test {{path}} {{args}}

alias ut-cov := unit-test-coverage
[group("ci")]
unit-test-coverage path="./..." *args="":
    @go test -cover -coverprofile=coverage.out {{path}} {{args}}

[group("ci")]
lint:
    @golangci-lint run -c .golangci.yaml