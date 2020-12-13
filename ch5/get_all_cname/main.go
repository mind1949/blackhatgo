package main

import (
	"flag"
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/miekg/dns"
)

func main() {
	var (
		flServerAddr = flag.String("server", "8.8.8.8:53", "The DNS server to use.")
		flFqdn       = flag.String("fqdn", "", "fqdn")
	)
	flag.Parse()
	if *flFqdn == "" {
		fmt.Println("-fqdn is required.")
		os.Exit(1)
	}

	var msg dns.Msg
	msg.SetQuestion(dns.Fqdn(*flFqdn), dns.TypeCNAME)
	in, err := dns.Exchange(&msg, *flServerAddr)
	if err != nil {
		fmt.Printf("lookup cname failed: %s.", err)
		os.Exit(1)
	}
	if len(in.Answer) == 0 {
		fmt.Println("no answers.")
		os.Exit(1)
	}

	w := tabwriter.NewWriter(os.Stdin, 0, 8, 0, ' ', 0)
	for _, rr := range in.Answer {
		if a, ok := rr.(*dns.CNAME); ok {
			fmt.Fprintf(w, "%s --<cname>--> %s\n", *flFqdn, a.Target)
		}
	}
	w.Flush()
}
