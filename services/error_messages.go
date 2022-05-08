package services

import (
	"fmt"
)

/* Error messages */
var (
	notFoundError  = fmt.Errorf("___ not found")
	emptyDbError   = fmt.Errorf("there are no ___ loaded")
	badParamsError = fmt.Errorf("not enough or invalid parameters")
)
