package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"

	"github.com/jackpal/bencode-go"
)

type Resume struct {
	Name          string
	Destination   string
	IncompleteDir string `bencode:"incomplete-dir"`

	Progress struct {
		Have string
	}
}

func printThing(r io.Reader) error {
	var res Resume
	if err := bencode.Unmarshal(r, &res); err != nil {
		return err
	}
	dest := filepath.Join(res.Destination, res.Name)
	if res.Progress.Have != "all" {
		dest = filepath.Join(res.IncompleteDir, res.Name)
	}

	fmt.Println(dest)
	return nil
}

func main() {
	flag.Parse()

	if flag.NArg() > 0 {
		for _, fl := range flag.Args() {
			fd, err := os.Open(fl)
			if err != nil {
				log.Fatal(err)
			}
			err = printThing(fd)
			fd.Close()
			if err != nil {
				log.Fatal(err)
			}
		}
	} else {
		bufred := bufio.NewReader(os.Stdin)
		for {
			if err := printThing(bufred); err != nil {
				if errors.Is(err, io.EOF) {
					break
				} else {
					log.Fatal(err)
				}

			}
		}

	}
}
