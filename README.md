## Go(lang) Git Fish!

Simple HTTP listener for Github post commit hooks.

#### Usage

```
NAME:
   go-git-fish - http listener and handler for github post commit hooks

USAGE:
   go-git-fish [global options] command [command options] [arguments...]

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
   --token, -t          require a token to be passed via query string [$FISH_TOKEN]
   --branches, -b       filter on branch names, comma delim [$FISH_BRANCH]
   --owner              filer, require repo owner push [$FISH_OWNER]
   --admin              filer, require repo admin push [$FISH_ADMIN]
   --master             filer, require branch be assigned as master branch [$FISH_MASTER]
   --help, -h           show help
   --version, -v        print the version
```
