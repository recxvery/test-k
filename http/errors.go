package http

import "errors"

var (
    incorrectMD5Length = errors.New("md5 must contain exactly 32 characters")
    incorrectMD5Format = errors.New("md5 must contain only hexadecimal characters")
    incorrectName      = errors.New("threat name cannot be empty")
    incorrectDescription = errors.New("threat description cannot be empty")
    incorrectSeverity  = errors.New("severity must be one of: low, medium, high, critical")
    incorrectThreatType = errors.New("threat type cannot be empty")
)