package beater

type ConfigSettings struct {
	Input JbConfig
}

type JbConfig struct {
	JsonFieldToEsType *string
}
