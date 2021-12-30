package socket

import (
	"fmt"
	"regexp"
)

func MainRegexp() {
	reg, err := regexp.Compile("[0-9]{1}")
	fmt.Println(err, reg)
	fmt.Println(reg.FindAllString("0-a23-b5454", 2))

}
