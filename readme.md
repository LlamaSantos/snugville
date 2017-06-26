# Snugville

Welcome snugs to Snugville.  It's only the snuggiest place on planet or and maybe... beyond.

## Setup

Clone the repo :) and run `brew install go`

* Note: If you have previous played with Go, we're going to migrate away from your prevous `GOPATH` to create a fresh environment to work from.  Check this by `echo $GOPATH` from the terminal to see if it is set to anything.

### Setting up your $GOPATH

In your (.bashrc, .profile, .zshrc, et al) add the following line `export GOPATH=/path/to/your/repo/snugville`, example `export GOPATH=$HOME/code/snugville`

To immediately load the change run `source ~/.zshrc` or whatever shell file you use.

Additionally, if you want to use any installed go binaries you'll want to add the `bin` folder in your `GOPATH` to your `PATH`

```
export GOPATH=$HOME/code/snugville
export PATH=$PATH:$GOPATH/bin
```

### Initializing

* cd `src/snugs`
* `chmod +x setup.sh`
* `./setup.sh`


### Visual Studio Code

Install the `Go` extension from `lukehoban`, this will make your experience with go beyond better.  The extension is found here `https://marketplace.visualstudio.com/items?itemName=lukehoban.Go`.  Enabling `go.useLanguageServer` in Settings will give additional information whilst go-coding.

After you've installed the go extension, go ahead and open the `src/snugs` folder.  Note you cannot open the project from the `snugville` root otherwise Visual Studio won't work properly.  Don't worry you don't need anything below the `src/snugs` folder anyways.  It may take a save on `main.go` to trigger the IDE to prompt you to install the Go toolchain.  When you do, click Install All and wait for magic to happen.

### What Now?

We code in go :)