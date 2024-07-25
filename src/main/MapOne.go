package main

import "fmt"

func main() {

	var sitemap map[string]string

	sitemap = make(map[string]string)

	sitemap["Google"] = "谷歌"
	sitemap["Runoob"] = "菜鸟教程"
	sitemap["Baidu"] = "百度"
	sitemap["Wiki"] = "维基百度"

	for site := range sitemap {
		fmt.Println(site, "首都是", sitemap[site])
	}

}
