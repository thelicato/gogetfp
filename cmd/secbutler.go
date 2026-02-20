package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/thelicato/gogetfp"
)

var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

func main() {
	country := flag.String("country", "", "Comma-separated country codes (e.g. US,GB,BR). Empty = any")
	timeout := flag.Float64("timeout", 1, "Timeout (seconds) for working-proxy check")
	random := flag.Bool("random", false, "Shuffle proxy list before selecting")
	anonym := flag.Bool("anonym", false, "Require anonymous proxies")
	elite := flag.Bool("elite", false, "Require elite proxies")
	google := flag.Bool("google", false, "Require 'Google' = yes (as per proxy source column)")
	https := flag.Bool("https", false, "Use HTTPS proxies (and https:// schema in output/check)")
	list := flag.Bool("list", false, "Print matching proxies without checking they work")
	limit := flag.Int("limit", 0, "Limit output count with --list (0 = no limit)")
	untested := flag.Bool("untested", false, "Return a single matching proxy without checking it works")
	plain := flag.Bool("plain", false, "Print host:port instead of scheme://host:port")
	showVersion := flag.Bool("version", false, "Print version info and exit")

	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "gogetfp - get free proxies (CLI for github.com/thelicato/gogetfp)\n\n")
		fmt.Fprintf(flag.CommandLine.Output(), "Usage:\n")
		fmt.Fprintf(flag.CommandLine.Output(), "  gogetfp [flags]\n\n")
		fmt.Fprintf(flag.CommandLine.Output(), "Examples:\n")
		fmt.Fprintf(flag.CommandLine.Output(), "  gogetfp --country US --https\n")
		fmt.Fprintf(flag.CommandLine.Output(), "  gogetfp --untested --random\n")
		fmt.Fprintf(flag.CommandLine.Output(), "  gogetfp --list --limit 20 --plain\n\n")
		fmt.Fprintf(flag.CommandLine.Output(), "Flags:\n")
		flag.PrintDefaults()
	}
	flag.Parse()

	if *showVersion {
		fmt.Printf("gogetfp %s (commit=%s, date=%s)\n", version, commit, date)
		return
	}

	cfg := gogetfp.FreeProxyConfig{
		CountryID: parseCountries(*country),
		Timeout:   *timeout,
		Random:    *random,
		Anonym:    *anonym,
		Elite:     *elite,
		Google:    *google,
		HTTPS:     *https,
	}

	fp := gogetfp.New(cfg)
	scheme := "http"
	if cfg.HTTPS {
		scheme = "https"
	}

	if *list {
		proxies, err := fp.GetProxyList()
		if err != nil {
			fatal(err)
		}
		if *limit > 0 && *limit < len(proxies) {
			proxies = proxies[:*limit]
		}
		for _, hp := range proxies {
			out := addScheme(scheme, hp)
			if *plain {
				out = stripScheme(out)
			}
			fmt.Println(out)
		}
		return
	}

	// default: working proxy
	var out string
	var err error
	if *untested {
		var hp string
		hp, err = fp.GetProxy()
		out = addScheme(scheme, hp)
	} else {
		out, err = fp.GetWorkingProxy()
	}
	if err != nil {
		fatal(err)
	}
	if *plain {
		out = stripScheme(out)
	}
	fmt.Println(out)
}

func parseCountries(s string) []string {
	s = strings.TrimSpace(s)
	if s == "" {
		return nil
	}
	parts := strings.Split(s, ",")
	out := make([]string, 0, len(parts))
	for _, p := range parts {
		p = strings.ToUpper(strings.TrimSpace(p))
		if p != "" {
			out = append(out, p)
		}
	}
	return out
}

func addScheme(scheme, hostport string) string {
	if strings.Contains(hostport, "://") {
		return hostport
	}
	return scheme + "://" + hostport
}

func stripScheme(s string) string {
	s = strings.TrimPrefix(s, "http://")
	s = strings.TrimPrefix(s, "https://")
	return s
}

func fatal(err error) {
	fmt.Fprintln(os.Stderr, "Error:", err)
	os.Exit(1)
}
