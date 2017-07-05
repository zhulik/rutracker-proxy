# rutracker-proxy
[![Build Status](https://travis-ci.org/zhulik/rutracker-proxy.svg?branch=master)](https://travis-ci.org/zhulik/rutracker-proxy)

GUI-less tool for proxying torrent client announces to blocked in Russia rutracker announcers.
Inspired by original [tool](https://github.com/RutrackerOrg/rutracker-proxy). Automatic proxy
rotation every 5 minutes(configurable).

Fedora package lives [here](http://koji.russianfedora.pro/koji/packageinfo?packageID=140), and it's specs
[here](https://github.com/RussianFedora/rutracker-proxy)(thanks to [@Vascom](https://github.com/vascom))

## Notes
* Transmission doesn't support proxies since 2.12 release. Deluge and QBittorrent does.
* For ArchLinux users: PKGBUILD can be found in archlinux directory. Config and service file are also there, they are the same as for Fedora(thanks to [@Vascom](https://github.com/vascom)). Package will be published on AUR after testing on my machine.

## Dependencies
* Golang compiler >= 1.6

For ArchLinux:

`pacman -S go`

## Installing

### Go way
`go install github.com/zhulik/rutracker-proxy`

### Archlinux way
`y -S rutracker-proxy`

### Fedora way
0. Install free repo if it is not installed yet

`# sudo dnf install --nogpgcheck http://mirror.yandex.ru/fedora/russianfedora/russianfedora/free/fedora/russianfedora-free-release-stable.noarch.rpm`

1. Install rutracker-proxy

`# sudo dnf install rutracker-proxy`

Alternatively you can download and install package manually from http://koji.russianfedora.pro/koji/packageinfo?packageID=140

2. Run rutracker-proxy by hand or via systemd unit with setting parameters in config file /etc/rutracker-proxy/rutracker-proxy.conf.

### Docker way

`docker pull zhulik/rutracker-proxy`

`docker run -d --rm zhulik/rutracker-proxy`

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
* Automatic binary builds on releases for all platforms on Travis
* Transport wrapper for handling proxy errors
* SSL support

## Contribution
You know;-)
