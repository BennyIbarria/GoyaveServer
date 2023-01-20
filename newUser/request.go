package main

import "goyave.dev/goyave/v4/validation"

var (
	PostRequest = validation.RuleSet{
		"Id":       validation.List{"required", "numeric"},
		"Name":     validation.List{"required", "string"},
		"LastName": validation.List{"required", "string"},
		"Age":      validation.List{"required", "numeric"},
		"Email":    validation.List{"required", "string"},
	}
)
