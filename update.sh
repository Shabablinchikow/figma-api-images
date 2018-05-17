#!/bin/bash
export GOPATH=/srv/figma
cd /srv/figma/figma-api
git pull origin master 
go build
systemctl restart figma.service
