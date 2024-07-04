package sconfig

type Jaeger struct {
	Agent string `mapstructure:"agent" json:"agent" yaml:"agent"`
}
