package initialize

import (
	"bufio"
	"groqai2api/global"
	"groqai2api/pkg/proxypool"
	"log/slog"
	"net/url"
	"os"
)

func InitProxy() {
	var proxies []string
	proxyUrl := os.Getenv("PROXY_URL")
	if proxyUrl != "" {
		proxies = append(proxies, proxyUrl)
	}

	if _, err := os.Stat("proxies.txt"); err == nil {
		file, _ := os.Open("proxies.txt")
		defer file.Close()
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			proxy := scanner.Text()
			parsedURL, err := url.Parse(proxy)
			if err != nil {
				slog.Warn("proxy url is invalid", "url", proxy, "err", err)
				continue
			}
			// Not a complete proxy link if port information is missing
			if parsedURL.Port() != "" {
				proxies = append(proxies, proxy)
			} else {
				continue
			}
		}
	}
	if len(proxies) == 0 {
		proxy := os.Getenv("http_proxy")
		if proxy != "" {
			proxies = append(proxies, proxy)
		}
	}
	global.ProxyPool = proxypool.NewIProxyIP(proxies)
}
