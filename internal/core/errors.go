package core

import "errors"

var (
    
    IncorrectMD5Length = errors.New("md5 must contain exactly 32 characters")
    IncorrectMD5Format = errors.New("md5 must contain only hexadecimal characters")
    IncorrectName      = errors.New("threat name cannot be empty")
    IncorrectDescription = errors.New("threat description cannot be empty")
    IncorrectSeverity  = errors.New("severity must be one of: low, medium, high, critical")
    IncorrectThreatType = errors.New("threat type cannot be empty")
)