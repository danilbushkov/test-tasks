#!/usr/bin/env bash




curl -i -X POST -H "Content-Type: application/json" -d '{"uuid": "435dd1a3-6862-4f9a-929b-b7e130462b42"}' http://localhost:3000/api/auth/get
