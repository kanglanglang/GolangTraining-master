
package main




import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"


)



func getARPEntries() (map[string]uint32, error) {
	file, err := os.Open("/proc/net/arp")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	entries, err := parseARPEntries(file)
	if err != nil {
		return nil, err
	}

	return entries, nil
}

// TODO: This should get extracted to the github.com/prometheus/procfs package
// to support more complete parsing of /proc/net/arp. Instead of adding
// more fields to this function's return values it should get moved and
// changed to support each field.
func parseARPEntries(data io.Reader) (map[string]uint32, error) {
	scanner := bufio.NewScanner(data)
	entries := make(map[string]uint32)

	for scanner.Scan() {
		columns := strings.Fields(scanner.Text())

		if len(columns) < 6 {
			return nil, fmt.Errorf("unexpected ARP table format")
		}

		if columns[0] != "IP" {
			deviceIndex := len(columns) - 1
			entries[columns[deviceIndex]]++
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("failed to parse ARP info: %s", err)
	}

	return entries, nil
}




func main() {

	fmt.Println("ssss")
	entries, err := getARPEntries()

	fmt.Println("info", entries)
	fmt.Println(err)

}
