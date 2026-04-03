package encoding

import (
	"strings"
)

// Codec defines the interface Transport uses to encode and decode messages.
type Codec interface {
	// Marshal returns the wire format of v.
	Marshal(v any) ([]byte, error)
	// Unmarshal parses the wire format into v.
	Unmarshal(data []byte, v any) error
	// Name returns the name of the Codec implementation.
	Name() string
}

var registeredCodecs = make(map[string]Codec)

// RegisterCodec registers the provided Codec for use with all Transport clients and
// servers.
func RegisterCodec(codec Codec) {
	if codec == nil {
		panic("cannot register a nil Codec")
	}
	if codec.Name() == "" {
		panic("cannot register Codec with empty string result for Name()")
	}
	contentSubtype := strings.ToLower(codec.Name())
	registeredCodecs[contentSubtype] = codec
}

// GetCodec gets a registered Codec by content-subtype, or nil if no Codec is
// registered for the content-subtype.
func GetCodec(contentSubtype string) Codec {
	return registeredCodecs[contentSubtype]
}
