#!/usr/bin/env bash

set -x

openssl req -new -x509 -days 36500 -extensions v3_ca -keyout ca.key -out ca.crt
openssl req -new -x509 -days 36500 -extensions v3_ca -keyout ca.key -out ca.pem

openssl genrsa -des3 -out server.key 2048
openssl req -out server.csr -key server.key -new
openssl x509 -req -in server.csr -CA ca.crt -CAkey ca.key -CAcreateserial -out server.crt -days 36500

openssl genrsa -out client-key.pem 2048
openssl req -out client.csr -key client-key.pem -new
openssl x509 -req -in client.csr -CA ca.crt -CAkey ca.key -CAcreateserial -out client-crt.pem -days 36500