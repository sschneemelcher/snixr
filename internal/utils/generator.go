package utils

import (
	"context"
	"crypto/sha512"
	"encoding/hex"
	"fmt"

	"github.com/redis/go-redis/v9"
)


func GenerateCode(url string, rdb *redis.Client) (string, error) {
    hash := CreateHash(url)

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

func CreateHash(url string) string {
    // Hash the url using SHA-512 algorithm
    hasher := sha512.New()
    hasher.Write([]byte(url))
    hash := hex.EncodeToString(hasher.Sum(nil))
    
    return hash

}
