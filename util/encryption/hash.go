package encryption

import (
	"github.com/pkg/errors"
	"github.com/speps/go-hashids/v2"
)

// EncodeHashIDs hashes a given uint-id with given settings to a hash-code string.
func EncodeHashIDs(id uint, alphabet, salt string, minLength int) (string, error) {
	hashData := hashids.HashIDData{
		Alphabet:  alphabet,
		Salt:      salt,
		MinLength: minLength,
	}

	h, err := hashids.NewWithData(&hashData)
	if err != nil {
		return "", errors.Wrap(err, "failed to create hash maker")
	}

	hashCode, err := h.Encode([]int{int(id)})
	if err != nil {
		return "", errors.Wrap(err, "failed to encode hash")
	}

	return hashCode, nil
}

// DecodeHashIDs decodes a given hash-code with given settings to a uint-id.
func DecodeHashIDs(hashCode, alphabet, salt string, minLength int) (uint, error) {
	hashData := hashids.HashIDData{
		Alphabet:  alphabet,
		Salt:      salt,
		MinLength: minLength,
	}

	h, err := hashids.NewWithData(&hashData)
	if err != nil {
		return 0, errors.Wrap(err, "failed to create hash maker")
	}

	ids, err := h.DecodeWithError(hashCode)
	if err != nil {
		return 0, errors.Wrap(err, "failed to decode hash")
	}

	id := uint(ids[0])
	return id, nil

}
