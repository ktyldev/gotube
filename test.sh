#!/bin/bash

contenttype="Content-Type: application/json"
rooturl="localhost:8080"

curl -H $contenttype -d '{"url":"memes"}' "$rooturl/queue/add"
