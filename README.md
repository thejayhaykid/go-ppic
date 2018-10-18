# <img src="./assets/go-ppic.png" width="32" height="32" valign="middle" title="go-ppic example"> <img src="./assets/hello-world.png" width="32" height="32" valign="middle" title="go-ppic hello-world example"> <img src="./assets/jackwilsdon.png" width="32" height="32" valign="middle" title="go-ppic jackwilsdon example"> go-ppic <a href="https://travis-ci.com/jackwilsdon/go-ppic" title="Build status"><img src="https://img.shields.io/travis/com/jackwilsdon/go-ppic.svg" valign="middle" title="Build status"></a> <a href="https://goreportcard.com/report/github.com/jackwilsdon/go-ppic" title="Go Report Card"><img src="https://goreportcard.com/badge/github.com/jackwilsdon/go-ppic" valign="middle" title="Go Report Card status"></a> <a href="https://godoc.org/github.com/jackwilsdon/go-ppic" title="GoDoc reference"><img src="https://godoc.org/github.com/jackwilsdon/go-ppic?status.svg" valign="middle" title="GoDoc reference"></a>

Profile picture generation service written in Go. A demo can be found at [ppic.now.sh](https://ppic.now.sh/hello).

`go-ppic` provides two commands; [`ppic`](#ppic) and [`ppicd`](#ppicd).

## ppic

`ppic` is used to generate profile pictures on the command line, without having to run a web server. `ppic` outputs the generated image to stdout.

### Installation

```Shell
go get -u github.com/jackwilsdon/go-ppic/cmd/ppic
```

### Usage

```Text
usage: ppic username [size] > image.png
```

> `size` defaults to 512 if not provided

### Examples

```Shell
ppic jackwilsdon 1024 > profile.png
```

## ppicd

`ppicd` is a web server providing image generation.

### Installation

```Shell
go get -u github.com/jackwilsdon/go-ppic/cmd/ppicd
```

### Usage

```Shell
  -d    enable pprof debug routes
  -h string
        host to run the server on
  -p uint
        port to run the server on (default 3000)
  -z    enable gzip compression
```

Once the server is running, you can retrieve an image by visiting `http://localhost:3000/image.png` (provided you haven't changed the port). Other supported file extensions are `.gif` and `.jpeg`, defaulting to `.png` if there is no file extension in the URL.

You can change the size of images which are generated by adding `?size=1024` to the end of the URL.

## Limitations

Here is a list of known limitations (subject to change);

* Image size must be a multiple of 8

