#!/bin/bash

# Key considerations for algorithm "RSA" ≥ 2048-bit
#openssl genrsa -out server.key 2048

# Key considerations for algorithm "ECDSA" ≥ secp384r1
# List ECDSA the supported curves (openssl ecparam -list_curves)
#openssl ecparam -genkey -name secp384r1 -out server.key

#openssl req -new -x509 -sha256 -key server.key -out server.crt -days 3650

env GOOS=darwin GOARCH=386 go build -o server_mac_386

env GOOS=linux GOARCH=386 go build -o server_linux_386

env GOOS=windows GOARCH=386 go build -o server_win_386