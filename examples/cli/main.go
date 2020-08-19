package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"

	"github.com/h4ckedneko/kagi"
	dotenv "github.com/joho/godotenv"
)

var (
	env   string
	size  int
	print bool
)

func init() {
	flag.StringVar(&env, "env", ".env", "Set the env file to use.")
	flag.IntVar(&size, "size", 32, "Set the size of key to generate.")
	flag.BoolVar(&print, "print", false, "Show the key instead of writing to file.")
	flag.Parse()

	if err := dotenv.Load(env); err != nil {
		log.Fatalln(err)
	}
}

func main() {
	key := kagi.New(size)
	if print {
		fmt.Println(key)
		return
	}

	ck := "APP_KEY"
	cv := regexp.QuoteMeta(os.Getenv(ck))

	r, err := regexp.Compile(ck + "=" + cv)
	if err != nil {
		log.Fatalln(err)
	}
	b, err := ioutil.ReadFile(env)
	if err != nil {
		log.Fatalln(err)
	}
	br := r.ReplaceAll(b, []byte(ck+"="+key))
	if err := ioutil.WriteFile(env, br, 0644); err != nil {
		log.Fatalln(err)
	}

	fmt.Println("Application key was successfully set.")
}
