package configreader

import (
	"fmt"

	"github.com/spf13/viper"
)

type Interface interface {
	ReadConfig(cfg interface{})
}

type Options struct {
	ConfigFile string
}

type configReader struct {
	viper *viper.Viper
	opt   Options
}

func Init(opt Options) Interface {
	v := viper.New()
	v.SetConfigFile(opt.ConfigFile)
	v.SetConfigType("json")
	if err := v.ReadInConfig(); err != nil {
		panic(fmt.Errorf("fatal error found during reading file. err : %w", err))
	}

	c := &configReader{
		viper: v,
		opt:   opt,
	}

	return c
}

func (c *configReader) ReadConfig(cfg interface{}) {
	if err := c.viper.Unmarshal(&cfg); err != nil {
		panic(fmt.Errorf("fatal error found during unmarshaling config. err: %w", err))
	}
}
