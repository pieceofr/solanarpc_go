package solanarpc

import (
	"errors"
)

const (
	invalidHost           = "invalid host"
	emptyHost             = "empty host"
	notSupportedProtocol  = "not a supported protocol"
	accountNotExist       = "account not exist"
	idNotMatch            = "id not match"
	notJsonType           = "not a Json type"
	noContentType         = "no content type"
	jsonParseError        = "json parse error"
	idMismatch            = "id mismatch"
	unKnownBlock          = "unknow block"
	timeStampNotAvailable = "timestamp is not available for this block"
)

var (
	ErrInvalidHost           error
	ErrEmptyHost             error
	ErrNotSupportProtocol    error
	ErrAccountNotExist       error
	ErrIDNotMatch            error
	ErrNotJSONType           error
	ErrNOContentType         error
	ErrJSONParseError        error
	ErrIDMismatch            error
	ErrUnknownBlock          error
	ErrTimeStampNotAvailable error
)

func init() {
	ErrInvalidHost = errors.New(invalidHost)
	ErrEmptyHost = errors.New(emptyHost)
	ErrNotSupportProtocol = errors.New(notSupportedProtocol)
	ErrAccountNotExist = errors.New(accountNotExist)
	ErrIDNotMatch = errors.New(idNotMatch)
	ErrNotJSONType = errors.New(notJsonType)
	ErrNOContentType = errors.New(noContentType)
	ErrJSONParseError = errors.New(jsonParseError)
	ErrIDMismatch = errors.New(idMismatch)
	ErrUnknownBlock = errors.New(unKnownBlock)
	ErrTimeStampNotAvailable = errors.New(timeStampNotAvailable)
}
