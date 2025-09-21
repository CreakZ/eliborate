package errs

import "errors"

var ErrEntityNotFound = errors.New("entity not found")
var ErrNoDataSentToUpdate = errors.New("no data sent to update")
