package command

import (
	"strings"
	"io/ioutil"
	"fmt"
	"os"
	//"github.com/gernest/front"
	"text/template"
	"bytes"
	"bufio"
	"encoding/json"
	patt "github.com/takaishi/patt/lib"
	"github.com/gernest/front"
	"path/filepath"
	"regexp"
)

type NewCommand struct {
	Meta
}


type Foo struct {
	Year string
	Month string
	Day string
	Week string
}

//func configFilePath() string {
//	return os.Getenv("HOME") + "/.patt.d/config.json"
//}


func (c *NewCommand) Run(args []string) int {
	// Write your code here
	name := args[0]
	configs := patt.ReadConfig()
	src := configs[name].Source
	//dst := configs[name].Destination

	data := Foo{
		Year: "2017",
		Month: "01",
		Day: "05",
		Week: "Thu",
	}
	//
	tmpl, err := template.ParseFiles(src)
	if err != nil {
		fmt.Println(err)
	}
	var doc bytes.Buffer
	err = tmpl.Execute(&doc, data)
	if err != nil {
		fmt.Println(err)
	}


	m := front.NewMatter()
	m.Handle("+++", front.JSONHandler)
	f, body, err := m.Parse(&doc)
	if err != nil {
		fmt.Println(err)
	}

	dst := f["destination"].(string)
	base := filepath.Base(dst)
	rep := regexp.MustCompile(base + "$")
	dstDir := rep.ReplaceAllString(dst, "")
	err = os.MkdirAll(dstDir, 0755)
	if err != nil {
		fmt.Println(err)
		return 1
	}

	fp, err := os.Create(dst)
	if err != nil {
		fmt.Println(err)
	}
	defer fp.Close()

	writer := bufio.NewWriter(fp)
	_, err = writer.WriteString(body)
	if err != nil {
		fmt.Println(err)
	}
	writer.Flush()
	return 0
}

func (c *NewCommand) Synopsis() string {
	return ""
}

func (c *NewCommand) Help() string {
	helpText := `

`
	return strings.TrimSpace(helpText)
}
