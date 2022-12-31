package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type CompletionReply struct {
	ID      string `json:"id"`
	Object  string `json:"object"`
	Created int    `json:"created"`
	Model   string `json:"model"`
	Choices []struct {
		Text         string      `json:"text"`
		Index        int         `json:"index"`
		Logprobs     interface{} `json:"logprobs"`
		FinishReason string      `json:"finish_reason"`
	} `json:"choices"`
	Usage struct {
		PromptTokens     int `json:"prompt_tokens"`
		CompletionTokens int `json:"completion_tokens"`
		TotalTokens      int `json:"total_tokens"`
	} `json:"usage"`
}

type CompletionRequest struct {
	Model            string  `json:"model"`
	Prompt           string  `json:"prompt"`
	Temperature      float64 `json:"temperature"`
	MaxTokens        int     `json:"max_tokens"`
	TopP             float64 `json:"top_p"`
	FrequencyPenalty float64 `json:"frequency_penalty"`
	PresencePenalty  float64 `json:"presence_penalty"`
	Stop             string  `json:"stop"`
}

func completion(prompt string) string {
	if creds.OpenAI != "" && prompt != "" {
		// Set up the API endpoint URL and request body
		endpoint := "https://api.openai.com/v1/completions"
		requestBody := CompletionRequest{
			Model:            "text-davinci-003",
			Prompt:           prompt,
			Temperature:      0.7,
			MaxTokens:        50,
			TopP:             1,
			FrequencyPenalty: 0.0,
			PresencePenalty:  0.0,
			Stop:             "",
		}

		jsonRequest, err := json.Marshal(requestBody)
		if err != nil {
			fmt.Println(err)
		}

		// Create an HTTP client and make the API request
		client := &http.Client{}
		req, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(jsonRequest))
		if err != nil {
			fmt.Println(err)
		}
		req.Header.Add("Content-Type", "application/json")
		req.Header.Add("Authorization", "Bearer "+creds.OpenAI)
		resp, err := client.Do(req)
		if err != nil {
			fmt.Println(err)
		}
		defer resp.Body.Close()

		// Read the response
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Println(err)
		}

		var completionreply CompletionReply
		err = json.Unmarshal([]byte(body), &completionreply)
		if err != nil {
			fmt.Println(err)
		}

		// Print the 1st response
		return strings.TrimSpace(completionreply.Choices[0].Text)
	}
	return "Please setup your OpenAI API key @ https://openai.com"
}
