#!/usr/bin/env bash

# Environment variables
echo "Load env variables..."

#--- Services ---
export HTTP_PORT=8000

#--- Database ---
export POSTGRES_DSN="host=localhost user=postgres password=postgres dbname=kore port=5432 sslmode=disable TimeZone=Asia/Jakarta"

#--- Auth ---
export TOKEN_SECRET=token-secret
export TOKEN_HEADER_KEY_ID=drPB_Wlr8gCYSaNp4GxJi6w61b8N1oosZQ8sxD9R1Is
