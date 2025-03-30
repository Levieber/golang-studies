// Writes an SVG clockface of the current time to Stdout.
package main

import (
	"os"
	"time"

	"levieber.com.br/golang-studies/maths/svg"
)

func main() {
	t := time.Now()
	svg.Write(os.Stdout, t)
}
