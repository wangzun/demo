#!/bin/bash
gomobile build -target=ios github.com/wangzun/demo/client
ios-deploy -r -b client.app
