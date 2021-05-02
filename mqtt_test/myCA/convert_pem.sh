#!/usr/bin/env bash

set -x

openssl rsa -in client.key -out client-key.pem

openssl x509 -in client.crt -out client-crt.pem