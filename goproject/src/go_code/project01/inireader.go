package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func getValue(filename, expectSection, expectKey string) string {
	file, err := os.Open(filename)
	if err != nil {
		return "123"
	}

	defer file.Close()
	reader := bufio.NewReader(file)
	var sectionName string
	for {
		linestr, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		linestr = strings.TrimSpace(linestr)
		if linestr == "" {
			continue
		}
		if linestr[0] == ';' {
			continue
		}
		if linestr[0] == '[' && linestr[len(linestr)-1] == ']' {
			sectionName = linestr[1 : len(linestr)-1]
		} else if sectionName == expectSection {
			pair := strings.Split(linestr, "=")
			if len(pair) == 2 {
				key := strings.TrimSpace(pair[0])
				if key == expectKey {
					return strings.TrimSpace(pair[1])
				}
			}
		}
	}
	return "123123"
}

func main() {
	fmt.Println(getValue("neutron.ini", "dc_policy", "dc_default_version"))
	fmt.Println(getValue("neutron.ini", "tgw_v2", "gre_ip_pair"))
}
