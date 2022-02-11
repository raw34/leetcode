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

    bookCmd.Flags().StringVarP(&Name, "name", "", "All", "课程名称")

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
        fmt.Println("name is " + name)
        if name != "All" {
            saveBock(name)
        } else {
            names := []string{"database", "network", "os", "design_pattern"}
            for _, name := range names {
                saveBock(name)
            }
        }
    },
}

type Node struct {
    Id       string
    Title    string
    ParentId string
}

func saveBock(name string) {
    // 获取文件
    bookPath := fmt.Sprintf("data/%s.json", name)
    content, err := ioutil.ReadFile(bookPath)
    if err != nil {
        panic(err)
    }
    // 解析文件
    book := gjson.Parse(string(content))
    questions := book.Get("data.leetbookBookDetail.pages").Array()
    // 遍历题目，拼接表格
    table := "| Done | Id | Chapter | Title | Mark |\n|:----:|-------|---------|------|"
    nodes := map[string]*Node{}
    for _, question := range questions {
        id := question.Get("id").String()
        title := question.Get("title").String()
        pageType := question.Get("pageType").String()
        if pageType == "CHAPTER" {
            nodes[id] = &Node{
                Id:    id,
                Title: title,
            }
            continue
        }
        parentId := question.Get("parentId").String()
        nodes[id] = &Node{
            Id:       id,
            Title:    title,
            ParentId: parentId,
        }
        chapter := ""
        for nodes[parentId] != nil {
            chapter = fmt.Sprintf("%s-%s", nodes[parentId].Title, chapter)
            parentId = nodes[parentId].ParentId
        }
        chapter = strings.TrimSuffix(chapter, "-")
        table += fmt.Sprintf("\n| %s | %s | %s | %s |   |", "⬜", id, chapter, title)
        detailTitle := strings.Replace(title, "/", "|", -1)
        detailPath := fmt.Sprintf("book/%s/%s.md", name, detailTitle)
        err = ioutil.WriteFile(detailPath, []byte("# "+title), 0644)
        if err != nil {
            panic(err)
        }
    }
    // 写入题目索引文件
    indexPath := fmt.Sprintf("book/%s/index.md", name)
    err = ioutil.WriteFile(indexPath, []byte(table), 0644)
    if err != nil {
        panic(err)
    }
}
