package configfetcher

import (
	"fmt"
	"testing"
)

func TestPiIp(t *testing.T) {
	cnfg, err := GetPiIP()
	if err != nil {
		t.Fatalf("error fetching config %+v", err)
	}

	if cnfg == "" {
		t.Fatalf("cnfg is empty")

	}
	if cnfg != "192.168.1.3" {
		t.Fatalf("incorrect ip %s", cnfg)
	}
	fmt.Println("config ip is ", cnfg)
}
