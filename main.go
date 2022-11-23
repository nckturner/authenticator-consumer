package main

import (
	"fmt"

	"sigs.k8s.io/aws-iam-authenticator/pkg/token"
)

func main() {

	identity := token.Identity{}

	fmt.Printf("%#v", identity)
}
