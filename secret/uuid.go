package secret

import "github.com/google/uuid"

/*
GenerateUUID - Generates a basic version 5 UUID to use in the header.Identifier field. The basis that is passed
in the parameter here is hashed along with the UUID namespace URL and a new UUID is generated from it. Using a
basis for this generation provides an additional layer of protection against duplication as if the same basis
is used, then the same UUID is generated
*/
func GenerateUUID(basis string) string {
	identifier := uuid.NewSHA1(uuid.NameSpaceURL, []byte(basis))

	return identifier.String()
}
