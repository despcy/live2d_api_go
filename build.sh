#!/bin/bash
env GOOS=darwin GOARCH=386 go build -o server_mac_386

env GOOS=linux GOARCH=386 go build -o server_linux_386

env GOOS=windows GOARCH=386 go build -o server_win_386