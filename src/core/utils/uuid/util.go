package uuid

import "github.com/google/uuid"

func GenerateSHA1UUID(s string) uuid.UUID {
	namespace := uuid.NewSHA1(uuid.NameSpaceDNS, []byte("gamen"))
	return uuid.NewSHA1(namespace, []byte(s))
}
