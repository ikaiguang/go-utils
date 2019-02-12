package gouuid

import "github.com/rs/xid"

// New new uuid
func New() xid.ID {
	return xid.New()
}

// NewUUID uuid string
func NewUUID() string {
	return xid.New().String()
}

// ParseUUID parse uuid
func ParseUUID(uuid string) (xid.ID, error) {
	return xid.FromString(uuid)
}
