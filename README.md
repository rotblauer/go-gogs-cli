# Gogs CLI
Accesses the [Gogs Client API](https://github.com/gogits/go-gogs-client), which is currently available exclusively on [Gogs' `develop` branch](https://github.com/gogits/gogs/tree/develop).

> [Cobra](https://github.com/spf13/cobra) & [Viper](https://github.com/spf13/viper) go packages handle the hard work for the CLI interface.


## Setup.
Clone the repo and build it yourself, or `go get github.com/irstacks/go-gogs-cli`. Just make sure the `gogs` executable winds up somewhere in your `$PATH`.

Oh, _but how do you build it?_, you ask? _What path?_ you ask? 

Ok, fine. Here's what's up. From the beginning. 

With the marvelous `go get`...
```bash
$ go get github.com/irstacks/go-gogs-cli
$ cd $GOPATH/src/irstacks/go-gogs-cli
$ go get ./...
```

With the almost-as-marvelous `git clone`...
```bash
$ cd where/i/like/to/put/funky/things
$ git clone https://github.com/irstacks/go-gogs-cli.git
# Pedantry explained immediately below...
$ go get stuff
$ go get morestuff
```
Now, I say almost-as-marvelous because if you use `git clone` you may run into issues about your $GOPATH. It happens. Since I haven't figured out how to consistently `go get` all the dependencies I need for a given Go project from esoteric locations outside my $GOPATH (in part because I have so many esoteric locations and don't want to haggle with forever adjusting/amending my $GOPATH extensioners), I usually just wind up running `go run main.go` or `go build -o gogs` and `go get`ting the dependencies it complains about one-by-one. I know, it sucks. But that's just how I roll sometimes. 
<br>
<br>
Finally, we can build the sucker.
```bash
$ go build -o gogs
```
Here we're using -o to tell go where to build the build it makes, in this case a file in the same directory we're with the name 'gogs' (because thats a lot shorter than go-gogs-cli). Note that whatever you name this file it what it will be accessible for you as on the CLI. So if you name it 'goo' (awesome.), then your commands will be all like: `$ goo repo new ...` 

Now here you've got some options (probably) about in which of your $PATH's paths you want to stick it. I like to keep my custom thingeys out of the dredges, so I stick mine in `$HOME/bin/`
<br>
... If you follow in my footsteps, make sure somewhere in your bash/fish/zsh shell you've added $HOME/bin (NO SECOND SLASH, you slashing fiend you) to your $PATH, with something like `export $PATH="$PATH:$HOME/bin"`.
```bash
$ cp gogs ~/bin/
```
<br>

#### This repo's build is for darwin (Mac). 
You can build for linux with the nifty `env GOOS=linux go build -o gogs`. You can probably build for Windows too but I don't trouble myself with such things.

## Usage.
So far, you can do things. What's that? You can _do things_? Yep! Do things!
```bash
# Create basic.
$ gogs repo create wheres-waldo # Create a repo owned by you.
$ gogs repo [create|new|c|n] wheres-waldo # All the same thing (aliases).
# Create fancy.
$ gogs repo create wheres-waldo --desc 'awesome stuff' --private --org GophersGophering # Descriptively, privately, for an org you own.
# Create fancy awesome.
# --> Create a gogs remote and add it to your current working directory, 
# initializing git if necessary.
$ gogs repo create where-waldo -r origin
# [Aliases] for create. 
$ gogs repo [create|new|c|n]

# List basic.
$ gogs repo list # Get all yo repos.
$ gogs repo list -l 2 # Get only 2 of your repos. 

# Search basic.
$ gogs repo search waldo
# Search fancy.
$ gogs repo search waldo --limit 1 --user thatguy # yep, flags are still optional (must specify user if searching for private repos)
# [Aliases] for search.
$ gogs repo [search|find|s|f]

# Destroy basic. (Watch out! Destroy won't ask twice. )
$ gogs repo destroy irstacks my-exterminable-repo
$ gogs repo destroy irstacks/my-other-exterminable-repo
# [Aliases] for destroy.
$ gogs repo [destroy|delete|d|rid]
```
<br>
<br>
__Oh, you're hot shit and use n > 1 Gogs?__ _Sweet_.
<br> 
You can override your api and token by flagging a config file with the `--config` flag (like such)
```bash
$ gogs --config="$HOME/sneaky/place/.go-gogs-cli.yaml" repo create newjunk
```
<br>
or, override your api url and token individually on the fly with flags `--token` and `--url` for any command, like so:
```bash
$ gogs --url=http://some.other.company --token=qo23ransdlfknaw3oijr2323rasldf repo search waldo
```

## Config.
There's a file called `.go-gogs-cli.yaml` which handles configuring your __Gogs url__ and __token__, like such
```yaml
token: 0e6709o05da4753dddf5f592374fdc263f02n801
api_url: http://my.goggers.me
```
Fill that in for your own self.
<br>
<br>
You may have noticed that we're pretty heavy on the `gogs repo` and pretty light on the `gogs somethingelse` side of things. The [Gogs Client API](https://github.com/gogits/go-gogs-client) makes a bunch of endpoints and methods accessible for Users, Organizations, Issues, Admins, and so forth (although it's still very much a work in progress). Myself, I mostly just want to be able to create, search, and destroy like a fiend. If you would :heart: something and are unable to help yourself, let me know by opening an issue. 

## Help out.
:clap: [chanting] Do it! Do it! Do it!
