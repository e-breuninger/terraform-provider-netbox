TEST?=netbox/*.go
GOFMT_FILES?=$$(find . -name '*.go' |grep -v vendor)

NETBOX_VERSION?=v3.1.3 
NETBOX_SERVER_URL=http://localhost:8001
SUPERUSER_API_TOKEN=0123456789abcdef0123456789abcdef01234567

default: testacc

# Run acceptance tests
.PHONY: testacc
testacc:
	# Set both NETBOX_TOKEN options to avoid collisions with existing environment settings
	TF_ACC=1 NETBOX_SERVER_URL=$(NETBOX_SERVER_URL) NETBOX_TOKEN='' NETBOX_API_TOKEN=$(SUPERUSER_API_TOKEN) go test -v -cover $(TEST)

.PHONY: test
test: 
	go test $(TEST) $(TESTARGS) -timeout=120s -parallel=4 -cover

# Run dockerized Netbox for acceptance testing
.PHONY: docker-up
docker-up: 
	echo "Startup and wait for Netbox to become ready"
	SUPERUSER_API_TOKEN=$(SUPERUSER_API_TOKEN) NETBOX_VERSION=$(NETBOX_VERSION) docker-compose -f docker/docker-compose.yml up --build wait
	docker-compose -f docker/docker-compose.yml logs
	echo "🚀 Netbox is up and running!"

.PHONY: docker-logs
docker-logs: 
	docker-compose -f docker/docker-compose.yml logs
	
.PHONY: docker-down
docker-down: 
	docker-compose -f docker/docker-compose.yml down

#! Development 
# The following make goals are only for local usage 
.PHONY: fmt
fmt:
	go fmt
	go fmt netbox/*.go
