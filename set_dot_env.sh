#!/bin/sh

# usage: ./set_dot_env "proxy_url"

# PROXY=$1
touch .env2

# echo "PROXY=$1" >> .env2

echo "UID=$(id -u $USER)" >> .env2
echo "GID=$(id -g $USER)" >> .env2
echo "UNAME=$USER" >> .env2
# echo "WANDB_API_KEY=" >> .env
