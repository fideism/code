package config

// Setting 配置对象
var Setting setting

// setting 配置
type setting struct {
	Gocolly struct {
		RepositoryPath string `toml:"repository_path"`
	}
}
