package utils

import (
	"fmt"
	"net/http"
	"time"
	"log"
	"github.com/go-resty/resty"
	"bufio"
	"os"
	"golang.org/x/crypto/ssh/terminal"
	"crypto/tls"
	"strings"
)

func InvokePOSTRequest(url string, headers map[string]string, body string) (*resty.Response, error) {
	resp, err := resty.R().SetHeaders(headers).SetBody(body).Post(url)

	return resp, err
}

func PromptForUsername() string {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter Username: ")
	username, _ := reader.ReadString('\n')

	return username
}

func PromptForPassword() string {
	fmt.Print("Enter Password: ")
	bytePassword, _ := terminal.ReadPassword(0)
	password := string(bytePassword)

	defer fmt.Println()
	return password
}

func ExportAPI(name string, version string, url string, accessToken string) *resty.Response {
	// append '/' to the end if there isn't one already
	if string(url[len(url)-1]) != "/" {
		url += "/"
	}
	url += "export/apis"

	query := "?query=" + name
	url += query
	fmt.Println("ExportAPI: URL:", url)
	headers := make(map[string]string)
	headers[HeaderAuthorization] = HeaderValueAuthBearerPrefix + " " + accessToken
	headers[HeaderAccept] = HeaderValueApplicationZip

	resp, err := resty.R().
		SetHeaders(headers).
		Get(url)

	if err != nil {
		fmt.Println("Error exporting API:", name)
		panic(err)
	}

	return resp
}

func ImportAPI(name string, version string, url string, accessToken string) *resty.Response {
	// append '/' to the end if there isn't one already
	if string(url[len(url)-1]) != "/" {
		url += "/"
	}
	url += "imports/api"

	file := "./exported/" + name + ".zip"
	fmt.Println("File:", file)
	fmt.Println("ImportAPI: URL:", url)
	body := strings.NewReader("------WebKitFormBoundary7MA4YWxkTrZu0gW\r" +
		"Content-Disposition: form-data; name=\"file\"; filename=\"harrypotter.zip\"\r\nContent-Type: application/zip\r" +
		"\r" +
		"\r" +
		"------WebKitFormBoundary7MA4YWxkTrZu0gW--")

	headers := make(map[string]string)

	headers[HeaderAuthorization] = HeaderValueAuthBearerPrefix + " " + accessToken
	// headers["Authorization"] = "Bearer " + accessToken

	headers[HeaderAccept] = "application/json"

	headers[HeaderContentType] = HeaderValueMultiPartFormData
	// headers["Content-Type"] = "multipart/form-data"

	headers[HeaderConsumes] = HeaderValueMultiPartFormData
	// headers["Consumes"] = "multipart/form-data"

	resty.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true}) // To bypass errors in HTTPS certificates
	resp, err := resty.R().
		SetHeaders(headers).
		SetBody(body).
		Put(url)

	if err != nil {
		fmt.Println("Error importing API:", name)
		panic(err)
	}

	return resp
}

func Authenticate() {

}

// Send the HTTP request to the server. Default timeout is one minute
func SendRequest(req *http.Request, timeout time.Duration) *http.Response {
	client := &http.Client{
		Timeout: timeout,
	}
	resp, err := client.Do(req)
	if err != nil {
		// TODO: Handle accordingly
	}
	return resp
}

// invoke the client request and handle error scenarios
func invokeRequest(req *http.Request, timeout time.Duration) *http.Response {
	httpResp := SendRequest(req, timeout)
	log.Printf("Status Code: %v\n", httpResp.Status)

	// TODO:: handle error responses

	return httpResp
}
