TEST?=netbox/*.go
TEST_FUNC?=TestAccNetboxMACAddr*
GOFMT_FILES?=$$(find . -name '*.go' | grep -v vendor)
DOCKER_COMPOSE=docker compose

export NETBOX_VERSION=v4.6.5
export NETBOX_SERVER_URL=http://localhost:8001
SUPERUSER_API_TOKEN := Iw2JZtpMdYjBuIVIZLItd4lB2RYEiv6xhY23W0tQ
SUPERUSER_API_KEY   := mrcLJrER9ves

# v4.5 switched to v2 tokens
# min(version, v4.5.0) == v4.5.0 means version >= v4.5.0.
IS_V2 := $(shell [ "$$(printf '%s\nv4.5.0\n' "$(NETBOX_VERSION)" | sort -V | head -1)" = "v4.5.0" ] && echo yes)

# v2 keys are randomly generated at container startup, so read the real key
# back from the DB. tail -n1 drops the manage.py banner; 2>/dev/null drops warnings.
GET_KEY = $(DOCKER_COMPOSE) -f docker/docker-compose.yml exec -T netbox \
  /opt/netbox/netbox/manage.py shell -c \
  "from users.models import Token; print(Token.objects.get(user__username='admin').key)" \
  2>/dev/null | tail -n1

# Lazy (=): the v2 key only exists after docker-up, so this runs per-recipe, not at parse time.
ifeq ($(IS_V2),yes)
NETBOX_API_TOKEN = nbt_$(shell $(GET_KEY)).$(SUPERUSER_API_TOKEN)
else
NETBOX_API_TOKEN = $(SUPERUSER_API_TOKEN)
endif

export NETBOX_API_TOKEN
export NETBOX_API_KEY   := $(SUPERUSER_API_KEY)
export SUPERUSER_API_TOKEN
export SUPERUSER_API_KEY
export API_TOKEN_PEPPER_1 := C9Bp3tmBgchWcWc2OkoFuaV9aoKZfhJgZUBG7g0PFxYzf1AkLd

default: testacc

# Run acceptance tests
.PHONY: testacc
testacc: docker-up
	@echo "⌛ Startup acceptance tests on $(NETBOX_SERVER_URL) with version $(NETBOX_VERSION)"
	TF_ACC=1 go test -timeout 20m -v -cover $(TEST)

.PHONY: testacc-specific-test
testacc-specific-test: docker-up
	@echo "⌛ Startup acceptance tests on $(NETBOX_SERVER_URL) with version $(NETBOX_VERSION)"
	@echo "⌛ Testing function $(TEST_FUNC)"
	TF_ACC=1 go test -timeout 20m -v -cover $(TEST) -run $(TEST_FUNC)

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

#! Development
# The following make goals are only for local usage
.PHONY: fmt
fmt:
	go fmt
	go fmt netbox/*.go
