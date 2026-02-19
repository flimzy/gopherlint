//golangcitest:config_path testdata/goimports.yml
package testdata

import (
	"fmt" // want "File is not properly formatted"
	"github.com/flimzy/gopherlint/pkg/config"
)

func Bar() {
	fmt.Print("x")
	_ = config.Config{}
}
