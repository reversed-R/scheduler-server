#!/bin/sh

curl -X POST -H "Content-Type: application/json" \
  -d '{
    "name":"John",
    "comment":"I am John!",
    "availabilities": [
      "OK", "OK", "BAD", "NOT_BAD", "BAD", "OK", "NOT_BAD", "BAD", "NOT_BAD", "OK"
    ]
  }' localhost:8080/api/v1/rooms/1/users
  
curl -X POST -H "Content-Type: application/json" \
  -d '{
    "name":"Smith",
    "comment":"I am Smith!",
    "availabilities": [
      "NOT_BAD", "NOT_BAD", "BAD", "NOT_BAD", "BAD", "OK", "OK", "OK", "OK", "BAD" 
    ]
  }' localhost:8080/api/v1/rooms/1/users
