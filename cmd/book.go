package cmd

import (
    "fmt"
    "github.com/spf13/cobra"
    "github.com/tidwall/gjson"
    "io/ioutil"
    "strings"
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
        fmt.Println("name is " + name)
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
            detailTitle := filterTitle(node.Title)
            detailPath := fmt.Sprintf("book/%s/%s.md", name, detailTitle)
            writeFile(detailPath, []byte(fmt.Sprintf("# %s\n\n", node.Title)))
        }
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
        node := &QuestionNode{Id: id, Title: title}
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
    detailTitle := strings.Replace(title, "/", "|", -1)
    detailTitle = strings.Replace(title, " ", "", -1)
    return detailTitle
}

func writeFile(path string, content []byte) {
    err := ioutil.WriteFile(path, content, 0644)
    if err != nil {
        panic(err)
    }
}
