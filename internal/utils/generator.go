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

// regular expression for URL validation
const urlRegex = `^(https?://)([\da-z.-]+)\.([a-z.]{2,6})([/\w.\?\&=-])*$`


func GenerateCode(url string, rdb *redis.Client) (string, error) {
    if !isValidURL(url) {
        return "", fmt.Errorf("not a valid URL")
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
        val, err := rdb.Get(context.Background(), fmt.Sprintf("shortcode:%s", hash[:i])).Result()

        if err != nil {
            err := rdb.Set(context.Background(), fmt.Sprintf("shortcode:%s", hash[:i]), url, 0).Err()
            if err != nil {
                return "", fmt.Errorf("error generating code: %s", err)
            }
            return hash[:i], nil
        } else if val == url {
            return hash[:i], nil
        }
    }
    
    val, _ := rdb.Get(context.Background(), hash).Result()
    return "", fmt.Errorf("error generating code: hash collision for %s and %s", url, val)
}


func isValidURL(url string) bool {
    // URL must be less than 250 characters
    if len(url) > 250 {
        return false
    }

    return regexp.MustCompile(urlRegex).MatchString(url)
}


func createHash(url string) string {
    hasher := sha256.New()
    hasher.Write([]byte(url))
    hash := hex.EncodeToString(hasher.Sum(nil))

    return hash
}

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
