#!/bin/sh
# author - S@M

## Create the ipfs container
cd docker-ipfs && docker build . -t ipfs-local

## Start the ipfs container
docker-compose up -d ipfs

