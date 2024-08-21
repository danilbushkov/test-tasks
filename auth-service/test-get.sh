#!/usr/bin/env bash




curl -i -X POST -H "Content-Type: application/json" -d '{"uuid": "'"$1"'"}' http://localhost:3000/api/auth/get
