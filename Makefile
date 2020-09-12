greet_proto  := greet/greetpb/greet.proto
greet_server := greet/greet_server/server.go
greet_client := greet/greet_client/client.go

calc_proto   := calculator/calculatorpb/calculator.proto
calc_server  := calculator/calculator_server/server.go
calc_client  := calculator/calculator_client/client.go

all_proto    := $(greet_proto) $(calc_proto)
all_proto_go := $(all_proto:%.proto=%.pb.go)

SERVER_CN ?= localhost
SSL_DIR   ?= ssl

.PHONY: gen-proto
gen-proto: $(all_proto_go)
%.pb.go: %.proto
	protoc $< --go_out=plugins=grpc:.

.PHONY: clean-proto
clean-proto:
	rm -f -- $(all_proto_go)

.PHONY: clean-all
clean-all: clean-proto

.PHONY: run-greet-server run-greet-client run-calc-server run-calc-client
run-greet-server: $(greet_server) $(greet_proto)
	go run $<
run-greet-client: $(greet_client) $(greet_proto)
	go run $<
run-calc-server: $(calc_server) $(calc_proto)
	go run $<
run-calc-client: $(calc_client) $(calc_proto)
	go run $<

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

.PHONY: ssl-all
ssl-all: $(SSL_DIR)/server.crt $(SSL_DIR)/server.pem

$(SSL_DIR):
	mkdir -p $@

# Step 1: Generate Certificate Authority + Trust Certificate (ca.crt)
$(SSL_DIR)/ca.key: $(SSL_DIR)
	openssl genrsa -passout pass:1111 -des3 -out $@ 4096
	chmod go-rwx $@
$(SSL_DIR)/ca.crt: $(SSL_DIR)/ca.key
	openssl req -passin pass:1111 -new -x509 -days 3650 -key $< -out $@ -subj "/CN=$(SERVER_CN)"

# Step 2: Generate the Server Private Key (server.key)
$(SSL_DIR)/server.key: $(SSL_DIR)
	openssl genrsa -passout pass:1111 -des3 -out $@ 4096
	chmod go-rwx $@

# Step 3: Get a certificate signing request from the CA (server.csr)
$(SSL_DIR)/server.csr: $(SSL_DIR)/server.key
	openssl req -passin pass:1111 -new -key $< -out $@ -subj "/CN=${SERVER_CN}"

# Step 4: Sign the certificate with the CA we created (it's called self signing) - server.crt
$(SSL_DIR)/server.crt: $(SSL_DIR)/server.csr $(SSL_DIR)/ca.crt $(SSL_DIR)/ca.key
	openssl x509 -req -passin pass:1111 -days 3650 -in $< -CA $(SSL_DIR)/ca.crt -CAkey $(SSL_DIR)/ca.key -set_serial 01 -out $@
	chmod go-rwx $@

# Step 5: Convert the server certificate to .pem format (server.pem) - usable by gRPC
$(SSL_DIR)/server.pem: $(SSL_DIR)/server.key
	openssl pkcs8 -topk8 -nocrypt -passin pass:1111 -in $< -out $@
	chmod go-rwx $@

clean-ssl:
	rm -fr -- ./$(SSL_DIR)
