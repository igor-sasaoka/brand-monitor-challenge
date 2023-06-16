package front 

import(
    "golang.org/x/crypto/bcrypt"
)

func HashPassword(p string) (string, error) {
    b, err := bcrypt.GenerateFromPassword([]byte(p), 14)
    return string(b), err
}

func CheckPassword(p, h string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(h), []byte(p))
    return err == nil
}
