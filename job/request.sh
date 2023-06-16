#!/bin/sh
#
curl -v -X POST http://localhost:8080/register -H 'content-type: application/json' \
    -d '{ "name": "teste", "email": "mail@mail.com" }' \
