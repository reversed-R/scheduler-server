#!/bin/sh

curl -X POST -H "Content-Type: application/json" \
  -d '{
    "name":"testroom",
    "description":"this is a test room",
    "beginTime":{
      "year": 2025,
      "month": 1,
      "day": 28,
      "hour": 20,
      "min": 10
    },
	  "dayLength": 5,
	  "dayPatternId": "AM_AND_PM",
	  "dayPatternLength": 2
  }' localhost:8080/api/v1/rooms
