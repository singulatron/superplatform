// This script only exists to generate example IDs
package main

import (
	"fmt"
	"os"

	sdk "github.com/singulatron/superplatform/sdk/go"
)

func main() {
	prefix := os.Args[1]
	fmt.Println(sdk.Id(prefix))
}
