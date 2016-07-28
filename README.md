# Gogs CLI
Accesses the [Gogs Client API](https://github.com/gogits/go-gogs-client), which is currently available exclusively on [Gogs' `develop` branch](https://github.com/gogits/gogs/tree/develop).

> [Cobra](https://github.com/spf13/cobra) & [Viper](https://github.com/spf13/viper) go packages handle the hard work for the CLI interface.


## Setup.

_TLDR:_ Just make sure the `gogs` executable winds up somewhere in your `$PATH`. Set your `GOGS_URL` and `GOGS_TOKEN` vars somehow.
<br>
<br>
Clone the repo and build and install it yourself, or `go get github.com/irstacks/go-gogs-cli`. 

Oh, _but how do you build it?_, you ask? _What path?_ you ask? 

Ok, fine. Here's what's up. From the beginning. 

With the marvelous `go get`...
```bash
$ go get github.com/irstacks/go-gogs-cli
$ cd $GOPATH/src/irstacks/go-gogs-cli
$ go get ./...
```
<br>
<br>
__Or__, with the almost-as-marvelous `git clone`...
```bash
$ cd where/i/like/to/put/funky/things
$ git clone https://github.com/irstacks/go-gogs-cli.git
# Pedantry explained immediately below...
$ go get stuff
$ go get morestuff
```
I say almost-as-marvelous because if you use `git clone` you may run into issues about your $GOPATH. It happens. Since I haven't figured out how to consistently `go get` all the dependencies I need for a given Go project from esoteric locations outside my $GOPATH (in part because I have so many esoteric locations and don't want to haggle with forever adjusting/amending my $GOPATH extensioners), I usually just wind up running `go run main.go` or `go build -o gogs` and `go get`ting the dependencies it complains about one-by-one. I know, it sucks. But that's just how I roll sometimes. 
<br>
<br>
Finally, we can build the sucker (this is __optional__ -- the repo comes with a build `gogs`).
```bash
$ go build -o gogs # Building it yourself will make sure the executable fits your OS and architecture. 
```
Here we're using -o to tell go where to build the build it makes, in this case a file in the same directory we're with the name `gogs` (because thats a lot shorter than 'go-gogs-cli'). Note that whatever you name this file is what it will be accessible for you as on the CLI. So if you name it 'goo' (awesome.), then your commands will be all like: `$ goo repo new ...` 

Now here you've got some options (probably) about in which of your $PATH's paths you want to stick it. I like to keep my custom thingeys out of the dredges, so I stick mine in `$HOME/bin/`
<br>
... If you follow in my footsteps, make sure somewhere in your bash/fish/zsh shell you've added $HOME/bin (NO SECOND SLASH, you slashing fiend you) to your $PATH, with something like `export $PATH="$PATH:$HOME/bin"`.
```bash
$ cp gogs ~/bin/
```

... Or let Go decide where to put the executable for you. Just make sure your `$GOPATH/bin` is actually in your `$PATH`.
```bash
# This builds and moves the executable to $GOPATH/bin. 
$ go install # From inside the base of the project.
```
<br>

#### This repo's build is for darwin (Mac). 
If you're on a Mac and want to build it for your server or something, you can build for linux with the nifty `env GOOS=linux go build -o gogs`. You can probably build for Windows too but I don't trouble myself with such things.

## Config.
Once you've got the project, you'll need to configure your own `GOGS_TOKEN` and `GOGS_URL` variables. 
<br>
You can use a file called `.go-gogs-cli.yaml` that likes to live at `$HOME/.go-gogs-cli.yaml`. It handles configuring your __Gogs url__ and __token__, like such:
```yaml
GOGS_TOKEN: 0e6709o05da4753dddf5f592374fdc263f02n801
GOGS_URL: http://my.goggers.me
```
Make that file (there's an example in the repo) and fill that in for your own self. 
<br>
<br>
__Or__, if you'd rather use environment variables, that'll work fine too. 
```bash
export GOGS_TOKEN=asdlfj239r029fzsfasf923r23f
export GOGS_URL=http://my.goggers.me
```

## Usage.
So far, you can do things. What's that? You can _do things_? Yep! Do things!
```bash
# Create basic: 
$ gogs repo create wheres-waldo # Create a repo owned by you.
# Create fancy: 
$ gogs repo create wheres-waldo --desc 'awesome stuff' --private --org GophersGophering
$ gogs repo create where-waldo -r origin
# [Aliases] for create. 
$ gogs repo [create|new|c|n]
# [Options] for create.
[-n | --name] # Name (or plain old args[0] also work, as above, obviously)
[-d | --desc] # Description
[-o | --org] # Owned by a an organization you own
[-p | --private] # Make repo private
[-r | --add-remote] # Add newly created gogs repo as a remote to your current git dir, initalizing git if necessary

# List basic:
$ gogs repo list # Get all yo repos.

# Search basic:
$ gogs repo search waldo # Search public repos for keyword 'waldo'.
# Search fancy:
$ gogs repo search waldo --limit 1 --user thatguy
# [Aliases] for search.
$ gogs repo [search|find|s|f]
# [Options] for search.
[-l | --limit] # Limit results
[-u | --user] # By user, required if you want to search private repos

# Destroy basic: 
$ gogs repo destroy irstacks my-exterminable-repo
$ gogs repo destroy irstacks/my-other-exterminable-repo
# [Aliases] for destroy.
$ gogs repo [destroy|delete|d|rid]

# Help?!
# Add --help after any command to see what's up, ie.
$ gogs --help
$ gogs repo --help
$ gogs repo create --help
# and so on...
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

<br>
<br>
You may have noticed that we're pretty heavy on the `gogs repo` and pretty light on the `gogs somethingelse` side of things. The [Gogs Client API](https://github.com/gogits/go-gogs-client) makes a bunch of endpoints and methods accessible for Users, Organizations, Issues, Admins, and so forth (although it's still very much a work in progress). Myself, I mostly just want to be able to create, search, and destroy like a fiend. If you would :heart: something and are unable to help yourself, let me know by opening an issue. 

## Help out.
:clap: [chanting] Do it! Do it! Do it!
