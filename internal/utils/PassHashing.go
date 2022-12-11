package utils

import (
    // "fmt"
    "crypto/sha1"
    "encoding/hex"
)

const salt = "qwerty"

func GetPswrdHash(password string) (string) {

    hash := sha1.New()
    hash.Write([]byte(password))

    hash_sum := hash.Sum([]byte(salt))
    hash_str := hex.EncodeToString(hash_sum)

    return hash_str
}
