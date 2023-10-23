package utils

import (
	"crypto/rand"
  "crypto/subtle"
  "encoding/base64"
  "fmt"
  "errors"
  "strings"

	"golang.org/x/crypto/argon2"
)

var (
    ErrInvalidHash         = errors.New("the encoded hash is not in the correct format")
    ErrIncompatibleVersion = errors.New("incompatible version of argon2")
)

type ArgonParams struct {
    Memory      uint32
    Iterations  uint32
    Parallelism uint8
    SaltLength  uint32
    KeyLength   uint32
}

func GenerateFromPassword(password string, params *ArgonParams) (encodedHash string, err error) {
  salt, err := generateRandomBytes(params.SaltLength)
  if err != nil {
    return "", err
  }

  hash := argon2.IDKey(
    []byte(password),
    salt,
    params.Iterations,
    params.Memory,
    params.Parallelism,
    params.KeyLength,
  )

  b64Salt := base64.RawStdEncoding.EncodeToString(salt)
  b64Hash := base64.RawStdEncoding.EncodeToString(hash)

  encodedHash = fmt.Sprintf(
    "$argon2id$v=%d$m=%d,t=%d,p=%d$%s$%s",
    argon2.Version,
    params.Memory,
    params.Iterations,
    params.Parallelism,
    b64Salt,
    b64Hash,
  )

  return encodedHash, nil
}

func ComparePasswordAndHash(password, encodedHash string) (match bool, err error) {
    params, salt, hash, err := decodeHash(encodedHash)
    if err != nil {
        return false, err
    }

    otherHash := argon2.IDKey(
      []byte(password),
      salt,
      params.Iterations,
      params.Memory,
      params.Parallelism,
      params.KeyLength,
    )

    if subtle.ConstantTimeCompare(hash, otherHash) == 1 {
        return true, nil
    }
    return false, nil
}

func generateRandomBytes(n uint32) ([]byte, error) {
  b := make([]byte, n)
  _, err := rand.Read(b)
  if err != nil {
    return nil, err
  }

  return b, nil
}

func decodeHash(encodedHash string) (params *ArgonParams, salt, hash []byte, err error) {
    vals := strings.Split(encodedHash, "$")
    if len(vals) != 6 {
        return nil, nil, nil, ErrInvalidHash
    }

    var version int
    _, err = fmt.Sscanf(vals[2], "v=%d", &version)
    if err != nil {
        return nil, nil, nil, err
    }
    if version != argon2.Version {
        return nil, nil, nil, ErrIncompatibleVersion
    }

    params = &ArgonParams{}
    _, err = fmt.Sscanf(
      vals[3],
      "m=%d,t=%d,p=%d",
      &params.Memory,
      &params.Iterations,
      &params.Parallelism,
    )
    if err != nil {
        return nil, nil, nil, err
    }

    salt, err = base64.RawStdEncoding.Strict().DecodeString(vals[4])
    if err != nil {
        return nil, nil, nil, err
    }
    params.SaltLength = uint32(len(salt))

    hash, err = base64.RawStdEncoding.Strict().DecodeString(vals[5])
    if err != nil {
        return nil, nil, nil, err
    }
    params.KeyLength = uint32(len(hash))

    return params, salt, hash, nil
}

