package utils

import (
	"context"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"regexp"

	"github.com/redis/go-redis/v9"
)

// The result of code generation.
type CodeGenerationResult struct {
	ShortCode string
	Result    CodeGenerationOutcome
	Err       error
}

// The outcome of code generation.
type CodeGenerationOutcome int

const (
	NewLink CodeGenerationOutcome = iota
	ExistingLink
)

// Error when there is a database error
type DBError string

func (e DBError) Error() string {
	return fmt.Sprintf("database error: %s", string(e))
}

// Error when there is an invalid URL
type InvalidURLError string

func (e InvalidURLError) Error() string {
	return fmt.Sprintf("invalid URL error: %s", string(e))
}

// Error when there is a hash collision
type HashCollisionError string

func (e HashCollisionError) Error() string {
	return fmt.Sprintf("hash collision error: %s", string(e))
}

// regular expression for URL validation
const urlRegex = `^(https?://)([\da-z.-]+)\.([a-z.]{2,6})([/\w.\?\&=-])*$`

// GenerateCode generates a shortcode for the given URL.
// If the URL already has a shortcode, the existing shortcode is returned.
// If the URL does not have a shortcode, a new shortcode is generated and
// returned. If an error occurs, an empty string and an error are returned.
func GenerateCode(url string, rdb *redis.Client) (*CodeGenerationResult, error) {
	if !isValidURL(url) {
		return nil, InvalidURLError(url)
	}

	// Hash the URL using the SHA-256 algorithm and a custom encoding scheme.
	// The resulting hash will be a byte slice that contains only lowercase and
	// uppercase letters as well as digits.
	//
	// We use a custom encoding scheme to represent the hash as a string with
	// only alphanumeric characters and no special characters. Specifically,
	// the encoding scheme uses all uppercase and lowercase letters as well as
	// digits, and duplicates the lowercase 'a' and uppercase 'A' characters to
	// ensure that the encoded string only contains alphanumeric characters.
	hash := createCustomHash(url)

	for i := 6; i < len(hash); i++ {
		// Check if start of hash already exists as shortcode in db
		val, err := rdb.HGet(context.Background(), fmt.Sprintf("shortcode:%s", hash[:i]), "url").Result()

		// If the shortcode does not exist, create a new entry in the db
		if err != nil {
			err := rdb.HSet(context.Background(), fmt.Sprintf("shortcode:%s", hash[:i]), "url", url, "clicks", 0).Err()
			if err != nil {
				return nil, DBError(err.Error())
			}
			return &CodeGenerationResult{ShortCode: hash[:i], Result: NewLink}, nil
		} else if val == url {
			return &CodeGenerationResult{ShortCode: hash[:i], Result: ExistingLink}, nil
		}
	}

	// If we reach this point, there is a hash collision.
	val, _ := rdb.HGet(context.Background(), hash, "url").Result()
	return nil, HashCollisionError(val)
}

// Check if the given URL is valid.
func isValidURL(url string) bool {
	// URL must be less than 250 characters
	if len(url) > 250 {
		return false
	}

	return regexp.MustCompile(urlRegex).MatchString(url)
}

// Hash the URL using the SHA-256 algorithm and return the hash as a string
// of hexadecimal digits.
func createHash(url string) string {
	hasher := sha256.New()
	hasher.Write([]byte(url))
	hash := hex.EncodeToString(hasher.Sum(nil))

	return hash
}

// Hash the URL using the SHA-256 algorithm and return the hash as a string
// of alphanumeric characters encoded using a custom encoding scheme.
func createCustomHash(url string) string {
	hasher := sha256.New()
	hasher.Write([]byte(url))

	// Custom base64 encoding alphabet that includes all uppercase and lowercase letters
	// as well as digits. Duplicate characters are used to ensure that the resulting
	// encoded string only contains alphanumeric characters and no special characters.
	// Specifically, the lowercase 'a' and uppercase 'A' characters are used twice each.
	// This custom encoding is only intended to be used to represent the hash as a string
	// and should not be used for general purpose encoding or decoding.
	encoder := base64.NewEncoding("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789aA").WithPadding(base64.NoPadding)
	hash := encoder.EncodeToString(hasher.Sum(nil))

	return hash

}
