# rutracker-proxy
[![Build Status](https://travis-ci.org/zhulik/rutracker-proxy.svg?branch=master)](https://travis-ci.org/zhulik/rutracker-proxy)

GUI-less tool for proxying torrent client announces to blocked in Russia rutracker announcers.
Inspired by original [tool](https://github.com/RutrackerOrg/rutracker-proxy). Automatic proxy
rotation every 5 minutes(configurable).

Fedora package lives [here](http://koji.russianfedora.pro/koji/packageinfo?packageID=140), and it's specs
[here](https://github.com/RussianFedora/rutracker-proxy)(thanks to [@Vascom](https://github.com/vascom))

## Dependencies
* Golang compiler

For ArchLinux:

`pacman -S go`

## Installing
`go install github.com/zhulik/rutracker-proxy`

## Building from sources
Proxy is written in Go, so build steps are ordinary for Go software. Clone the repository and
then in it's root directory run

`go get && go build`

## Running

`rutracker-proxy`

## Options

`rutracker-proxy --help`

## TODO
* Unit tests
* Testing in production
* Automatic binary builds on releases for all platforms on Travis
* PKGBUILD
* Transport wrapper for handling proxy errors
* SSL support

## Contribution
You know;-)