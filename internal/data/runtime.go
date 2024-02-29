package data

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

type Runtime int32

var ErrInvalidRuntimeFormat = errors.New("invalid runtime format")

func (r Runtime) MarshalJSON() ([]byte, error) {
	duration := time.Duration(r) * time.Minute
	hours := int(duration.Hours())
	mins := int(duration.Minutes()) % 60

	formattedTime := fmt.Sprintf("%02d:%02d", hours, mins)
	jsonValue := fmt.Sprintf(`"%s mins"`, formattedTime)

	return []byte(jsonValue), nil
}

func (r *Runtime) UnmarshalJSON(jsonValue []byte) error {
	unquotedJSONValue, err := strconv.Unquote(string(jsonValue))
	if err != nil {
		return ErrInvalidRuntimeFormat
	}

	parts := strings.Split(unquotedJSONValue, " ")

	if len(parts) != 2 || parts[1] != "mins" {
		return ErrInvalidRuntimeFormat
	}

	i, err := strconv.ParseInt(parts[0], 10, 32)
	if err != nil {
		return ErrInvalidRuntimeFormat
	}

	*r = Runtime(i)

	return nil
}
