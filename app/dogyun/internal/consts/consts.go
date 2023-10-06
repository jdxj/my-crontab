package consts

const (
	Target = "https://vm.dogyun.com/server/create/product/data"
)

var (
	Header = map[string]string{
		"authority":          "vm.dogyun.com",
		"accept":             "application/json, text/javascript, */*; q=0.01",
		"accept-language":    "zh-CN,zh;q=0.9,en;q=0.8",
		"cache-control":      "no-cache",
		"content-type":       "application/x-www-form-urlencoded; charset=UTF-8",
		"dnt":                "1",
		"origin":             "https://vm.dogyun.com",
		"pragma":             "no-cache",
		"referer":            "https://vm.dogyun.com/server/create",
		"sec-ch-ua":          `"Not/A)Brand";v="99", "Google Chrome";v="115", "Chromium";v="115"`,
		"sec-ch-ua-mobile":   "?0",
		"sec-ch-ua-platform": `"Linux"`,
		"sec-fetch-dest":     "empty",
		"sec-fetch-mode":     "cors",
		"sec-fetch-site":     "same-origin",
		"user-agent":         "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/115.0.0.0 Safari/537.36",
		"x-requested-with":   "XMLHttpRequest",
	}
)
