TEST?=./internal/provider/
TEST_FUNC?=TestAccNetboxMACAddr*
DOCKER_COMPOSE=docker compose

export NETBOX_VERSION=v4.6.10
export NETBOX_SERVER_URL=http://localhost:8001
# The admin token docker/docker-compose.yml configures, joined as
# nbt_<SUPERUSER_API_KEY>.<SUPERUSER_API_TOKEN>.
export NETBOX_API_TOKEN := nbt_mrcLJrER9ves.Iw2JZtpMdYjBuIVIZLItd4lB2RYEiv6xhY23W0tQ
# The parallel suite against one container overwhelms the provider's 10s default.
export NETBOX_REQUEST_TIMEOUT := 60
# Refresh the state after every apply: the related-object counts a write response carries go stale
# when the same apply creates or deletes a related object afterwards, and the harness compares
# that state on import.
export TF_ACC_REFRESH_AFTER_APPLY := 1

default: testacc

# Run acceptance tests
.PHONY: testacc
testacc: docker-up
	@echo "⌛ Startup acceptance tests on $(NETBOX_SERVER_URL) with version $(NETBOX_VERSION)"
	TF_ACC=1 go test -tags acctest -timeout 30m -v -cover $(TEST)

.PHONY: testacc-specific-test
testacc-specific-test: docker-up
	@echo "⌛ Startup acceptance tests on $(NETBOX_SERVER_URL) with version $(NETBOX_VERSION)"
	@echo "⌛ Testing function $(TEST_FUNC)"
	TF_ACC=1 go test -tags acctest -timeout 30m -v -cover $(TEST) -run $(TEST_FUNC)

.PHONY: test
test:
	go test $(TEST) $(TESTARGS) -timeout=120s -parallel=4 -cover

# Run dockerized Netbox for acceptance testing
.PHONY: docker-up
docker-up:
	@echo "⌛ Startup Netbox $(NETBOX_VERSION) and wait for it to become healthy"
	$(DOCKER_COMPOSE) -f docker/docker-compose.yml up --build --detach --wait --wait-timeout 600 netbox
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

# Validates every docs example alone against the provider built from this checkout and checks that
# every resource and data source has one; see scripts/validate_examples.sh.
.PHONY: validate-examples
validate-examples:
	bash scripts/validate_examples.sh

#! Development
# The following make goals are only for local usage
.PHONY: fmt
fmt:
	go fmt ./...
