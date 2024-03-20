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
		fmt.Println("Usage: format <package-directory>")
		return
	}
	var directory = osx.Args[1]

	// Parse the package file.
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
	var model = parser.ParseSource(source)

	// Reformat the package file.
	var formatter = pac.Formatter().Make()
	source = formatter.FormatModel(model)
	bytes = []byte(source)
	err = osx.WriteFile(packageFile, bytes, 0644)
	if err != nil {
		panic(err)
	}
}
