package main

import (
	"errors"
	"plugin"
)

func Parser(headers string, pluginer *plugin.Plugin) (err error, code int, resfunc func(string, []byte)) {
	srm, err := pluginer.Lookup("StatusCode")
	code = *srm.(*int)
	srm, err = pluginer.Lookup("RespParser")
	if err == nil {
		newfun, ok := srm.(func(string, []byte))
		if ok {
			resfunc = newfun
		} else {
			err = errors.New("cant find RespParser function")
		}
	}
	return
}
