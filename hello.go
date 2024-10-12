package hello

import (
    "os"
    "rsc.io/quote"
)

func Hello() string {
    os.Setenv("LANG", "en")
    return quote.Hello()
}
