//************************************************************************//
// API "cellar": Application Resource Href Factories
//
// Generated with goagen v0.0.1, command line:
// $ goagen.exe
// --out=$(GOPATH)\src\github.com\goadesign\goa-cellar
// --design=github.com/goadesign/goa-cellar/design
// --pkg=app
//
// The content of this file is auto-generated, DO NOT MODIFY
//************************************************************************//

package app

import "fmt"

// AccountHref returns the resource href.
func AccountHref(accountID interface{}) string {
	return fmt.Sprintf("/cellar/accounts/%v", accountID)
}

// BottleHref returns the resource href.
func BottleHref(accountID, bottleID interface{}) string {
	return fmt.Sprintf("/cellar/accounts/%v/bottles/%v", accountID, bottleID)
}
