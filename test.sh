#!/bin/bash

contenttype="Content-Type: application/json"
rooturl="localhost:8080"

route="search"
query=$1

curl -H $contenttype -d "{\"query\":\"$query\"}" "$rooturl/$route"
