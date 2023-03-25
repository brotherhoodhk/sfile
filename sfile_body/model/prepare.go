package model

import (
	"encoding/xml"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"plugin"
	"strings"

	"github.com/oswaldoooo/octools/toolsbox"
)

type plugincnf struct {
	XMLName xml.Name     `xml:"plugin_conf"`
	Plugin  []plugininfo `xml:"plugin"`
}
type plugininfo struct {
	XMLName   xml.Name `xml:"plugin"`
	ClassName string   `xml:"classname,attr"`
	Name      string   `xml:"name,attr"`
}

var ParserFunc func(string, *plugin.Plugin) (error, int, func(string, []byte)) = DefaultParser //解析插件的方法，用插件解析插件的方式，实现高度自由化开发.只负责respstatus的插件解析
var ROOTPATH = os.Getenv("SFILE_HOME")
var errorlog = toolsbox.LogInit("error", ROOTPATH+"/log/error.log")

func init() {
	cnflist, err := toolsbox.ParseList(ROOTPATH + "/conf/site.cnf")
	if filename, ok := cnflist["plugincnf"]; err == nil && ok {
		err = loadparser()
		if err == nil {
			ReadPluginConf(filename)
		} else {
			errorlog.Println("extension function is ban,error>>", err.Error())
		}
	}

}
func ReadPluginConf(filename string) {
	if !strings.Contains(filename, ".xml") {
		filename = filename + ".xml"
	}
	content, err := ioutil.ReadFile(ROOTPATH + "/conf/" + filename)
	if err == nil {
		cnf := new(plugincnf)
		err = xml.Unmarshal(content, cnf)
		if err == nil {
			bad := 0
			for _, info := range cnf.Plugin {
				pluginer, err := toolsbox.ScanPluginByName(info.Name, ROOTPATH+"/plugin/")
				if err == nil {
					err, code, newfunc := ParserFunc(strings.ToLower(info.ClassName), pluginer)
					//check the function is whether existed
					if _, ok := SpecialParser[code]; err == nil && !ok {
						SpecialParser[code] = newfunc
					}
				}
				if err != nil {
					//write error into errorlog
					errorlog.Println(err)
					bad++
				}
			}
			errorlog.Printf("read %v plugin info,bad info %v\n", len(cnf.Plugin), bad)
		} else {
			fmt.Printf("read conf file failed,%v\n", err.Error())
		}
	}
}
func loadparser() (err error) {
	pluginer, err := toolsbox.ScanPluginByName("parser", ROOTPATH+"/plugin/")
	if err == nil {
		srm, err := pluginer.Lookup("Parser")
		if err == nil {
			newfunc, ok := srm.(func(string, *plugin.Plugin) (error, int, func(string, []byte)))
			if !ok {
				err = errors.New("cant match the parser function")
			} else {
				ParserFunc = newfunc
			}
		}
	}
	return
}
func DefaultParser(classname string, pluginer *plugin.Plugin) (err error, code int, resfunc func(string, []byte)) {
	srm, err := pluginer.Lookup("StatusCode")
	if err == nil {
		code = *srm.(*int)
		srm, err = pluginer.Lookup("RespParser")
		if newfunc, ok := srm.(func(string, []byte)); err == nil && ok {
			resfunc = newfunc
		} else {
			err = errors.New("cant find respparser function in plugin file")
		}
	}
	return
}
