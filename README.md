le_go
=============

## Installation

`go get github.com/rpmoore/le_go`

## Usage

To send data to Log Entries from Golang you create an instance of the client which implements the [Writer](http://golang.org/pkg/io/#Writer) interface.  This allows you to pass the client to the built in Golang [Logger](http://golang.org/pkg/log/#New).  Here is a concret example of creating the client and then using it in a Logger.

```go

package main

import (
    "log"

    "github.com/rpmoore/le_go"
)

func main() {
    // If you want to send your log data via SSL replace false with true
    le, err := logentries.NewLogEntriesWriter("token", false)

    if err != nil {
        log.Fatal(err)
    }

    Log := log.New(le, "Info: ", log.Ldate|log.Ltime)

    Log.Printf("My log message")

}

```

If you were successful in setting the logger up you should see a message similar to the following in your Log Entries log:

`08:58:09.420 Info: 2014/06/02 08:58:09 My log message`
