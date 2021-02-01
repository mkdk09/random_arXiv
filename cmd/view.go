/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/spf13/cobra"
)

type Feed struct {
	XMLName xml.Name `xml:"feed"`
	Text    string   `xml:",chardata"`
	Xmlns   string   `xml:"xmlns,attr"`
	Link    struct {
		Text string `xml:",chardata"`
		Href string `xml:"href,attr"`
		Rel  string `xml:"rel,attr"`
		Type string `xml:"type,attr"`
	} `xml:"link"`
	Title struct {
		Text string `xml:",chardata"`
		Type string `xml:"type,attr"`
	} `xml:"title"`
	ID           string `xml:"id"`
	Updated      string `xml:"updated"`
	TotalResults struct {
		Text       string `xml:",chardata"`
		Opensearch string `xml:"opensearch,attr"`
	} `xml:"totalResults"`
	StartIndex struct {
		Text       string `xml:",chardata"`
		Opensearch string `xml:"opensearch,attr"`
	} `xml:"startIndex"`
	ItemsPerPage struct {
		Text       string `xml:",chardata"`
		Opensearch string `xml:"opensearch,attr"`
	} `xml:"itemsPerPage"`
	Entry struct {
		Text      string `xml:",chardata"`
		ID        string `xml:"id"`
		Updated   string `xml:"updated"`
		Published string `xml:"published"`
		Title     string `xml:"title"`
		Summary   string `xml:"summary"`
		Author    struct {
			Text string `xml:",chardata"`
			Name string `xml:"name"`
		} `xml:"author"`
		Comment struct {
			Text  string `xml:",chardata"`
			Arxiv string `xml:"arxiv,attr"`
		} `xml:"comment"`
		JournalRef struct {
			Text  string `xml:",chardata"`
			Arxiv string `xml:"arxiv,attr"`
		} `xml:"journal_ref"`
		Link []struct {
			Text  string `xml:",chardata"`
			Href  string `xml:"href,attr"`
			Rel   string `xml:"rel,attr"`
			Type  string `xml:"type,attr"`
			Title string `xml:"title,attr"`
		} `xml:"link"`
		PrimaryCategory struct {
			Text   string `xml:",chardata"`
			Arxiv  string `xml:"arxiv,attr"`
			Term   string `xml:"term,attr"`
			Scheme string `xml:"scheme,attr"`
		} `xml:"primary_category"`
		Category []struct {
			Text   string `xml:",chardata"`
			Term   string `xml:"term,attr"`
			Scheme string `xml:"scheme,attr"`
		} `xml:"category"`
	} `xml:"entry"`
}

// viewCmd represents the view command
var viewCmd = &cobra.Command{
	Use:   "view",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: main,
}

func main(cmd *cobra.Command, args []string) {
	// baseUrl := "http://export.arxiv.org/api/query?sortBy=lastUpdatedDate&sortOrder=descending&max_results=1&start=29500&search_query=cat:cs.AI"
	// baseUrl := "http://export.arxiv.org/api/query?sortBy=lastUpdatedDate&sortOrder=descending&max_results=1"
	// addCategoryUrl := baseUrl + "&search_query=" + categoryGet()
	// addRandomUrl := addCategoryUrl + "&start=" + randGet()
	url := makeUrl()
	data := httpGet(url)
	fmt.Println(data)
	result := Feed{}
	err := xml.Unmarshal([]byte(data), &result)
	if result.Entry.Title == "" {
		fmt.Println("タイトルが空の時(リクエストがエラーの時)エラーがなくなるまでリクエストを繰り返す")
	}
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}

	// タイトル
	fmt.Println(result.Entry.Title)
	// 発行日
	fmt.Println(result.Entry.Published)
	// 更新日
	fmt.Println(result.Entry.Updated)
	// サマリー
	fmt.Println(result.Entry.Summary)
	// 著者
	fmt.Println(result.Entry.Author.Name)
	// リンク
	// fmt.Println(result.Entry.Link[0].Href)
	// pdfリンク
	// fmt.Println(result.Entry.Link[1].Href)
	for _, link := range result.Entry.Link {
		fmt.Println(link.Href)
	}
	// 主要カテゴリ
	fmt.Println(result.Entry.PrimaryCategory.Term)
	// サブカテゴリー
	for _, category := range result.Entry.Category {
		fmt.Println(category.Term)
	}
}

func makeUrl() string {
	// baseUrl := "http://export.arxiv.org/api/query?sortBy=lastUpdatedDate&sortOrder=descending&max_results=1&start=29500&search_query=cat:cs.AI"
	baseUrl := "http://export.arxiv.org/api/query?sortBy=lastUpdatedDate&sortOrder=descending&max_results=1"
	addCategoryUrl := baseUrl + "&search_query=" + categoryGet()
	addRandomUrl := addCategoryUrl + "&start=" + randGet()
	return addRandomUrl
}

func httpGet(url string) string {
	response, _ := http.Get(url)
	body, _ := ioutil.ReadAll(response.Body)
	defer response.Body.Close()
	return string(body)
}

func categoryGet() string {
	category := []string{"cs.AI", "cs.LG", "cs.CL", "cs.CV", "cs.NE", "stat.ML"}
	rand.Seed(time.Now().UnixNano())
	if globalFlags.category == "random" {
		return category[rand.Intn(len(category))]
	} else {
		return globalFlags.category
	}
}

func randGet() string {
	rand.Seed(time.Now().UnixNano())
	return strconv.Itoa(rand.Intn(30000))
}

func init() {
	rootCmd.AddCommand(viewCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// viewCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// viewCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
