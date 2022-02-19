package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"os/exec"
	"time"

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

var f *os.File

func main() {
	var err error
	fmt.Println(fmt.Sprintf(`%s/gitee.md`, config.Setting.Gocolly.RepositoryPath))
	f, err = os.OpenFile(fmt.Sprintf(`%s/gitee.md`, config.Setting.Gocolly.RepositoryPath), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println(`处理日志文件失败`, err.Error())
	}

	fmt.Println(io.WriteString(f, fmt.Sprintf("\n\n**%s**\n", time.Now().Format(`2006-01-02`))))

	defer f.Close()

	c := colly.NewCollector(colly.UserAgent(`Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/98.0.4758.82 Safari/537.36`), colly.MaxDepth(1))

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})
	c.OnResponse(func(r *colly.Response) {
		fmt.Println("Visited", r.Request.URL)
	})
	c.OnHTML("a[class='title project-namespace-path']", func(element *colly.HTMLElement) {
		dealProject(project{
			Href: element.Attr(`href`),
			Name: element.Text,
		})
	})

	c.OnScraped(func(r *colly.Response) {
		fmt.Println("Finished", r.Request.URL)

		// git提交
		execCmd(fmt.Sprintf(`cd %s && git add . && git commit -m '%s' && git push origin main`, config.Setting.Gocolly.RepositoryPath, time.Now().String()))
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
	fmt.Println(`开始处理`, data)
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

	if err := util.DelDir(fmt.Sprintf(`%s/.git`, dir)); nil != err {
		_ = fmt.Sprintf(`删除git文件目录失败:%s`, err.Error())
		return
	}

	// 保存记录
	writeReadme(data)

}

func writeReadme(data project) {

	content := fmt.Sprintf("- [%s](https://gitee.com%s)\n", data.Name, data.Href)

	fmt.Println(io.WriteString(f, content))
}

// execCmd xxx
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
