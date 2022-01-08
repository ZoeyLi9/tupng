package id

import (
	"fmt"
	"github.com/google/uuid"
)

const (
	Version1 = 1
	Version2 = 2
	Version3 = 3
	Version4 = 4
	Version5 = 5
)


func NewUUID(version int, data []byte) (uuid.UUID, error) {
	switch version {
	case Version1:
		return uuid.NewUUID()
	case Version2:
		return uuid.NewDCEGroup()
	case Version3:
		return uuid.NewMD5(uuid.NameSpaceDNS, data), nil
	case Version4:
		return uuid.NewRandom()
	case Version5:
		return uuid.NewSHA1(uuid.NameSpaceOID, data), nil
	default:
		return uuid.UUID{},fmt.Errorf("u need to choose the version number in [1,5]")
	}

}
