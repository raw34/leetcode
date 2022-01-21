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
        // 全部题目
        questions := getAllQuestions()
        // 已完成的题目
        doneQuestions := getDoneQuestions()
        // 当前标签已收录的题目
        topicQuestions := getTopicQuestions(name)
        // 题目写入标签文件
        question2Topic(name, questions, doneQuestions, topicQuestions)
    },
}

func getAllQuestions() []gjson.Result {
    // 读取全部题目文件
    questionsPath := "questions.json"
    content, err := ioutil.ReadFile(questionsPath)
    if err != nil {
        panic(err)
    }
    // 解析全部题目文件
    questions := gjson.Parse(string(content)).Array()
    return questions
}

func getDoneQuestions() map[int]bool {
    doneQuestions := map[int]bool{}
    // 读取README文件
    readmePath := "README.md"
    content, err := ioutil.ReadFile(readmePath)
    if err != nil {
        panic(err)
    }
    // 解析README文件
    markdown := goldmark.New()
    readme := (markdown.Parser()).Parse(text.NewReader(content))
    // 遍历README文件，获取已完成题目编号
    readme = readme.FirstChild()
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
                    doneQuestions[i] = true
                }
            }
            break
        }
    }
    return doneQuestions
}

func getTopicQuestions(name string) map[int]string {
    topicQuestions := map[int]string{}
    // 读取topic文件
    topicPath := fmt.Sprintf("topic/%s.md", strings.Replace(name, " ", "", -1))
    content, err := ioutil.ReadFile(topicPath)
    if err != nil {
        panic(err)
    }
    // 解析topic文件
    markdown := goldmark.New()
    topic := (markdown.Parser()).Parse(text.NewReader(content))
    // 遍历topic文件，获取全部题目
    topic = topic.FirstChild()
    for topic.NextSibling() != nil {
        topic = topic.NextSibling()
        kind := topic.Kind().String()
        if kind == "Paragraph" {
            child := topic.FirstChild().NextSibling()
            for child.NextSibling() != nil {
                child = child.NextSibling()
                row := string(child.Text(content))
                arr := strings.Split(row, "|")
                no, _ := strconv.Atoi(strings.Trim(arr[1], " "))
                topicQuestions[no] = row
            }
            break
        }
    }
    return topicQuestions
}

func question2Topic(name string, questions []gjson.Result, doneQuestions map[int]bool, topicQuestions map[int]string) {
    topicPath := "topic/" + strings.Replace(name, " ", "", -1) + ".md"
    header := fmt.Sprintf("## %s\n", name)
    header += "| No.  | Title                                                       | Mark |\n"
    header += "|------|-------------------------------------------------------------|------|\n"
    err := ioutil.WriteFile(topicPath, []byte(header), 0644)
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
                if !doneQuestions[no] {
                    continue
                }
                row, ok := topicQuestions[no]
                if !ok {
                    title := question.Get("title").String()
                    row = fmt.Sprintf("| %d | %s | |", no, title)
                }
                row += "\n"
                // TODO 格式化一下当前行，补若干个空格
                _, err = topicFile.WriteString(row)
                if err != nil {
                    panic(err)
                }
            }
        }
    }
}

func getAllTopics() []string {
    return []string{
        "Array",
        "String",
        "Hash Table",
        "Dynamic Programming",
        "Math",
        "Depth-First Search",
        "Sorting",
        "Breadth-First Search",
        "Greedy",
        "Tree",
        "Binary Tree",
        "Binary Search",
        "Database",
        "Matrix",
        "Two Pointers",
        "Bit Manipulation",
        "Stack",
        "Design",
        "Heap (Priority Queue)",
        "Backtracking",
        "Graph",
        "Linked List",
        "Simulation",
        "Prefix Sum",
        "Sliding Window",
        "Counting",
        "Union Find",
        "Recursion",
        "Binary Search Tree",
        "Divide and Conquer",
        "Trie",
        "Monotonic Stack",
        "Ordered Set",
        "Queue",
        "Memoization",
        "Geometry",
        "Bitmask",
        "Enumeration",
        "Topological Sort",
        "Segment Tree",
        "Game Theory",
        "Hash Function",
        "Data Stream",
        "Binary Indexed Tree",
        "String Matching",
        "Interactive",
        "Shortest Path",
        "Combinatorics",
        "Rolling Hash",
        "Randomized",
        "Merge Sort",
        "Monotonic Queue",
        "Doubly-Linked List",
        "Quickselect",
        "Number Theory",
        "Iterator",
        "Brainteaser",
        "Probability and Statistics",
        "Concurrency",
        "Bucket Sort",
        "Counting Sort",
        "Suffix Array",
        "Minimum Spanning Tree",
        "Line Sweep",
        "Shell",
        "Reservoir Sampling",
        "Eulerian Circuit",
        "Strongly Connected Component",
        "Rejection Sampling",
        "Radix Sort",
        "Biconnected Component",
    }
}
