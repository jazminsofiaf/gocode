# BC

Block Chain exploration

more info would be available later...

## Getting Started
Fist of all instal go tools 'https://golang.org/doc/install#install'

create a directory of your wish, for example `Users/john/go`
inside this directory create three folders:
* src (source code, all .go files)
* bin (compiled binary files)
* pkg (packages objects, all .a files)
clone this project inside src directory

it is important to set 3 environment variables in your .bashrc or .zshrc

`export GOPATH=/Users/john/go`
`export GOBIN=$GOPATH/bin`
`export PATH=$PATH:$GOPATH/bin`

you can check that the environment variables are set by command `go env GOPATH`
any time you import a package as `import "foo/bar"` it would be searched in:
 * $GOROOT (env var already set in installation)
 * $GOPATH

with the command 'go install' the program would be compiled and the binary would be save in $GOBIN path

as GOPATH/bin is added to PATH list variable you can execute the program wherever you want as `bcexplore-struct-sfr`
