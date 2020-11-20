#!/usr/bin/env bash
# bash

echo "go test: github.com/d3ta-go/ddd-mod-geolocation/modules/geolocation/infrastructure/service/rdbms... "
echo "-------------------------------------------------------------------------------"
echo ""

go test -timeout 120s  github.com/d3ta-go/ddd-mod-geolocation/modules/geolocation/infrastructure/service/rdbms -v -cover

echo ""
echo "-------------------------------------------------------------------------------"
echo "go test: DONE "
echo ""