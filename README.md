# gogetfp

`gogetfp` (**Go Get** a **F**ree **P**roxy) is a lightweight Go tool (both a CLI and a library) facilitating the retrieval of free proxies for seamless integration into your applications. With support for popular proxy sources like `free-proxy-list.net`, `sslproxies.org` and `us-proxy.org`. `gogetfp` offers customizable options, including **country filtering** and **timeout settings**. Enhance privacy, security, and access geo-restricted content effortlessly, making `gogetfp` the go-to solution for integrating free proxies into your Go projects.

Heavily inspired by the Python library [freeproxy](https://github.com/jundymek/free-proxy).

## 🚀 CLI

If you want to use it as a CLI you can run:

```sh
go install -v github.com/thelicato/gogetfp/cmd@latest
```

Here's the help of the tool:

```
gogetfp - Go Get a Free Proxy (https://github.com/thelicato/gogetfp)

Usage:
  gogetfp [flags]

Examples:
  gogetfp --country US --https
  gogetfp --untested --random
  gogetfp --list --limit 20 --plain

Flags:
  -anonym
        Require anonymous proxies
  -country string
        Comma-separated country codes (e.g. US,GB,BR). Empty = any
  -elite
        Require elite proxies
  -google
        Require 'Google' = yes (as per proxy source column)
  -https
        Use HTTPS proxies (and https:// schema in output/check)
  -limit int
        Limit output count with --list (0 = no limit)
  -list
        Print matching proxies without checking they work
  -plain
        Print host:port instead of scheme://host:port
  -random
        Shuffle proxy list before selecting
  -timeout float
        Timeout (seconds) for working-proxy check (default 1)
  -untested
        Return a single matching proxy without checking it works
  -version
        Print version info and exit
```
## 📚 Library

On the other hand, if you want to use it in your Golang project you can run the following:

```bash
go get github.com/thelicato/gogetfp@latest
```

After this command `gogetfp` library source will be in your current `go.mod`.

### ⚙️ Options

The options are basically the same provided by [freeproxy](https://github.com/jundymek/free-proxy):

| Name      | Type     | Example      | Default value |
| --------- | -------- | ------------ | ------------- |
| CountryID | []string | ['US', 'BR'] | []string{}    |
| Timeout   | float64  | 0.1          | 1             |
| Random    | bool     | True         | False         |
| Anonym    | bool     | True         | False         |
| Elite     | bool     | True         | False         |
| Google    | bool     | False        | False         |
| HTTPS     | bool     | True         | False         |

### 💡 Example

```go
package main

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/thelicato/gogetfp"
)

func main() {
	fp := gogetfp.New(gogetfp.FreeProxyConfig{})

	proxy, err := fp.GetWorkingProxy()
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Working Proxy:", proxy)
	}
}
```

## 🪪 License

_gogetfp_ is made with 🖤 and released under the [MIT LICENSE](https://github.com/thelicato/gogetfp/blob/main/LICENSE).
