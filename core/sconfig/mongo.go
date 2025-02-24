package sconfig

type Mongo struct {
	Uri string `yaml:"uri" json:"uri" mapstructure:"uri"`
}
