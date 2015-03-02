### GitFish

Simple HTTP listener for Github post commit hooks.

> Why 'GitFish'?
>
> What receives a hook?!

[![Build Status](https://travis-ci.org/jmervine/gitfish.svg?branch=master)](https://travis-ci.org/jmervine/gitfish)

#### Install

```
go get -u github.com/jmervine/gitfish
```

> Note: ensure `GOBIN` environment variable is set.
>
> Something like:
>
> `$ test "$GOBIN" || (mkdir ~/.gobin && export GOBIN=~/.gobin)`

#### Download

Currently I only have the `linux/x86_64` binaries ready.

```
$ curl -sS -O http://static.mervine.net/go/linux/x86_64/gitfish && chmod 755 gitfish
$ ./gitfish help
```

#### Usage

```
NAME:
   gitfish - http listener and handler for github post commit hooks

USAGE:
   gitfish [global options] command [command options] [arguments...]

VERSION:
   0.0.1

AUTHOR:
  Joshua Mervine - <joshua@mervine.net>

COMMANDS:
   help, h	Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --command, -c        command to execute *required* [$FISH_COMMAND]
   --port, -p "8888"    http listener port [$FISH_PORT]
   --verify             run command on startup to verify [$FISH_VERIFY]
   --secret, -s         require a secret to be passed by github [$FISH_SECRET]
   --branches, -b       filter on branch names, comma delim [$FISH_BRANCH]
   --owner              filer, require repo owner push [$FISH_OWNER]
   --admin              filer, require repo admin push [$FISH_ADMIN]
   --master             filer, require branch be assigned as master branch [$FISH_MASTER]
   --help, -h           show help
   --version, -v        print the version
```
