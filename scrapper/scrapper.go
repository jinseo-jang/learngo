package scrapper

import (
	"encoding/csv"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type extractedJob struct {
	id       string
	title    string
	location string
	salary   string
	summary  string
}

//Scrape job infomation in indeed.com by a term
func Scrape(term string) {
	// jobs :=   make([]extractedJob, 1)
	var baseURL string = "https://kr.indeed.com/jobs?q=" + term + "&limit=50"
	var jobs []extractedJob
	c := make(chan []extractedJob)
	totalPages := getPages(baseURL)

	for i := 0; i < totalPages; i++ {
		// extractedJobs := getPage(i)
		go getPage(i, baseURL, c)
		// jobs = append(jobs, extractedJobs...)
	}
	//fmt.Println(jobs)

	for i := 0; i < totalPages; i++ {
		extractedJob := <-c
		jobs = append(jobs, extractedJob...)
	}
	writeJobs(jobs)
	fmt.Println("Done, extracted", len(jobs))

}

func getPage(page int, url string, mainC chan<- []extractedJob) {
	var jobs []extractedJob
	c := make(chan extractedJob)

	pageURL := url + "&start=" + strconv.Itoa(page*50)
	fmt.Println("Requesting URL: ", pageURL)
	// http.Get(pageURL)
	res, err := http.Get(pageURL)
	checkErr(err)
	checkCode(res)

	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)

	searchCards := doc.Find(".jobsearch-SerpJobCard")
	searchCards.Each(func(i int, card *goquery.Selection) {
		go extractJob(card, c)
		// jobs = append(jobs, job)
		// id, _ := card.Attr("data-jk")
		// title := cleanString(card.Find(".title>a").Text())
		// location := cleanString(card.Find(".sjcl").Text())
		// fmt.Println(id, title, location)
	})
	for i := 0; i < searchCards.Length(); i++ {
		job := <-c
		jobs = append(jobs, job)
	}
	// return jobs
	mainC <- jobs

}

func extractJob(card *goquery.Selection, c chan<- extractedJob) {
	id, _ := card.Attr("data-jk")
	title := CleanString(card.Find(".title>a").Text())
	location := CleanString(card.Find(".sjcl").Text())
	salary := CleanString(card.Find(".salaryText").Text())
	summary := CleanString(card.Find(".summary").Text())
	c <- extractedJob{
		id:       id,
		title:    title,
		location: location,
		salary:   salary,
		summary:  summary}
	// fmt.Println(id, title, location, salary, summary)
}

func getPages(url string) int {
	pages := 0
	res, err := http.Get(url)
	checkErr(err)
	checkCode(res)

	defer res.Body.Close()
	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)
	// fmt.Println(doc)
	doc.Find(".pagination").Each(func(i int, s *goquery.Selection) {
		pages = s.Find("a").Length()
	})

	return pages
}

func writeJobs(jobs []extractedJob) {
	// jobAllData := []string{}
	var jobAllData [][]string
	file, err := os.Create("jobs.csv")
	checkErr(err)

	w := csv.NewWriter(file)
	defer w.Flush()

	headers := []string{"ID", "TITLE", "LOCATION", "SALARY", "SUMMARY"}
	wErr := w.Write(headers)
	checkErr(wErr)

	c := make(chan []string)
	for _, job := range jobs {
		// jobSlice := []string{"https://kr.indeed.com/viewjob?jk=" + job.id, job.title, job.location, job.salary, job.summary}
		// jwErr := w.Write(jobSlice)
		// checkErr(jwErr)
		go getJobData(job, c)
	}

	for i := 0; i < len(jobs); i++ {
		jobData := <-c
		fmt.Println("jobData: ", jobData)
		jobAllData = append(jobAllData, jobData)
	}
	fmt.Println("length:", len(jobAllData))
	fmt.Println(jobAllData[0])
	jwErr := w.WriteAll(jobAllData)
	checkErr(jwErr)
}

func getJobData(job extractedJob, c chan<- []string) {
	jobLink := "https://kr.indeed.com/viewjob?jk=" + job.id
	jobInfo := []string{jobLink, job.title, job.location, job.salary, job.summary}
	c <- jobInfo
}

func checkErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func checkCode(res *http.Response) {
	if res.StatusCode != 200 {
		log.Fatalln("Request failed with status: ", res.StatusCode)
	}
}

//CleanString to trim string input
func CleanString(str string) string {
	return strings.Join(strings.Fields(strings.TrimSpace(str)), " ")
}
