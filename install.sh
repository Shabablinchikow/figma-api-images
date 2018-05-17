#!/bin/bash
apt install golang-go -y
ln -s /usr/local/go/bin/go /usr/bin/go
mkdir /srv/figma
export GOPATH=/srv/figma
go get github.com/hoisie/web
cd /srv/figma
git clone https://git.shbbl.ru/shabablinchikow/figma-api.git
cd figma-api
go build
cp figma.service /etc/systemd/system/figma.service
systemctl daemon-reload
systemctl enable figma.service
