# restless

This is a schemaless rest api

## Prerequisites

- [GNU Make](https://www.gnu.org/software/make/)
- [Golang](https://golang.org/) ~= 1.12.4
- [Docker](https://www.docker.com/) ~= 19.03.1
- [Docker Compose](https://docs.docker.com/compose/) ~= 1.24.1

## Installation & setup guide (Linux/macOS)

### 1. Installing Go with [Homebrew](https://brew.sh/)

```sh
brew install go
```

**Note:** You can install Go by whichever mean you see fit. Homebrew is just our choice of preference.

### 2. Setup your [GOPATH](https://github.com/golang/go/wiki/GOPATH)

```sh
# Source this in your favorite shell .rc (.bashrc, .zshrc, etc...)
# or just export it on the shell you're currently working on
export GOPATH=$HOME/place/to/put/my/go/code
```

**Note:** If you're using IDE embedded test runners like [Goland](https://www.jetbrains.com/go/) configure this as an environment variable

```sh
# Make sure GO111MODULE=on is set on your working shell
# The process should be the same as how you setup your GOPATH (See section 2.)
export GO111MODULE=on

# You can use go built-in tool to clone the code
go get -d -v github.com/xuanit/testing

# Or man-handling it
mkdir -p $GOPATH/src/github.com/xuanit
git clone https://github.com/xuanit/testing.git $GOPATH/src/github.com/xuanit/testing
```

### 4. Installing dependencies (Optional)

We love localizing project level dependencies so now we're going to download them using a go module utility.

```sh
# Make sure you're inside of the project directory first
go mod vendor
```

We usually refer to this practice as **vendoring**, hence the name of the utility

### 5. Run the tests
