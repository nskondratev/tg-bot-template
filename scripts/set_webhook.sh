#!/bin/bash

ENV_FILE=.env

# Check existence of .env file
if [ ! -f "$ENV_FILE" ]; then
  echo ".env doesn't exist. Rename example file and fill it: mv .env.example .env && nano .env"
  exit 1
fi

# Load variables from .env file
source $ENV_FILE

# Parse -c option
CLEAR=false
while getopts ":c" opt; do
  case ${opt} in
    c )
      CLEAR=true
      ;;
    \? ) echo "Usage: ./scripts/set_webhook.sh [-c], where -c - clear existing webhook"
      exit 2
      ;;
  esac
done

# Build request params
REQ_PARAMS="url=$WEBHOOK_URL"

if [ $CLEAR = true ]; then
  REQ_PARAMS='url='
fi

echo "Response from Telegram on /setWebhook request:"

curl -F $REQ_PARAMS "https://api.telegram.org/bot$TELEGRAM_API_TOKEN/setWebhook"

printf "\n"
