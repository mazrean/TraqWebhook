package webhook

import (
	"os"
	"crypto/hmac"
	"crypto/sha1"
	"encoding/hex"
	"github.com/labstack/echo"
	"log"
	"net/http"
	"strings"
)

var (
	traqWebhookID     = os.Getenv("TRAQ_WEBHOOK_ID")
	traqWebhookSecret = os.Getenv("TRAQ_WEBHOOK_SECRET")
	traqChannelID = os.Getenv("X_TRAQ_Channel_Id")
)

// postMessage Webhookにメッセージを投稿します
func postMessage(message string) error {
	url := "https://q.trap.jp/api/1.0/webhooks/" + traqWebhookID
	req, err := http.NewRequest("POST",
		url,
		strings.NewReader(message))
	if err != nil {
		log.Printf("Error occured while creating a new request: %s\n", err)
		return err
	}

	req.Header.Set(echo.HeaderContentType, echo.MIMETextPlainCharsetUTF8)
	req.Header.Set("X-TRAQ-Signature", generateSignature(message))
	if traqChannelID!=""{
		req.Header.Set("X-TRAQ-Channel-Id", traqChannelID)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	response := make([]byte, 512)
	_, err = resp.Body.Read(response)
	if err != nil {
		log.Printf("Error occured while reading response from traq webhook: %s\n", err)
	}

	log.Printf("Message sent to %s,\nmessage: %s, response: %s\n", url, message, response)

	return nil
}

func generateSignature(message string) string {
	mac := hmac.New(sha1.New, []byte(traqWebhookSecret))
	_, _ = mac.Write([]byte(message))
	return hex.EncodeToString(mac.Sum(nil))
}