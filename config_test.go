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
	fmt.Println("config ip is ", cnfg)
}

func TestPiPort(t *testing.T) {
	cnfg, err := GetPiPort()
	if err != nil {
		t.Fatalf("error fetching config %+v", err)
	}

	if cnfg == "" {
		t.Fatalf("cnfg is empty")

	}
	fmt.Println("config port is ", cnfg)
}
