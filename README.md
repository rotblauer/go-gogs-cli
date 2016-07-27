# Gogs CLI

Accesses the [Gogs Client API](https://github.com/gogits/go-gogs-client), currently available on [Gogs' `develop` branch](https://github.com/gogits/gogs/tree/develop).

[Cobra](https://github.com/spf13/cobra) & [Viper](https://github.com/spf13/viper) go packages handle the hard work for the CLI interface.

Very much still a __work in progress__. That's in bold.

## Setup

Clone the repo and build it yourself, or `go get github.com/irstacks/gogs-cli`.

Make sure the `gogs` executable is available somewhere in your `$PATH`.

This repo's build is for darwin. You can build for linux with the nifty `env GOOS=linux go build -o gogs`.

## Usage

So far...

```bash
$ gogs repo # index all your repos
$ gogs repo -n my-new-repo
$ gogs repo -o anorganizationiown -n my-new-repo
```

## Help out

Pull request at will.
