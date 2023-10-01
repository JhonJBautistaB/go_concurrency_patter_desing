package main

import "fmt"

type HashAlgorithm interface {
	Hash(p *PasswordProctector)
}

type PasswordProctector struct {
	user          string
	passwordName  string
	hashAlgorithm HashAlgorithm
}

func NewPasswordProtector(user string, passwordName string, hash HashAlgorithm) *PasswordProctector {
	return &PasswordProctector{
		user:          user,
		passwordName:  passwordName,
		hashAlgorithm: hash,
	}
}

func (p *PasswordProctector) SetHashAlgorithm(hash HashAlgorithm) {
	p.hashAlgorithm = hash
}

func (p *PasswordProctector) Hash() {
	p.hashAlgorithm.Hash(p)
}

type SHA struct{}

func (SHA) Hash(p *PasswordProctector) {
	fmt.Printf("Hashing using SHA for %s\n", p.passwordName)
}

type MD5 struct{}

func (MD5) Hash(p *PasswordProctector) {
	fmt.Printf("Hashing using MD5 for %s\n", p.passwordName)
}

func main() {
	sha := &SHA{}
	md5 := &MD5{}

	PasswordProctector := NewPasswordProtector("nestos", "gmail password", sha)
	PasswordProctector.Hash()
	PasswordProctector.SetHashAlgorithm(md5)
	PasswordProctector.Hash()
}
