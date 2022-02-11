package cmd

import (
    "fmt"
    "github.com/spf13/cobra"
    "github.com/tidwall/gjson"
    "io/ioutil"
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
        if name == "All" {
            saveBock(name)
        } else {
            names := []string{"database", "network", "os", "design_pattern"}
            for _, name := range names {
                saveBock(name)
            }
        }
    },
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
    table := "| Done | Chapter | Title | Mark |\n|:----:|-------|---------|------|"
    chapters := map[string]string{}
    for _, question := range questions {
        id := question.Get("id").String()
        title := question.Get("title").String()
        pageType := question.Get("pageType").String()
        if _, ok := chapters[id]; !ok && pageType == "CHAPTER" {
            chapters[id] = title
            continue
        }
        pid := question.Get("parentId").String()
        chapter := chapters[pid]
        table += fmt.Sprintf("\n| %s | %s | %s |   |", "⬜", chapter, title)
    }
    // 写入题目索引文件
    indexPath := fmt.Sprintf("book/%s.md", name)
    err = ioutil.WriteFile(indexPath, []byte(table), 0644)
    if err != nil {
        panic(err)
    }
}
