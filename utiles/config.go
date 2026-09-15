package utiles

import (
	"encoding/json"
	"os"
)

type Config struct {
	WorkerNum int `json:"worker_num"`
}

func ReadConf(path string) (Config, error) {
	var cfg Config
	data, err := os.ReadFile(path)
	if err != nil {
		return cfg, err
	}
	err = json.Unmarshal(data, &cfg)
	return cfg, err
}
