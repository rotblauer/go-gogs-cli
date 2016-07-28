# Gogs CLI
Accesses the [Gogs Client API](https://github.com/gogits/go-gogs-client), which is currently available exclusively on [Gogs' `develop` branch](https://github.com/gogits/gogs/tree/develop).

> [Cobra](https://github.com/spf13/cobra) & [Viper](https://github.com/spf13/viper) go packages handle the hard work for the CLI interface.

Very much still a __work in progress__. That's in bold. 

## Setup.
Clone the repo and build it yourself, or `go get github.com/irstacks/go-gogs-cli`. Just make sure the `gogs` executable winds up somewhere in your `$PATH`.

Oh, _but how do you build it?_, you ask? _What path?_ you ask? 

Here's what's up. From the beginning. 

With the marvelous `go get`...
```bash
$ go get github.com/irstacks/go-gogs-cli
$ cd $GOPATH/src/irstacks/go-gogs-cli
```

With the almost-as-marvelous `git clone`...
```
$ cd where/i/like/to/put/funky/things
$ git clone https://github.com/irstacks/go-gogs-cli.git
```

I say almost-as-marvelous because if you use `git clone` you may run into issues about your $GOPATH. It happens. Since I haven't figured out how to consistently `go get` all the dependencies I need for a given Go project from esoteric location outside my $GOPATH (in part because I have so many esoteric locations and don't want to haggle with forever adjusting/amending my $GOPATH extension), I usually just wind up running `go run main.go` or `go build -o gogs` and `go get`ting the dependencies it complains about one-by-one. I know, it sucks. But that's just how I roll sometimes. 

```bash
$ go get ./... # Make sure we've got all the dependencies we need available.
$ go build -o gogs # Here we're using -o to tell go where to build the build it makes, in this case a file in the same directory we're with the name 'gogs' (because thats a lot shorter than go-gogs-cli). Note that whatever you name this file it what it will be accessible for you as on the CLI. So if you name it 'goo' (awesome.), then your commands will be all like: $ goo repo new ... 

# Now here you've got some options (probably) about in which of your $PATH's paths you want to stick it. I like to keep my custom thingeys out of the dredges, so I stick mine in $HOME/bin/
# ... If you follow in my footsteps, make sure somewhere in your bash/fish/zsh shell you've added $HOME/bin (NO SECOND SLASH, you slashing fiend you) to your $PATH, with something like `export $PATH="$PATH:$HOME/bin".
$ cp gogs ~/bin/
```


This repo's build is for darwin (Mac). You can build for linux with the nifty `env GOOS=linux go build -o gogs`. You can probably build for Window too but I don't trouble myself with such things.

## Usage.
So far, you can do things like...
```bash
$ gogs repo create my-new-repo --desc 'awesome stuff' --private --org GophersGophering # optional flag [-n|--name] if you want to be very particular
$ gogs repo new my-new-repo # new is an alias for create
$ gogs repo list # get all yo repos
$ gogs repo destroy irstacks my-exterminable-repo
$ gogs repo destroy irstacks/my-other-exterminable-repo
```

## Help out.
Please do!
