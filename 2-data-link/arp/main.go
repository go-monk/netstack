package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"runtime"
	"strings"
)

type arpEntry struct {
	IP     string
	MAC    string
	Device string // or interface
}

func main() {
	iface, err := net.InterfaceByName("eth0")
	if err != nil {
		log.Fatal(err)
	}

	entries, err := linuxARPEntries(iface.Name)
	if err != nil {
		log.Fatal(err)
	}

	if len(entries) == 0 {
		fmt.Println("arp: no entries available")
		return
	}

	for _, entry := range entries {
		fmt.Printf("arp ip=%s mac=%s dev=%s\n",
			entry.IP, entry.MAC, entry.Device)
	}
}

func linuxARPEntries(device string) ([]arpEntry, error) {
	if runtime.GOOS != "linux" {
		return nil, nil
	}

	f, err := os.Open("/proc/net/arp")
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var entries []arpEntry
	scanner := bufio.NewScanner(f)
	first := true
	for scanner.Scan() {
		if first {
			first = false
			continue
		}

		fields := strings.Fields(scanner.Text())
		if len(fields) < 6 {
			continue
		}

		entry := arpEntry{
			IP:     fields[0],
			MAC:    fields[3],
			Device: fields[5],
		}

		if device == "" || entry.Device == device {
			entries = append(entries, entry)
		}
	}

	return entries, scanner.Err()
}
