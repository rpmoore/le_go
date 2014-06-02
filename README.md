le_go
=============

# Installation

`go get github.com/rpmoore/le_go`

# Usage

To send data to Log Entries from Golang you create an instance of the client which implements the [Writer](http://golang.org/pkg/io/#Writer) interface.  This allows you to pass the client into the built in Golang [Logger](http://golang.org/pkg/log/#New).  Here is a concret example of creating the client and then using it in a Logger.

```go

import (
    "log"

    "github.com/rpmoore/le_go"
)

func main() {
    // If you want to send your log data via SSL replace false with true
    le = logentries.NewLogEntriesWriter("token", false)

    if err != nil {
        log.Fatal(err)
    }

    Log = log.New(le, "Info:", log.Ldate|log.Ltime)

    Log.Printf("My log message")

}

```
