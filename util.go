package sparkapi

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// AssembleAuthURL creates the authenticated URL for the WebSocket connection.
func AssembleAuthURL(httpMethod string, hostURL, apiKey, apiSecret string) string {
	ul, err := url.Parse(hostURL)
	if err != nil {
		log.Fatalf("URL parse error: %v", err)
	}

	date := time.Now().UTC().Format(time.RFC1123)
	signString := []string{"host: " + ul.Host, "date: " + date, httpMethod + " " + ul.Path + " HTTP/1.1"}
	signature := HmacWithShaToBase64("hmac-sha256", strings.Join(signString, "\n"), apiSecret)
	authorization := base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("hmac username=\"%s\", algorithm=\"%s\", headers=\"%s\", signature=\"%s\"", apiKey, "hmac-sha256", "host date request-line", signature)))

	v := url.Values{}
	v.Add("host", ul.Host)
	v.Add("date", date)
	v.Add("authorization", authorization)
	return hostURL + "?" + v.Encode()
}

// HmacWithShaToBase64 computes HMAC with SHA-256 and returns the result in base64 encoding.
func HmacWithShaToBase64(algorithm, data, key string) string {
	mac := hmac.New(sha256.New, []byte(key))
	mac.Write([]byte(data))
	encodedData := mac.Sum(nil)
	return base64.StdEncoding.EncodeToString(encodedData)
}

// readResp reads the HTTP response and returns a formatted string.
func readResp(resp *http.Response) string {
	if resp == nil {
		return ""
	}
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Failed to read response body: %v", err)
	}
	return fmt.Sprintf("code=%d,body=%s", resp.StatusCode, string(b))
}
