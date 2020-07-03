package main

import (
	"fmt"

	uuid "github.com/satori/go.uuid"
)

func main() {
	// uuid.Must(uuid.NewV3(uuid.), nil)
	u1 := uuid.Must(uuid.NewV4(), nil)
	// u1 := uuid.NewMD5(uuid.Must(uuid.NewRandom()), []byte(""))
	fmt.Printf("UUIDv4: %s\n", u1)
}
