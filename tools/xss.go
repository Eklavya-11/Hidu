// Maybe another repo

package main

import (
        "bufio"
        "crypto/tls"
        "flag"
        "fmt"
        "io/ioutil"
        "net"
        "net/http"
        "net/url"
        "os"
        "strings"
        "regexp"
        "sync"
        "time"
)

func main() {
          
        var xsspayload string
        flag.StringVar(&xsspayload, "payload", "", "")
        flag.StringVar(&xsspayload, "p", "", "")

        var proxy string
        flag.StringVar(&proxy,"proxy", "0","")
        flag.StringVar(&proxy,"x", "0","")
          // Headers flag
        flag.Var(&headers, "headers", "")
        flag.Var(&headers, "H", "")
}
