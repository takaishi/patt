package command

import (
	"bufio"
	"fmt"
	"github.com/urfave/cli"
	"io/ioutil"
	"os"
	"path/filepath"
)

func createPattDir(pattern string) error {
	err := os.MkdirAll(templatePath(""), 0755)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func writePatternFile(pattern string) error {
	src, err := os.Open(pattern)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer src.Close()
	srcReader := bufio.NewReader(src)

	dst, err := os.Create(templatePath(pattern))
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer dst.Close()

	contents, _ := ioutil.ReadAll(srcReader)

	writer := bufio.NewWriter(dst)
	_, err = writer.Write(contents)
	if err != nil {
		fmt.Println(err)
		return err
	}
	writer.Flush()

	return nil

}

func templatePath(pattern string) string {
	return os.Getenv("HOME") + "/.patt.d/templates/" + filepath.Base(pattern)
}

func RunAddCommand(c *cli.Context) error {
	// Write your code here
	pattern := c.Args().Get(0)
	fmt.Printf("pattern = %v\n", pattern)

	err := createPattDir(pattern)
	if err != nil {
		return err
	}

	err = writePatternFile(pattern)
	if err != nil {
		return err
	}

	return nil
}
