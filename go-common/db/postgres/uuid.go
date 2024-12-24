package postgres

import (
	"fmt"

	"github.com/jackc/pgx/v5/pgtype"
)

type Stringer[V any] interface {
	String() string
}

func ToUUIDs[V Stringer[V]](s []V) (uuidSlice []pgtype.UUID, err error) {
	for _, v := range s {
		var uuid pgtype.UUID
		err = uuid.Scan(v.String())
		uuidSlice = append(uuidSlice, uuid)
	}
	return
}

func UUIDToString(u pgtype.UUID) string {
	if !u.Valid {
		return ""
	}
	src := u.Bytes
	return fmt.Sprintf("%x-%x-%x-%x-%x", src[0:4], src[4:6], src[6:8], src[8:10], src[10:16])
}

func StringToUUID(s string) (pgtype.UUID, error) {
	var u pgtype.UUID
	err := u.Scan(s)
	return u, err
}
