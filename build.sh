#!/bin/bash

# build web UI

cd /root/golang/gopath/src/StreamingWebsite/web
go install
cp /root/golang/gopath/bin/web /root/golang/goapth/bin/video_server_web_ui/web
cp -R /root/golang/gopath/src/StreamingWebsite/templates /root/golang/gopath/bin/video_server_web_ui/

