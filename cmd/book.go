package cmd

import (
    "fmt"
    "github.com/gocolly/colly/v2"
    "github.com/spf13/cobra"
    "github.com/tidwall/gjson"
    "io/ioutil"
    "log"
    "strings"
    "time"
)

func init() {
    var Name string
    var Ftype string

    bookCmd.Flags().StringVarP(&Name, "name", "", "All", "课程名称")
    bookCmd.Flags().StringVarP(&Ftype, "ftype", "", "Index", "文件类型")

    rootCmd.AddCommand(bookCmd)
}

// bookCmd represents the book command
var bookCmd = &cobra.Command{
    Use:   "book",
    Short: "课程",
    Long:  `Leetcode课程。Usage：go run main.go book --name="database"`,
    Run: func(cmd *cobra.Command, args []string) {
        fmt.Println("book called")
        name := cmd.Flag("name").Value.String()
        ftype := cmd.Flag("ftype").Value.String()
        if name != "All" {
            saveBock(name, ftype)
        } else {
            names := []string{"database", "network", "os", "design_pattern"}
            for _, name := range names {
                saveBock(name, ftype)
            }
        }
    },
}

type QuestionNode struct {
    Id       string
    Title    string
    ParentId string
    Book     string
    Chapter  string
}

func saveBock(name string, ftype string) {
    // 获取题目信息
    questionList := getQuestions(name)

    // 写入题目索引文件
    if ftype == "Index" {
        table := "| Done | Id | Chapter | Title | Mark |\n|:----:|-------|---------|------|------|"
        for _, node := range questionList {
            table += fmt.Sprintf("\n| ⬜ | %s | %s | [%s](%s.md) |   |",
                node.Id,
                node.Chapter,
                node.Title,
                filterTitle(node.Title),
            )
        }
        indexPath := fmt.Sprintf("book/%s/index.md", name)
        writeFile(indexPath, []byte(table))
    }

    // 写入题目详情文件
    if ftype == "Detail" {
        for _, node := range questionList {
            saveQuestionDetail(node)
        }
    }
}

func saveQuestionDetail(node *QuestionNode) {
    detailTitle := filterTitle(node.Title)
    detailPath := fmt.Sprintf("book/%s/%s.md", node.Book, detailTitle)

    c := colly.NewCollector(
        colly.AllowedDomains("leetcode-cn.com"),
    )

    c.Limit(&colly.LimitRule{
        //DomainGlob:  "*httpbin.*",
        Parallelism: 2,
        RandomDelay: 5 * time.Second,
    })

    c.OnRequest(func(r *colly.Request) {
        fmt.Println("Visiting", r.URL)
        r.Headers.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/76.0.3809.100 Safari/537.36")
        r.Headers.Set("Content-Type", "application/json")
    })

    c.OnError(func(_ *colly.Response, err error) {
        log.Println("Something went wrong:", err)
    })

    c.OnResponse(func(r *colly.Response) {
        fmt.Println("Visited", r.Request.URL)
        blocks := gjson.Get(string(r.Body), "data.leetbookPage.blocks")
        content := ""
        for _, block := range blocks.Array() {
            blockType := block.Get("type").String()
            if blockType == "MARKDOWN" {
                content += block.Get("value").String() + "\n"
            }
        }
        content = fmt.Sprintf("# %s\n\n%s", node.Title, content)
        writeFile(detailPath, []byte(fmt.Sprintf(content, node.Title)))
    })

    url := "https://leetcode-cn.com/graphql/"
    payload := `{"operationName":"leetbookPageDetail","variables":{"pageId":"` + node.Id + `"},"query":"query leetbookPageDetail($pageId: ID!) {\n  leetbookPage(pageId: $pageId) {\n    title\n    subtitle\n    id\n    pageType\n    blocks {\n      type\n      value\n      __typename\n    }\n    commonTags {\n      nameTranslated\n      name\n      slug\n      __typename\n    }\n    qaQuestionUuid\n    ...leetbookQuestionPageNode\n    __typename\n  }\n}\n\nfragment leetbookQuestionPageNode on LeetbookQuestionPage {\n  question {\n    questionId\n    envInfo\n    judgeType\n    metaData\n    enableRunCode\n    sampleTestCase\n    judgerAvailable\n    langToValidPlayground\n    questionFrontendId\n    style\n    content\n    translatedContent\n    questionType\n    questionTitleSlug\n    editorType\n    mysqlSchemas\n    codeSnippets {\n      lang\n      langSlug\n      code\n      __typename\n    }\n    topicTags {\n      slug\n      name\n      translatedName\n      __typename\n    }\n    __typename\n  }\n  __typename\n}\n"}`
    err := c.PostRaw(url, []byte(payload))
    if err != nil {
        fmt.Println("Post err is :", err)
        return
    }
}

func getQuestions(name string) []*QuestionNode {
    // 获取文件
    bookPath := fmt.Sprintf("data/%s.json", name)
    content, err := ioutil.ReadFile(bookPath)
    if err != nil {
        panic(err)
    }
    // 解析文件
    book := gjson.Parse(string(content))
    questions := book.Get("data.leetbookBookDetail.pages").Array()
    // 遍历题目，获取题目信息
    questionList := make([]*QuestionNode, 0)
    questionMap := map[string]*QuestionNode{}
    for _, question := range questions {
        id := question.Get("id").String()
        title := question.Get("title").String()
        pageType := question.Get("pageType").String()
        node := &QuestionNode{Id: id, Title: title, Book: name}
        questionMap[id] = node
        if pageType == "CHAPTER" {
            continue
        }
        parentId := question.Get("parentId").String()
        node.ParentId = parentId
        node.Chapter = getChapter(questionMap, parentId)
        questionList = append(questionList, node)
    }
    return questionList
}

func getChapter(nodes map[string]*QuestionNode, parentId string) string {
    chapter := ""
    for nodes[parentId] != nil {
        chapter = fmt.Sprintf("%s-%s", nodes[parentId].Title, chapter)
        parentId = nodes[parentId].ParentId
    }
    chapter = strings.TrimSuffix(chapter, "-")
    return chapter
}

func filterTitle(title string) string {
    detailTitle := strings.ReplaceAll(title, "/", "|")
    detailTitle = strings.ReplaceAll(detailTitle, " ", "")
    return detailTitle
}

func writeFile(path string, content []byte) {
    err := ioutil.WriteFile(path, content, 0644)
    if err != nil {
        panic(err)
    }
}
