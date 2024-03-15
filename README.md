<img src="https://craterdog.com/images/CraterDog.png" width="50%">

## Go Class-Based Package Framework Tools

### Overview
This project provides a set of Go based tools that can validate and format an
abstract `Model.go` file associated with a class-based package.  It also can
generate the corresponding concrete class files for each abstract class defined
in the package file.  For full details on the Go Class-Based Package Framework
click [here](https://github.com/craterdog/go-package-framework/wiki)

### Getting Started
To install the class-based package tools do the following from a terminal
window:
```
$ git clone git@github.com:craterdog/go-package-tools.git
$ cd go-package-tools/
$ ./etc/build.sh
$ ls
LICENSE		bin		go.mod		src
README.md	etc		go.sum

$ls bin/
format		generate	validate
```

### Using the Tools
The `validate` command reads in a `Model.go` file and ensures that it contains
only abstract type and interface definitions:
```
$ go-package-tools/bin/validate example/Model.go
```

The `format` command reads in a `Model.go` file and reformats it in its
canonical format as follows:
```
$ go-package-tools/bin/format example/Model.go
```

The `generate` command reads in a `Model.go` file and generates a separate
concrete class file—in the same directory—for each abstract class defined in the
`Model.go` file:
```
$ go-package-tools/bin/generate example/Model.go
```

<H5 align="center"> Copyright © 2009 - 2024  Crater Dog Technologies™. All rights reserved. </H5>
