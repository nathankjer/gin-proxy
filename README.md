# Gin Proxy

A reverse proxy written in golang that uses the gin gonic web framework.

## Setup

1. sudo -u postgres psql -c 'create database requests;'
2. export PROXY_HOST=https://example.com
3. go run .