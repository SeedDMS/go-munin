package main

import (
    "fmt"
    "os"
    "net/url"
    "seeddms.org/seeddms/apiclient"
)

var DisplayConfig = false
var Target = ""
var Apikey = ""
var Mode = ""

func main() {
    if (len(os.Args) == 2 && os.Args[1] == "help") {
		fmt.Println("SeedDMS Munin PlugIn")
		fmt.Println("https://www.seeddms.org/")
		fmt.Println("(c)2022-2023 Uwe Steinmann")
		os.Exit(0)
	}

    if (len(os.Args) == 2 && os.Args[1] == "autoconf") {
		fmt.Println("yes")
		os.Exit(0)
	}

    if (len(os.Args) == 2 && os.Args[1] == "config") { DisplayConfig = true }

    // url of rest api
    Target = os.Getenv("target")
    // api key of rest api
    Apikey = os.Getenv("apikey")
    // mode is currently not used but is for getting other data than total stats
    Mode = os.Getenv("mode")
    if(Target == "") {
        Target = "http://localhost/restapi/index.php"
    }

    u, err := url.Parse(Target)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Could not parse target")
        os.Exit(-1)
    }

    switch Mode {
    default:
        if (DisplayConfig) {
            fmt.Println("graph_title", u.Host)
            fmt.Println("graph_vlabel #")
            fmt.Println("graph_info This graph show the number of documents, folders, and users")
            fmt.Println("graph_category seeddms")
            fmt.Printf("documents.label Documents\n")
            fmt.Printf("documents.draw LINE1\n")
            fmt.Printf("documents.info The total number of documents\n")
            fmt.Printf("folders.label Folders\n")
            fmt.Printf("folders.draw LINE1\n")
            fmt.Printf("folders.info The total number of folders\n")
            fmt.Printf("users.label Users\n")
            fmt.Printf("users.draw LINE1\n")
            fmt.Printf("users.info The total number of users\n")
        } else {
            c := apiclient.Connect(Target, Apikey)
            res, err := c.Statstotal()
            if err == nil {
                fmt.Printf("documents.value %d\n", res.Data.Docstotal)
                fmt.Printf("folders.value %d\n", res.Data.Folderstotal)
                fmt.Printf("users.value %d\n", res.Data.Userstotal)
            } else {
                fmt.Fprintf(os.Stderr, "Could not connect to %s", Target)
                os.Exit(-1)
            }
        }
    }
}
