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
	Max_Tokens       int     `json:"max_tokens"`
	Top_P            float64 `json:"top_p"`
	FrequencyPenalty float64 `json:"frequency_penalty"`
	PresencePenalty  float64 `json:"presence_penalty"`
	Stop             string  `json:"stop"`
}

func completion(prompt, profile string) string {
	var completionrequest CompletionRequest
	var requestBody CompletionRequest

	completionrequest = settings.API.OpenAI

	if creds.OpenAI != "" && prompt != "" {
		// Set up the API endpoint URL and request body
		endpoint := "https://api.openai.com/v1/completions"

		completionrequest.Prompt = prompt

		if profile == "fact" {
			completionrequest.Temperature = 0.0
			requestBody = completionrequest
		} else {
			completionrequest.Temperature = 0.7
			requestBody = completionrequest
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
		fmt.Println("gpt reply: ", completionreply)
		// Print the 1st response
		if len(completionreply.Choices) > 0 {
			return strings.TrimSpace(completionreply.Choices[0].Text)
		} else {
			return "Please setup your OpenAI API key @ https://openai.com"
		}
	}
	return "Please setup your OpenAI API key @ https://openai.com"
}
