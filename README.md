# hash-go

Go abstraction for hashing a string

## Install

```bash
$ go get github.com/garavan/hash-go
```

## Usage

```go
import (
	hash "github.com/garavan/hash-go"
	bcryptHash "github.com/garavan/hash-go/bcrypt"
)

type userService struct {
	hash hash.Hash
}

func NewUserService(hash hash.Hash) User {
	return &userService{hash}
}

func (u *userService) Create(data map[string]string) error {
	passHashed, err := u.hash.Make(data["password"])
	...
}


userSrv := NewUserService(bcryptHash.DefaultBcryptHash())
...
```

## Methods

1. [bcrypt](https://godoc.org/golang.org/x/crypto/bcrypt)

## Test

```bash
$ go test ./...
```
