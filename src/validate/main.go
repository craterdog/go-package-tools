/*
................................................................................
.    Copyright (c) 2009-2024 Crater Dog Technologies.  All Rights Reserved.    .
................................................................................
.  DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               .
.                                                                              .
.  This code is free software; you can redistribute it and/or modify it under  .
.  the terms of The MIT License (MIT), as published by the Open Source         .
.  Initiative. (See http://opensource.org/licenses/MIT)                        .
................................................................................
*/

package main

import (
	fmt "fmt"
	pac "github.com/craterdog/go-package-framework/v2"
	osx "os"
	sts "strings"
)

// MAIN PROGRAM

func main() {
	// Validate the commandline arguments.
	if len(osx.Args) < 2 {
		fmt.Println("Usage: validate <package-directory>")
		return
	}
	var directory = osx.Args[1]

	// Validate the package file.
	if !sts.HasSuffix(directory, "/") {
		directory += "/"
	}
	var packageFile = directory + "Model.go"
	var bytes, err = osx.ReadFile(packageFile)
	if err != nil {
		panic(err)
	}
	var source = string(bytes)
	var parser = pac.Parser().Make()
	var package_ = parser.ParseSource(source)
	var validator = pac.Validator().Make()
	validator.ValidatePackage(package_)
}
