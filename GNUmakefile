TEST?=netbox/*.go
GOFMT_FILES?=$$(find . -name '*.go' |grep -v vendor)

# Export variables to downstream processes
.EXPORT_ALL_VARIABLES:

# These variables can be changed by the user
NETBOX_HOST?=localhost
NETBOX_PORT?=8001
NETBOX_VERSION?=v2.11.12

# Set static in order to avoid collisions with other netbox intergrations like netbox
NETBOX_API_TOKEN=0123456789abcdef0123456789abcdef01234567
NETBOX_SERVER_URL=http://${NETBOX_HOST}:$(NETBOX_PORT)

default: testacc

# Run acceptance tests
.PHONY: testacc
testacc: netbox.start
	TF_ACC=1 gotestsum --format testname --hide-summary skipped -- -v -cover $(TEST)

# Run unit tests
.PHONY: test
test: 
	gotestsum --format testname --hide-summary skipped -- $(TEST) $(TESTARGS) -timeout=120s -parallel=4 -cover

.PHONY: download
download:
	@echo Download go.mod dependencies
	@go mod download

.PHONY: install-tools
install-tools: download
	@echo Installing tools from tools.go
	@cat tools.go | grep _ | awk -F'"' '{print $$2}' | xargs -tI % go install %


#! Development 
# The following make goals are only for local usage 

.PHONY: netbox.start 
netbox.start: 
	sh docker/start_netbox.sh

.PHONY: netbox.stop 
netbox.stop: 
	docker-compose -f  docker/docker-compose.yml down

.PHONY: fmt
fmt:
	go fmt
	go fmt netbox/*.go