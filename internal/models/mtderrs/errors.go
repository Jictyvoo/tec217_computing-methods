package mtderrs

import (
	"errors"
	"strconv"
)

func ErrMaxIterations(iterations uint32) error {
	return errors.New(
		"failed to find root after " + strconv.Itoa(int(iterations)) + " iterations",
	)
}
