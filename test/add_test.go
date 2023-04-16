package test

import (
	"fmt"
	"sales-user-web/utils"
	"testing"
)

func TestGetPort(t *testing.T) {
	port, _ := utils.GetFreePort()
	fmt.Println(port)
}
