package schedule

type NewsApi struct {
	HotApi struct {
		Zhihu        string `yaml:"zhihu"`
		Weibo        string `yaml:"weibo"`
		Wangyi       string `yaml:"wangyi"`
		Toutiao      string `yaml:"toutiao"`
		Tengxun      string `yaml:"tengxun"`
		Lol          string `yaml:"lol"`
		Juejin       string `yaml:"juejin"`
		Hyper        string `yaml:"hyper"`
		Douyin       string `yaml:"douyin"`
		BilibiliRank string `yaml:"bilibili_rank"`
		Bilibili     string `yaml:"bilibili"`
		Baidu        string `yaml:"baidu"`
	} `yaml:"hot_api"`
}
