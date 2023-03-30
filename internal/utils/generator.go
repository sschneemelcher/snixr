package utils

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"regexp"

	"github.com/redis/go-redis/v9"
)


func GenerateCode(url string, rdb *redis.Client) (string, error) {
    if !isValidURL(url) {
        return "", fmt.Errorf("not a valid URL")
    }

    hash := createHash(url)

    for i := 5; i < len(hash); i++ {
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

const urlRegex = `^(https?://)([\da-z.-]+)\.([a-z.]{2,6})([/\w.-]*)*/?$`

func isValidURL(url string) bool {
    // regular expression for URL validation
    return regexp.MustCompile(urlRegex).MatchString(url)
}

func createHash(url string) string {
    // Hash the url using SHA-512 algorithm
    hasher := sha256.New()
    hasher.Write([]byte(url))
    hash := hex.EncodeToString(hasher.Sum(nil))
    
    return hash

}
