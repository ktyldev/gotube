#!/bin/bash

contenttype="Content-Type: application/json"
rooturl="localhost:8080"

downloadurl=$1

curl -H $contenttype -d "{\"url\":\"$downloadurl\"}" "$rooturl/queue/add"
