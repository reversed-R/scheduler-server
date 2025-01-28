#!/bin/sh

curl -X POST -H "Content-Type: application/json" \
  -d '{
    "name":"John",
    "comment":"I am John!"
  }' localhost:8080/api/v1/rooms/2/users
