package hackrf

import "errors"

var (
	errInvalidParam = errors.New("Invalid parameter")
	errInvalidLna   = errors.New("Invalid LNA value")
	errInvalidVga   = errors.New("Invalid VGA value")
	errInvalidTxVga = errors.New("Invalid TxVGA value")
)
