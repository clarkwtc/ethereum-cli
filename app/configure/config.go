package configure

import "os"

type Config struct {
    URL string
}

func NewConfing() *Config {
    urlType := os.Getenv("BLOCK_CHAIN_URL_TYPE")

    url := GetURLByType(urlType)
    return &Config{url}
}

func GetURLByType(urlType string) string {
    var url string
    switch urlType {
    case "main":
        url = os.Getenv("BLOCK_MAIN_CHAIN_URL")
    default:
        url = os.Getenv("BLOCK_TEST_CHAIN_URL")
    }
    return url
}
