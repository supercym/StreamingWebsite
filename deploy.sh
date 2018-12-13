#!/bin/bash

kill -9 $(pgrep webserver)
cd /root/StreamingWebsite/
git pull https://github.com/supercym/StreamingWebsite.git
cd webserver/
chmod a+x webserver
./webserver &
