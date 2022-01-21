package cmd

import (
    "fmt"
    "github.com/spf13/cobra"
    "github.com/tidwall/gjson"
    "github.com/yuin/goldmark"
    "github.com/yuin/goldmark/text"
    "io/ioutil"
    "os"
    "strconv"
    "strings"
)

func init() {
    var Name string

    topicCmd.Flags().StringVarP(&Name, "name", "", "", "标签名称")
    _ = topicCmd.MarkFlagRequired("name")

    rootCmd.AddCommand(topicCmd)
}

// topicCmd represents the topic command
var topicCmd = &cobra.Command{
    Use:   "topic",
    Short: "标签名称",
    Long:  `题目分类标签名称。Usage：go run main.go topic --name="Hash Table"`,
    Run: func(cmd *cobra.Command, args []string) {
        fmt.Println("topic called")
        name := cmd.Flag("name").Value.String()
        fmt.Println("name is", name)
        // 读取README文件（可选）
        readmePath := "README.md"
        content, err := ioutil.ReadFile(readmePath)
        if err != nil {
            panic(err)
        }
        // 解析README文件（可选）
        markdown := goldmark.New()
        readme := (markdown.Parser()).Parse(text.NewReader(content))
        // 遍历README文件，获取已完成题目编号
        readme = readme.FirstChild()
        questionsDone := map[int]bool{}
        for readme.NextSibling() != nil {
            readme = readme.NextSibling()
            kind := readme.Kind().String()
            if kind == "Paragraph" {
                child := readme.FirstChild().NextSibling()
                for child.NextSibling() != nil {
                    child = child.NextSibling()
                    row := strings.Split(string(child.Text(content)), "|")
                    done, no := strings.Trim(row[1], " "), strings.Trim(row[2], " ")
                    if done == "✅" {
                        i, _ := strconv.Atoi(no)
                        questionsDone[i] = true
                    }
                }
                break
            }
        }

        // 读取全部题目文件
        questionsPath := "questions.json"
        content, err = ioutil.ReadFile(questionsPath)
        if err != nil {
            panic(err)
        }
        // 解析全部题目文件
        questions := gjson.Parse(string(content)).Array()
        // 创建标签文件
        topicPath := "topic/" + strings.Replace(name, " ", "", -1) + ".md"
        title := fmt.Sprintf("## %s\n", name)
        title += "| No.  | Title                                                       | Mark |\n"
        title += "|------|-------------------------------------------------------------|------|\n"
        err = ioutil.WriteFile(topicPath, []byte(title), 0644)
        if err != nil {
            panic(err)
        }
        topicFile, err := os.OpenFile(topicPath, os.O_WRONLY|os.O_APPEND, 0644)
        if err != nil {
            panic(err)
        }
        defer topicFile.Close()
        // 遍历全部题目，找到包含当前标签的题目，逐行写入文件
        for _, question := range questions {
            topics := question.Get("topicTags").Array()
            for _, topic := range topics {
                if topic.Get("name").String() == name {
                    no := int(question.Get("questionId").Int())
                    if !questionsDone[no] {
                        continue
                    }
                    title := question.Get("title").String()
                    row := fmt.Sprintf("| %d | %s | |\n", no, title)
                    // TODO 格式化一下当前行，补若干个空格
                    _, err = topicFile.WriteString(row)
                    if err != nil {
                        panic(err)
                    }
                }
            }
        }
    },
}
