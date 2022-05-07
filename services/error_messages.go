package services

import (
	"fmt"
)

/* Error messages */
var (
	notFoundError  = fmt.Errorf("oil bin not found")
	emptyDbError   = fmt.Errorf("there are no bins loaded")
	badParamsError = fmt.Errorf("not enough or invalid parameters")
)
