package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"sync/atomic"

	"github.com/joho/godotenv"
)

func main() {
	type token struct{}

	limit := 100
	flag.Parse()
	args := flag.Args()

	if len(args) < 1 {
		log.Fatalln("Enter the folder path")
	}

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error cargando el archivo .env:", err)
	}

	HTTP_API := os.Getenv("HTTP_API")
	ORG_ID := os.Getenv("ORG_ID")
	STREAM := os.Getenv("STREAM")
	USER := os.Getenv("ZO_ROOT_USER_EMAIL")
	PASSWORD := os.Getenv("ZO_ROOT_USER_PASSWORD")

	client := NewOpenObserverClient(
		HTTP_API,
		ORG_ID,
		STREAM,
		USER,
		PASSWORD,
	)

	folderPath := args[0]
	out := make(chan string)
	go getPathFile(folderPath, out)

	var wg sync.WaitGroup
	sem := make(chan token, limit)
	fileReadChannels := make(chan map[string]string)

	go func() {
		for pathFile := range out {
			sem <- token{}
			wg.Add(1)

			go func(p string) {
				defer wg.Done()
				fileReadChannels <- readFile(p)
				<-sem
			}(pathFile)
		}

		wg.Wait()
		close(fileReadChannels)
	}()

	maxChunk := 2000
	chunk := make([]map[string]string, 0, maxChunk)
	var sendWg sync.WaitGroup
	semOpen := make(chan token, 10)
	for fileRead := range fileReadChannels {
		chunk = append(chunk, fileRead)
		if len(chunk) >= maxChunk {
			sendWg.Add(1)
			data := append([]map[string]string{}, chunk...)
			go func(data []map[string]string) {
				defer sendWg.Done()
				semOpen <- token{}
				client.send(data)
				<-semOpen
			}(data)

			chunk = chunk[:0]
		}
	}

	if len(chunk) > 0 {
		sendWg.Add(1)
		data := append([]map[string]string{}, chunk...)
		go func(data []map[string]string) {
			defer sendWg.Done()
			client.send(data)
		}(data)
	}

	sendWg.Wait()

}

func getPathFile(pathDir string, out chan string) {
	err := filepath.WalkDir(pathDir, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !d.IsDir() {
			out <- path
		}
		return nil
	})

	if err != nil {
		log.Fatal(err)
	}
	close(out)
}

func readFile(path string) map[string]string {

	var commonHeader map[string]string
	var commonHeaderOnce sync.Once

	initCommonHeader := func() {

		commonHeader = make(map[string]string)
		for _, v := range []string{
			"Message-ID",
			"Date",
			"From",
			"To",
			"Subject",
			"Mime-Version",
			"Content-Type",
			"Content-Transfer-Encoding",
			"X-From",
			"X-To",
			"X-cc",
			"X-bcc",
			"Cc",
			"X-Folder",
			"X-Origin",
			"X-FileName",
		} {
			commonHeader[v] = v
		}
	}

	valideKey := func(key string) bool {
		commonHeaderOnce.Do(initCommonHeader)
		if v := commonHeader[key]; v != "" {
			return true
		}
		return false
	}

	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var bodyLines []string
	var currentKey string

	headers := make(map[string]string)
	scanner := bufio.NewScanner(file)

	isHeader := true

	for scanner.Scan() {

		line := scanner.Text()

		if isHeader {

			if len(line) == 0 {
				isHeader = false
				continue
			}

			key, value, ok := strings.Cut(line, ":")

			if ok {

				if !valideKey(key) {
					headers[currentKey] += strings.TrimSpace(line)
					continue
				}

				currentKey = strings.ToLower(key)
				headers[currentKey] += strings.TrimSpace(value)
				continue
			}

			headers[currentKey] += " " + strings.TrimSpace(key)
		}

		if !isHeader {
			bodyLines = append(bodyLines, line)
		}
	}

	headers["body"] = strings.Join(bodyLines, "\n")

	return headers
}

var total int32

type OpenObserverClient struct {
	client     *http.Client
	orgID      string
	streamName string
	user       string
	password   string
	Http_api   string
	baseURL    string
}

func NewOpenObserverClient(http_api, orgID, streamName, user, password string) *OpenObserverClient {
	url := fmt.Sprintf("%s/%s/%s/_json", http_api, orgID, streamName)
	return &OpenObserverClient{
		client:     &http.Client{},
		orgID:      orgID,
		streamName: streamName,
		user:       user,
		password:   password,
		baseURL:    url,
	}
}

func (ooc *OpenObserverClient) send(chunk []map[string]string) (string, error) {

	jsonData, err := json.Marshal(chunk)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("POST", ooc.baseURL, bytes.NewBuffer(jsonData))
	if err != nil {
		return "", err
	}

	req.SetBasicAuth(ooc.user, ooc.password)
	req.Header.Set("Content-Type", "application/json")

	resp, err := ooc.client.Do(req)

	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	fmt.Println(string(body))
	fmt.Println("count:", atomic.AddInt32(&total, 1))
	return string(body), nil
}
