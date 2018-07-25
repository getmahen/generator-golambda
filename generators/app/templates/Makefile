NO_COLOR=\033[0m
OK_COLOR=\033[0;32m
TS_COLOR=\033[1;30m
BUILD_DATE_VER=""
TIMESTAMP := $(shell date -u +%Y%m%d%H%M%S)

.PHONY: test
test:
	@echo "$(OK_COLOR)==> Testing$(NO_COLOR)"
	go test -cover -race ./...

.PHONY: clean
clean:
	@echo "$(TS_COLOR)$(shell date "+%Y/%m/%d %H:%M:%S")$(NO_COLOR)$(OK_COLOR) ==> Cleaning$(NO_COLOR)"
	@go clean
	@rm -f <%= lambdaName %> <%= lambdaName %>.zip

.PHONY: build
build: clean
	@echo "$(TS_COLOR)$(shell date "+%Y/%m/%d %H:%M:%S")$(NO_COLOR)$(OK_COLOR)==> Building$(NO_COLOR)"
	GOOS=linux go build --ldflags "-X bitbucket.org/credomobile/<%= lambdaName %>/handler.version=`git rev-parse HEAD`" -o <%= lambdaName %>
	zip <%= lambdaName %>.zip <%= lambdaName %> vault-cas.crt

.PHONY: upload
upload:
	@echo "$(TS_COLOR)$(shell date "+%Y/%m/%d %H:%M:%S")$(NO_COLOR)$(OK_COLOR)==> Deploying Zip to s3$(NO_COLOR)"
	aws s3 cp <%= lambdaName %>.zip s3://credo-dev-lambdas/<%= lambdaName %>.zip --metadata GitHash=`git rev-parse HEAD`