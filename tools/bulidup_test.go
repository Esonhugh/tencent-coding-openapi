package tools

import (
	"bytes"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/bashtian/jsonutils"
	"github.com/guonaihong/gout"
	"io/fs"
	"os"
	"strings"
	"testing"
	"text/template"
)

var (
	outputDir = "./output"
)

var (
	packageNewer = template.Must(template.ParseFiles("struct.tmpl"))
	api          = template.Must(template.ParseFiles("api.tmpl"))
	openAPI      = template.Must(template.ParseFiles("openAPI.tmpl"))
)

var chinese2english = map[string]string{
	"团队":     "team",
	"项目":     "project",
	"项目设置":   "projectSetting",
	"需求":     "requirement",
	"事项":     "issue",
	"迭代":     "iteration",
	"工时":     "work",
	"代码托管":   "git",
	"提交":     "commit",
	"标签":     "tag",
	"版本":     "release",
	"分支":     "branch",
	"保护分支":   "protectedBranch",
	"保护分支规则": "branchProtection",
	"合并请求":   "mergeReq",
	"持续集成":   "ci",
	"持续部署":   "cd",
	"制品仓库":   "artifact",
	"测试管理":   "test",
	"文档管理":   "wiki",
}

type openApiInfo struct {
	Info []packageInfo
}

type packageInfo struct {
	PackageName      string
	PackageNameLower string
	PackageDesc      string
	PackageNameShort string // 包首字母小写
}

type actionInfo struct {
	ReqStruct  string
	RespStruct string
	ActionName string // 英文
	ActionDesc string // 中文
	packageInfo
}

func TestTool(t *testing.T) {
	var body string
	err := gout.GET("https://coding.net/help/openapi").BindBody(&body).Do()
	if err != nil {
		t.Fatal(err)
	}
	page, err := goquery.NewDocumentFromReader(strings.NewReader(body))
	if err != nil {
		t.Fatal(err)
	}

checkOutputDir:
	if _, err := os.Stat(outputDir); os.IsNotExist(err) {
		_ = os.Mkdir(outputDir, fs.ModePerm)
		goto checkOutputDir
	}

	allPkgInfo := openApiInfo{Info: make([]packageInfo, 0, 20)}

	section := page.Find("article > div:nth-of-type(1) > div:nth-of-type(3) > div > section")
	section.Each(func(i int, selection *goquery.Selection) {
		category := selection.Find("section > div:nth-of-type(1) h1").Contents().Text()
		categoryDesc := category
		category = chinese2english[category]
		apiDoc := selection.Find("section > div:not(:nth-of-type(1))")

	checkPackageDir:
		if _, err := os.Stat(outputDir + "/" + category); os.IsNotExist(err) {
			_ = os.Mkdir(outputDir+"/"+category, fs.ModePerm)
			goto checkPackageDir
		}

		fmt.Println(category)
		pkgInfo := packageInfo{
			PackageName:      strings.Title(category),
			PackageNameLower: strings.ToLower(category[:1]) + category[1:],
			PackageDesc:      categoryDesc,
			PackageNameShort: strings.ToLower(category[:1]),
		}
		allPkgInfo.Info = append(allPkgInfo.Info, pkgInfo)
		file, err := os.Create(outputDir + "/" + category + "/" + category + ".go")
		if err != nil {
			return
		}
		err = packageNewer.Execute(file, pkgInfo)
		if err != nil {
			t.Fatal(err)
			return
		}
		apiDoc.Each(func(i int, apiSelection *goquery.Selection) {
			apiName := apiSelection.Find("div > div:nth-of-type(1) > h2 > span").Text()
			actionName := strings.ReplaceAll(apiSelection.Find("div > div:nth-of-type(1) > div:nth-of-type(2) > span:nth-of-type(2)").Text(), "/open-api?Action=", "")
			apiReqExample := apiSelection.Find("div > div:nth-of-type(2) > div > div:nth-of-type(1) > pre > code").Contents().Not("span.comment").Text()
			apiRespExample := apiSelection.Find("div > div:nth-of-type(2) > div > div:nth-of-type(2) > pre > code").Contents().Not("span.comment").Text()
			fmt.Println("	" + apiName)

			file, err := os.Create(outputDir + "/" + category + "/" + strings.ToLower(actionName[:1]) + actionName[1:] + ".go")
			if err != nil {
				return
			}
			reqStruct := ""
			if apiReqExample != "" {
				reqStruct = json2struct(apiReqExample, actionName+"Req")
			}
			respStruct := ""
			if apiRespExample != "" {
				respStruct = json2struct(apiRespExample, actionName+"Resp")
			}
			err = api.Execute(file, actionInfo{
				ReqStruct:   reqStruct,
				RespStruct:  respStruct,
				ActionName:  strings.Title(strings.TrimSpace(actionName)),
				ActionDesc:  apiName,
				packageInfo: pkgInfo,
			})
			if err != nil {
				t.Fatal(err)
				return
			}
		})
	})

	file, err := os.Create(outputDir + "/openAPI.go")
	if err != nil {
		return
	}
	openAPI.Execute(file, allPkgInfo)
}

func json2struct(json string, jsonName string) string {
	reqModel, _ := jsonutils.ParseJson([]byte(json))
	buf := new(bytes.Buffer)
	jsonutils.WriteGo(buf, reqModel, jsonName)
	return buf.String()
}
