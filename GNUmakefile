TEST?=netbox/*.go
TEST_FUNC?=TestAccNetboxVirtualMachine*
GOFMT_FILES?=$$(find . -name '*.go' | grep -v vendor)
DOCKER_COMPOSE=docker compose

export NETBOX_VERSION=v4.1.11
export NETBOX_SERVER_URL=http://localhost:8001
export NETBOX_API_TOKEN=0123456789abcdef0123456789abcdef01234567
export NETBOX_TOKEN=$(NETBOX_API_TOKEN)

default: testacc

# Run acceptance tests
.PHONY: testacc
testacc: docker-up
	@echo "⌛ Startup acceptance tests on $(NETBOX_SERVER_URL) with version $(NETBOX_VERSION)"
	TF_ACC=1 go test -timeout 20m -v -cover $(TEST)

.PHONY: testacc-specific-test
testacc-specific-test: # docker-up
	@echo "⌛ Startup acceptance tests on $(NETBOX_SERVER_URL) with version $(NETBOX_VERSION)"
	@echo "⌛ Testing function $(TEST_FUNC)"
	TF_ACC=1 go test -timeout 20m -v -cover $(TEST) -run $(TEST_FUNC)

.PHONY: test
test:
	go test $(TEST) $(TESTARGS) -timeout=120s -parallel=4 -cover

# Run dockerized Netbox for acceptance testing
.PHONY: docker-up
docker-up:
	@echo "⌛ Startup Netbox $(NETBOX_VERSION) and wait for service to become ready"
	$(DOCKER_COMPOSE) -f docker/docker-compose.yml up --build wait
	$(DOCKER_COMPOSE) -f docker/docker-compose.yml logs
	@echo "🚀 Netbox is up and running!"

.PHONY: docker-logs
docker-logs:
	$(DOCKER_COMPOSE) -f docker/docker-compose.yml logs

.PHONY: docker-down
docker-down:
	$(DOCKER_COMPOSE) -f docker/docker-compose.yml down --volumes

.PHONY: docs
docs:
	NETBOX_API_TOKEN="" NETBOX_SERVER_URL="" go generate ./...

#! Development
# The following make goals are only for local usage
.PHONY: fmt
fmt:
	go fmt
	go fmt netbox/*.go
