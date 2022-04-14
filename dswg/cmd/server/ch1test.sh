#!/usr/bin/zsh
curl -X POST localhost:8080 -d '{"record":{"value":"1234"}}'
curl -X POST localhost:8080 -d '{"record":{"value":"5678"}}'
curl -X POST localhost:8080 -d '{"record":{"value":"90AB"}}'
curl -X GET localhost:8080 -d '{"offset":0}'
curl -X GET localhost:8080 -d '{"offset":1}'
curl -X GET localhost:8080 -d '{"offset":2}'
