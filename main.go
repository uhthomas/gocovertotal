package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"

	"golang.org/x/tools/cover"
)

func percent(covered, total int64) float64 {
	if total == 0 {
		total = 1 // Avoid zero denominator.
	}
	return 100.0 * float64(covered) / float64(total)
}

func Main(ctx context.Context) error {
	name := flag.String("profile", "cover.out", "coverage profile")
	flag.Parse()

	f, err := os.Open(*name)
	if err != nil {
		return fmt.Errorf("open: %w", err)
	}
	defer f.Close()

	profiles, err := cover.ParseProfilesFromReader(f)
	if err != nil {
		return fmt.Errorf("parse profiles: %w", err)
	}

	var covered, total int64
	for _, p := range profiles {
		for _, block := range p.Blocks {
			total += int64(block.NumStmt)
			if block.Count > 0 {
				covered += int64(block.NumStmt)
			}
		}
	}
	fmt.Printf("total statement coverage: %f%%\n", percent(covered, total))
	return nil
}

func main() {
	if err := Main(context.Background()); err != nil {
		log.Fatal(err)
	}
}
