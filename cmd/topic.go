package cmd

import (
    "fmt"
    "github.com/spf13/cobra"
    "github.com/tidwall/gjson"
    "github.com/yuin/goldmark"
    "github.com/yuin/goldmark/text"
    "io/ioutil"
    "strconv"
    "strings"
)

func init() {
    var Name string

    topicCmd.Flags().StringVarP(&Name, "name", "", "All", "标签名称")

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
        fmt.Println("name is " + name)
        if name != "All" {
            saveTopic(name)
        } else {
            names := getAllTopics()
            for _, name := range names {
                saveTopic(name)
            }
        }
    },
}

func saveTopic(name string) {
    name = strings.Replace(name, " ", "", -1)
    // 全部题目
    questions := getAllQuestions()
    // 已完成的题目
    doneQuestions := getReadmeQuestions()
    // 当前标签已收录的题目
    topicQuestions := getTopicQuestions(name)
    // 题目写入标签文件
    question2Topic(name, questions, doneQuestions, topicQuestions)
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

func getReadmeQuestions() map[int]string {
    // 读取README文件
    readmePath := "README.md"
    questions := getDoneQuestions(readmePath, "readme")
    for no, question := range questions {
        row := strings.Split(question, "|")
        done := strings.Trim(row[1], " ")
        if done != "✅" {
            delete(questions, no)
        }
    }

    return questions
}

func getDoneQuestions(filePath, fileType string) map[int]string {
    noIndex := map[string]int{"topic": 1, "readme": 2}
    questions := map[int]string{}
    // 读取topic文件
    content, err := ioutil.ReadFile(filePath)
    if err != nil {
        return questions
    }
    // 解析topic文件
    markdown := goldmark.New()
    node := (markdown.Parser()).Parse(text.NewReader(content))
    // 遍历topic文件，获取全部题目
    node = node.FirstChild()
    for node.NextSibling() != nil {
        node = node.NextSibling()
        kind := node.Kind().String()
        if kind != "Paragraph" {
            continue
        }
        child := node.FirstChild().NextSibling()
        for child.NextSibling() != nil {
            child = child.NextSibling()
            row := string(child.Text(content))
            arr := strings.Split(row, "|")
            i := noIndex[fileType]
            no, _ := strconv.Atoi(strings.Trim(arr[i], " "))
            questions[no] = row
        }
    }
    return questions
}

func getTopicQuestions(name string) map[int]string {
    // 读取topic文件
    topicPath := fmt.Sprintf("topic/%s.md", name)
    return getDoneQuestions(topicPath, "topic")
}

func question2Topic(name string, questions []gjson.Result, doneQuestions map[int]string, topicQuestions map[int]string) {
    topicPath := fmt.Sprintf("topic/%s.md", name)
    content := fmt.Sprintf("## %s\n", name)
    content += "| No.  | Title                                                       | Mark |\n"
    content += "|------|-------------------------------------------------------------|------|\n"
    // 遍历全部题目，找到包含当前标签的题目，逐行写入文件
    for _, question := range questions {
        topics := question.Get("topicTags").Array()
        for _, topic := range topics {
            no := int(question.Get("questionId").Int())
            topicName := topic.Get("name").String()
            if _, ok := doneQuestions[no]; !ok || topicName != name {
                continue
            }
            row, ok := topicQuestions[no]
            if ok {
                row += "\n"
                delete(topicQuestions, no)
            } else {
                title := question.Get("title").String()
                row = fmt.Sprintf("| %d | %s | |\n", no, title)
            }
            // TODO 格式化一下当前行，补若干个空格
            content += row
        }
    }
    // 补全手动覆盖标签的数据
    for no, row := range topicQuestions {
        if no == 0 {
            continue
        }
        content += row + "\n"
    }
    // 写入文件
    err := ioutil.WriteFile(topicPath, []byte(content), 0644)
    if err != nil {
        panic(err)
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
        //"Database",
        "Matrix",
        "Two Pointers",
        "Bit Manipulation",
        "Stack",
        "Design",
        "Heap (Priority Queue)",
        "Backtracking",
        "Graph",
        "Linked List",
        //"Simulation",
        "Prefix Sum",
        "Sliding Window",
        //"Counting",
        "Union Find",
        "Recursion",
        "Binary Search Tree",
        "Divide and Conquer",
        "Trie",
        "Monotonic Stack",
        //"Ordered Set",
        "Queue",
        //"Memoization",
        //"Geometry",
        //"Bitmask",
        //"Enumeration",
        //"Topological Sort",
        //"Segment Tree",
        //"Game Theory",
        //"Hash Function",
        //"Data Stream",
        //"Binary Indexed Tree",
        //"String Matching",
        //"Interactive",
        //"Shortest Path",
        //"Combinatorics",
        //"Rolling Hash",
        //"Randomized",
        //"Merge Sort",
        "Monotonic Queue",
        "Doubly-Linked List",
        //"Quickselect",
        //"Number Theory",
        //"Iterator",
        //"Brainteaser",
        //"Probability and Statistics",
        //"Concurrency",
        //"Bucket Sort",
        //"Counting Sort",
        //"Suffix Array",
        //"Minimum Spanning Tree",
        //"Line Sweep",
        //"Shell",
        //"Reservoir Sampling",
        //"Eulerian Circuit",
        //"Strongly Connected Component",
        //"Rejection Sampling",
        //"Radix Sort",
        //"Biconnected Component",
    }
}
