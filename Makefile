# 单元测试
.PHONY: ut
ut:
	@go test -race ./... -count=1

.PHONY: setup
setup:
	@sh ./.script/setup.sh

.PHONY: lint
lint:
	golangci-lint run -c ./.script/.golangci.yml

.PHONY: fmt
fmt:
	@sh ./.script/fmt.sh

.PHONY: tidy
tidy:
	@go mod tidy -v

.PHONY: check
check:
	@$(MAKE) --no-print-directory fmt
	@$(MAKE) --no-print-directory tidy

# e2e 测试
.PHONY: e2e
e2e:
	sh ./.script/integrate_test.sh

.PHONY: e2e_up
e2e_up:
	docker compose -p dbproxy -f .script/integration_test_compose.yml up -d

.PHONY: e2e_down
e2e_down:
	docker compose -p dbproxy -f .script/integration_test_compose.yml down -v