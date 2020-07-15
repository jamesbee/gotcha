package main

import (
	"encoding/csv"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

func main() {
	records := ParseCsv()
	if len(records) == 0 {
		fmt.Println("Pokemon names not found, quit now.")
		os.Exit(1)
	}
	fmt.Printf("\nSuccessfully loaded %d pokemon names!\n\n", len(records))

	cmd := "n"
	r := rand.New(rand.NewSource(time.Now().Unix()))
	idx := r.Intn(len(records))

	for {
		rec := records[idx]
		fmt.Printf("Gocha! %s ( %s / %s ) at %s!\n\n", rec[1], rec[2], rec[3], rec[0])
		fmt.Print("Need another one? [y/N] ")
		fmt.Scanf("%s", &cmd)
		if strings.ToLower(cmd) == "y" {
			idx = r.Intn(len(records))
			cmd = "n"
			fmt.Println("")
		} else {
			break
		}
	}
	fmt.Println("Bye!")
}

func ParseCsv() [][]string {
	fi, err := os.Open("./name.csv")
	if err != nil {
		panic(err)
	}
	r := csv.NewReader(fi)
	records, err := r.ReadAll()
	if err != nil {
		panic(err)
	}
	return records
}
