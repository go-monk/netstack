package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"text/tabwriter"
)

func main() {
	ifaces, err := net.Interfaces()
	if err != nil {
		log.Fatal(err)
	}

	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	format := "%v\t%v\n"
	fmt.Fprintf(w, format, "Name", "MAC")
	fmt.Fprintf(w, format, "----", "---")

	for _, iface := range ifaces {
		fmt.Fprintf(w, format,
			iface.Name, iface.HardwareAddr,
		)
	}
	w.Flush()
}
