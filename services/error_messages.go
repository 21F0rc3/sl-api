package services

import (
	"fmt"
)

/* Error messages */
var (
	NotFoundError  = fmt.Errorf("___ not found")
	EmptyDbError   = fmt.Errorf("there are no ___ loaded")
	BadParamsError = fmt.Errorf("not enough or invalid parameters")
)
