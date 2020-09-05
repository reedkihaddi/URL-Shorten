package encode

import (
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"log"
	"math"
	"math/rand"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// For base62 hashing.
const base = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const b = 62

// HashLink hashes the link using SHA-1.
func HashLink(str string) string {

	// Create a random int and append it to the link to get unique hashes.
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	x := r.Intn(100000)

	// Concatenate string.
	var sb strings.Builder
	sb.WriteString(str)
	t := strconv.Itoa(x)
	sb.WriteString(t)

	// Use sha256 to hash.
	h := sha256.New()
	h.Write([]byte(sb.String()))
	sha := base64.URLEncoding.EncodeToString(h.Sum(nil))

	// Remove everything except alphanumeric.
	reg, err := regexp.Compile("[^a-zA-Z0-9]+")
	if err != nil {
		fmt.Println("Error in regex")
		log.Println(err)
	}
	replacedString := reg.ReplaceAllString(sha, "")

	// Convert to byte and return first six letters of the hash, might be a cleaner way to do this.
	b := []byte(replacedString)
	i := b[0:6]
	processedString := string(i[:])

	return processedString

}

//ToBase62  generates a base62 string.
func ToBase62(num int) string {
	r := num % b
	res := string(base[r])
	div := num / b
	q := int(math.Floor(float64(div)))

	for q != 0 {
		r = q % b
		temp := q / b
		q = int(math.Floor(float64(temp)))
		res = string(base[int(r)]) + res
	}
	return string(res)
}

//ToBase10 takes a base62 string and gives the original number.
func ToBase10(str string) int {
	res := 0
	for _, r := range str {
		res = (b * res) + strings.Index(base, string(r))
	}
	return res
}
