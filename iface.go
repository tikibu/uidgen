package uidgen

type UIDGenerator interface {
	GenUID() uint64
}
