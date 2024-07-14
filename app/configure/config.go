package configure

import (
    "fmt"
    "os"
    "strings"
)

type Config struct {
    URL string
}

func NewConfing() *Config {
    urlType := os.Getenv("BLOCK_CHAIN_URL_TYPE")

    url := GetURLByType(urlType)
    return &Config{url}
}

func GetURLByType(urlType string) string {
    return os.Getenv(fmt.Sprintf("BLOCK_%s_CHAIN_URL", strings.ToUpper(urlType)))
}
