package main

import (
	"flag"
	"fmt"
	"os/exec"

	"github.com/fideism/code/internal/util"

	"github.com/fideism/code/internal/config"
	"github.com/gocolly/colly"
)

func init() {
	var conf string
	flag.StringVar(&conf, "c", "", "指定配置文件位置")
	flag.Parse()
	if conf == "" {
		panic("请指定配置文件位置")
	}
	config.Load(conf)

}

func main() {
	fmt.Println(config.Setting.Gocolly)
	c := colly.NewCollector(colly.UserAgent(`Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/98.0.4758.82 Safari/537.36`), colly.MaxDepth(1))

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})
	c.OnResponse(func(r *colly.Response) {
		fmt.Println("Visited", r.Request.URL)
	})
	c.OnHTML("a[class='title project-namespace-path']", func(element *colly.HTMLElement) {
		href := element.Attr(`href`)
		fmt.Println(fmt.Sprintf(`https://gitee.com%s.git`, href))
		fmt.Println(element.Text)
		fmt.Println(fmt.Sprintf(fmt.Sprintf(`%s%s`, config.Setting.Gocolly.RepositoryPath, href)))
		dealProject(project{
			Href: element.Attr(`href`),
			Name: element.Text,
		})
	})

	c.OnScraped(func(r *colly.Response) {
		fmt.Println("Finished", r.Request.URL)
	})

	// gitee 最新推荐
	if err := c.Visit(`https://gitee.com/explore/all?order=recommend`); nil != err {
		fmt.Println(err)
	}
}

type project struct {
	Name string
	Href string
}

func dealProject(data project) {
	fmt.Println(fmt.Sprintf(`https://gitee.com%s.git`, data.Href))
	fmt.Println(data.Name)
	dir := fmt.Sprintf(`%s%s`, config.Setting.Gocolly.RepositoryPath, data.Href)

	if err := util.DelDir(dir); nil != err {
		_ = fmt.Sprintf(`删除文件夹[%s]错误:%s\n`, dir, err.Error())
		return
	}

	if err := util.Mkdir(dir); nil != err {
		_ = fmt.Sprintf(`创建文件夹[%s]错误:%s\n`, dir, err.Error())
		return
	}

	// git clone https://gitee.com/theajack/cnchar.git /xxx/theajack/cnchar
	if err := execCmd(fmt.Sprintf(`git clone https://gitee.com%s.git %s`, data.Href, dir)); nil != err {
		_ = fmt.Sprintf(`执行git命令错误:%s`, err.Error())
		return
	}

}

// DelDir 删除文件夹
func execCmd(c string) error {
	cmd := exec.Command("/bin/bash", "-c", c)

	if err := cmd.Start(); nil != err {
		return err
	}

	if err := cmd.Wait(); nil != err {
		return err
	}

	return nil
}
