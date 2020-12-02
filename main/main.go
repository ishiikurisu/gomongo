package main

import "github.com/ishiikurisu/gomongo"

func main() {
    mongoUrl := "mongodb://localhost:27017"

    gomongo.Run(mongoUrl)
}
