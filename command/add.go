package command

import (
	"fmt"
	"os"
	"path/filepath"
	"github.com/gernest/front"
	"bufio"
	"io/ioutil"
	patt "github.com/takaishi/patt/lib"
	"github.com/urfave/cli"
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

func readFrontMatter(pattern string) (string, patt.Config, error) {
	fp, err := os.Open(pattern)
	if err != nil {
		fmt.Println(err)
		return "", patt.Config{}, err
	}
	defer fp.Close()

	reader := bufio.NewReader(fp)
	m := front.NewMatter()
	m.Handle("+++", front.JSONHandler)
	f, _, err := m.Parse(reader)
	if err != nil {
		fmt.Println(err)
		return "", patt.Config{}, err
	}

	dst := f["destination"].(string)
	name := f["name"].(string)
	defer fp.Close()

	fm := patt.Config{
		Source: templatePath(pattern),
		Destination: dst,
	}
	return name, fm, nil
}

func configExists(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil
}

func templatePath(pattern string) string {
	return os.Getenv("HOME") + "/.patt.d/templates/" + filepath.Base(pattern)
}


func configFilePath() string {
	return os.Getenv("HOME") + "/.patt.d/config.json"
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

	name, fm, err := readFrontMatter(pattern)
	if err != nil {
		return err
	}

	if !configExists(configFilePath()) {
		fp, err := os.Create(configFilePath())
		if err != nil {
			fmt.Println(err)
		}
		writer := bufio.NewWriter(fp)
		_, err = writer.WriteString("{}")
		if err != nil {
			fmt.Println(err)

		}
		writer.Flush()

		fp.Close()
	}

	configs := patt.ReadConfig()
	configs[name] = fm

	err = patt.WriteConfig(configs)
	if err != nil {
		return err
	}

	return nil
}
