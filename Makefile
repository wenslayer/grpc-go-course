# ======================================================================
# Makefile for Go GRPC Cource
# ======================================================================

# ======================================================================
# Help
# ======================================================================

SHELL     := /bin/bash
GREEN     := \033[0;32m
COLOR_OFF := \033[0m
CHECK     := \xE2\x9C\x93
PREMSG    := \n$(GREEN)$(CHECK) # include trailing space
POSTMSG   := $(COLOR_OFF)\n\n

.PHONY: default help
default: help
help: ## Show help message
	@printf "$(PREMSG) usage: make [target]\n$(POSTMSG)"
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%s\033[0m#%s\n", $$1, $$2}' $(MAKEFILE_LIST) | column -t -s# | sort

# ======================================================================
# SSL/TLS Support
# ======================================================================

# Inspired from: https://github.com/grpc/grpc-java/tree/master/examples#generating-self-signed-certificates-for-use-with-grpc

# Output files
# ca.key: Certificate Authority private key file (this shouldn't be shared in real-life)
# ca.crt: Certificate Authority trust certificate (this should be shared with users in real-life)
# server.key: Server private key, password protected (this shouldn't be shared)
# server.csr: Server certificate signing request (this should be shared with the CA owner)
# server.crt: Server certificate signed by the CA (this would be sent back by the CA owner) - keep on server
# server.pem: Conversion of server.key into a format gRPC likes (this shouldn't be shared)

# Summary
# Private files: ca.key, server.key, server.pem, server.crt
# "Share" files: ca.crt (needed by the client), server.csr (needed by the CA)

SSL_DIR   ?= ssl

SSL_CA_CRT = $(SSL_DIR)/ca.crt
SSL_CA_KEY = $(SSL_DIR)/ca.key
SSL_SERVER_CRT = $(SSL_DIR)/server.crt
SSL_SERVER_CSR = $(SSL_DIR)/server.csr
SSL_SERVER_KEY = $(SSL_DIR)/server.key
SSL_SERVER_PEM = $(SSL_DIR)/server.pem

SSL_FILES = $(SSL_CA_CRT) $(SSL_CA_KEY) $(SSL_SERVER_CRT) $(SSL_SERVER_CSR) $(SSL_SERVER_KEY) $(SSL_SERVER_PEM)

# SERVER_CN ?= localhost
SERVER_CN ?= localhost.localdomain
# SERVER_CN ?= $(shell hostname)
SSL_SUBJ  := /C=US/ST=WA/O=DreamBox Learning/CN=$(SERVER_CN)
# SSL_SUBJ  := /C=US/ST=WA/O=DreamBox Learning
# SSL_SAN_CONFIG := subjectAltName=DNS:$(SERVER_CN)
		# -reqexts SAN \
		# -config <(cat /etc/ssl/openssl.cnf; printf "\n[SAN]\n$(SSL_SAN_CONFIG)") \
		#

# Set this so that we can use CN in the X509 certificate.
# Otherwise, you get connection error:
# 	transport: authentication handshake failed:x509: certificate relies on legacy Common Name field,
#	use SANs or temporarily enable Common Name matching with GODEBUG=x509ignoreCN=0"
export GODEBUG = x509ignoreCN=0

.PHONY: ssl-all
ssl-all: $(SSL_FILES) ## Generate all SSL files

$(SSL_DIR):
	mkdir -p "$@"

# Step 1
$(SSL_CA_KEY): $(SSL_DIR) $(MAKEFILE_LIST)
	@printf "$(PREMSG)Generate Certificate Authority ($@)$(POSTMSG)"
	openssl genrsa \
		-passout pass:1111 \
		-des3 \
		-out "$@" \
		4096
	chmod -v go-rwx "$@"
$(SSL_CA_CRT): $(SSL_CA_KEY) $(MAKEFILE_LIST)
	@printf "$(PREMSG)Generate Trust Certificate ($@)$(POSTMSG)"
	# openssl req -passin pass:1111 -new -x509 -days 3650 -key $< -out $@ -subj "/CN=$(SERVER_CN)"
	openssl req \
		-passin pass:1111 \
		-new \
		-x509 \
		-days 3650 \
		-key "$<" \
		-subj "$(SSL_SUBJ)" \
		-out "$@"

# Step 2
$(SSL_SERVER_KEY): $(SSL_DIR) $(MAKEFILE_LIST)
	@printf "$(PREMSG)Generate the Server Private Key ($@)$(POSTMSG)"
	openssl genrsa \
		-passout pass:1111 \
		-des3 \
		-out "$@" \
		4096
	chmod -v go-rwx "$@"

# Step 3
$(SSL_SERVER_CSR): $(SSL_SERVER_KEY) $(MAKEFILE_LIST)
	@printf "$(PREMSG)Get a certificate signing request from the CA ($@)$(POSTMSG)"
	openssl req \
		-passin pass:1111 \
		-new \
		-key "$<" \
		-subj "$(SSL_SUBJ)" \
		-out "$@"
	openssl req -text -noout -in "$@"

# Step 4
$(SSL_SERVER_CRT):$(SSL_SERVER_CSR) $(SSL_CA_CRT) $(SSL_CA_KEY) $(MAKEFILE_LIST)
	@printf "$(PREMSG)Self-sign the certificate with the CA we created ($@)$(POSTMSG)"
	openssl x509 \
		-req \
		-passin pass:1111 \
		-days 3650 \
		-in "$<" \
		-CA "$(SSL_CA_CRT)" \
		-CAkey "$(SSL_CA_KEY)" \
		-set_serial 01 \
		-out "$@"
	chmod -v go-rwx "$@"

# Step 5
$(SSL_SERVER_PEM):$(SSL_SERVER_KEY) $(MAKEFILE_LIST)
	@printf "$(PREMSG)Convert the server certificate to .pem format ($@)$(POSTMSG)"
	openssl pkcs8 \
		-topk8 \
		-nocrypt \
		-passin pass:1111 \
		-in "$<" \
		-out "$@"
	chmod -v go-rwx "$@"

.PHONY: clean-ssl
clean-ssl:
	rm -frv -- "./$(SSL_DIR)"

# ======================================================================
# This project
# ======================================================================

greet_proto  := greet/greetpb/greet.proto
greet_server := greet/greet_server/server.go
greet_client := greet/greet_client/client.go

calc_proto   := calculator/calculatorpb/calculator.proto
calc_server  := calculator/calculator_server/server.go
calc_client  := calculator/calculator_client/client.go

all_proto    := $(greet_proto) $(calc_proto)
all_proto_go := $(all_proto:%.proto=%.pb.go)

GO_RUN_FLAGS := -ldflags "-X main.HostAndPort=$(SERVER_CN):50051 -X main.CACertFile=$(SSL_CA_CRT) -X main.ServerCertFile=$(SSL_SERVER_CRT) -X main.ServerKeyFile=$(SSL_SERVER_PEM)"

.PHONY: gen-proto
gen-proto: $(all_proto_go) ## Generate code from all proto files
%.pb.go: %.proto $(MAKEFILE_LIST)
	@printf "$(PREMSG)Generate code from all proto files$(POSTMSG)"
	protoc "$<" --go_out=plugins=grpc:.

.PHONY: clean-proto
clean-proto: ## Clean up all generated code from proto files
	@printf "$(PREMSG)Clean up all generated code from proto files$(POSTMSG)"
	rm -f -- $(all_proto_go)

.PHONY: clean-all
clean-all: clean-proto clean-ssl ## Run all 'clean-*' targets

.PHONY: run-greet-server run-greet-client run-calc-server run-calc-client
run-greet-server: $(greet_server) $(greet_proto) ssl-all ## Start the greet server
	@printf "$(PREMSG)Start the greet server$(POSTMSG)"
	go run $(GO_RUN_FLAGS) $<

run-greet-client: $(greet_client) $(greet_proto) ## Run the greet client
	@printf "$(PREMSG)Run the greet client$(POSTMSG)"
	go run $(GO_RUN_FLAGS) $<

run-calc-server: $(calc_server) $(calc_proto) ssl-all ## Start the calc server
	@printf "$(PREMSG)Start the calc server$(POSTMSG)"
	go run $(GO_RUN_FLAGS) $<

run-calc-client: $(calc_client) $(calc_proto) ## Run the calc client
	@printf "$(PREMSG)Run the calc client$(POSTMSG)"
	go run $(GO_RUN_FLAGS) $<
