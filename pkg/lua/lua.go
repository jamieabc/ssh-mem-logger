package lua

import (
	"github.com/yuin/gluamapper"
	l "github.com/yuin/gopher-lua"
)

func Parse(configPath string, obj interface{}) error {
	L := l.NewState()
	defer L.Close()

	L.OpenLibs()

	arg := &l.LTable{}
	arg.Insert(0, l.LString(configPath))
	L.SetGlobal("arg", arg)

	if err := L.DoFile(configPath); nil != err {
		return err
	}

	mapperOption := gluamapper.Option{
		NameFunc: func(s string) string {
			return s
		},
		TagName: "lua",
	}

	mapper := gluamapper.Mapper{Option: mapperOption}
	if err := mapper.Map(L.Get(L.GetTop()).(*l.LTable), obj); nil != err {
		return err
	}
	return nil
}
