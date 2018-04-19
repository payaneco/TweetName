package main

import (
	"gopkg.in/kyokomi/emoji.v1"
	"os"
        "fmt"
        "io/ioutil"
)

func main() {
        text, _ := ioutil.ReadAll(os.Stdin)
	TweetRename(string(text))
}

func TweetRename(text string) {
        fmt.Println(text)
	s := emoji.Sprint(text)
        emoji.Println(":beer:")
	Rename(s)
}
