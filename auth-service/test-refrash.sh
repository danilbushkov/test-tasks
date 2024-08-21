#!/usr/bin/env bash




curl -i -X POST -H "Content-Type: application/json" -d '{"refreshToken": "'"$1"'"}' http://localhost:3000/api/auth/refresh
