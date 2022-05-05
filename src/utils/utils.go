package utils

import (
	"math/rand"
	"time"
	"strings"
)

var ascii_lower = "abcdefghijklmnopqrstuvwxyz"
var ascii_upper = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

func init() {
    rand.Seed(time.Now().UnixNano())
}

func randLetter() byte {
    return byte(rand.Intn(26) + 97)
}

func RandTimestamp() time.Time {
    t := time.Now()
    // select random date from next 10 years
    t = t.AddDate(rand.Intn(11), rand.Intn(13), rand.Intn(32))
    // Add random hours upto 24
    t = t.Add(time.Duration(rand.Intn(int(time.Second*60*60*24))))
    return t
}

func RandomString(n int) string {
    s := strings.Builder{}
    for i:=0; i<n; i++ {
        s.WriteByte(randLetter())
    }
    return s.String()
}

func RandomEmail() string {
    email := strings.Builder{}

    for i:=0; i<9; i++ {
        email.WriteByte(randLetter())
    }
    email.WriteByte('@')
    for i:=0; i<6; i++ {
        email.WriteByte(randLetter())
    }
    email.WriteByte('.')
    for i:=0; i<3; i++ {
        email.WriteByte(randLetter())
    }
    return email.String()
}
