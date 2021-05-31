TEST?=netbox/*.go
GOFMT_FILES?=$$(find . -name '*.go' |grep -v vendor)

NETBOX_SERVER_URL?=http://localhost:8001

default: testacc

# Run acceptance tests
.PHONY: testacc
testacc: 
	@NETBOX_VERSION=v2.10.10 sh -c "'$(CURDIR)/docker/start.sh' $(NETBOX_SERVER_URL)"
	TF_ACC=1 NETBOX_SERVER_URL=$(NETBOX_SERVER_URL) NETBOX_API_TOKEN=0123456789abcdef0123456789abcdef01234567 go test -v -cover $(TEST)

.PHONY: test
test: 
	go test $(TEST) $(TESTARGS) -timeout=120s -parallel=4 -cover

#! Development 
# The following make goals are only for local usage 

.PHONY: fmt
fmt:
	go fmt
	go fmt netbox/*.go
