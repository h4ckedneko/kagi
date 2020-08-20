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

func main() {
	flag.StringVar(&env, "env", ".env", "Set the env file to use.")
	flag.IntVar(&size, "size", 32, "Set the size of key to generate.")
	flag.BoolVar(&print, "print", false, "Show the key instead of writing to file.")
	flag.Parse()

	if err := dotenv.Load(env); err != nil {
		log.Fatalln(err)
	}

	var err error
	defer func() {
		if err != nil {
			log.Fatalln(err)
		}
	}()

	switch flag.Arg(0) {
	case "generate":
		err = runGenerate()
	case "decode":
		err = runDecode()
	default:
		err = fmt.Errorf("unknown command")
	}
}

func runGenerate() error {
	key := kagi.New(size)
	if print {
		fmt.Println(key)
		return nil
	}

	ck := "APP_KEY"
	cv := regexp.QuoteMeta(os.Getenv(ck))

	// Replace the application from env file.
	r, err := regexp.Compile(ck + "=" + cv)
	if err != nil {
		return err
	}
	b, err := ioutil.ReadFile(env)
	if err != nil {
		return err
	}
	br := r.ReplaceAll(b, []byte(ck+"="+key))
	if err := ioutil.WriteFile(env, br, 0644); err != nil {
		return err
	}

	fmt.Println("Application key was successfully set.")
	return nil
}

func runDecode() error {
	keyd := kagi.Decode(os.Getenv("APP_KEY"))
	fmt.Println(string(keyd))
	return nil
}
