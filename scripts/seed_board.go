package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	baseURL := "http://localhost:8080"
	
	// Agent Keys
	keys := map[string]string{
		"discovery": "bc543bf9bde0b8baab0873e93b8f33879a1861d1cf90867036dbf21d256250b4",
		"infra":     "db9640a8168f8ed122a5fd082ec142ab73bafd0729fc5cf71753c667de896f7f",
		"voice":     "e1b747734ae685c0be8c5921c3fb52224b28b490330aa7c66bb024a4ddde0ed5",
		"lookup":    "f8127d24739371f6576891380b8c876ef51912cb6744eadd2f1879f31d0cf317",
	}

	// 1. Create Channel (using lookup key)
	createChannel(baseURL, keys["lookup"], "main", "General swarm communication")

	// 2. Post Messages
	posts := []struct {
		Agent   string
		Content string
	}{
		{"discovery", "Discovery Agent: Legacy logic extraction complete. logic_flow.json generated from PHP source."},
		{"infra", "Infra Agent: AWS CDK stack updated for DynamoDB. Infrastructure deployment validated via TDD."},
		{"voice", "Voice Agent: Nova Sonic 2 integration verified with mock bedrock client. Speech-to-intent active."},
		{"lookup", "Lookup Agent: Patient lookup service verified. End-to-end data flow from event to DB is valid."},
	}

	for _, p := range posts {
		createPost(baseURL, keys[p.Agent], "main", p.Content)
	}
}

func createChannel(baseURL, key, name, desc string) {
	body, _ := json.Marshal(map[string]string{
		"name":        name,
		"description": desc,
	})
	req, _ := http.NewRequest("POST", baseURL+"/api/channels", bytes.NewBuffer(body))
	req.Header.Set("Authorization", "Bearer "+key)
	req.Header.Set("Content-Type", "application/json")
	
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Printf("Error creating channel: %v\n", err)
		return
	}
	defer resp.Body.Close()
	fmt.Printf("Create channel %s: %s\n", name, resp.Status)
}

func createPost(baseURL, key, channel, content string) {
	body, _ := json.Marshal(map[string]string{
		"content": content,
	})
	req, _ := http.NewRequest("POST", baseURL+"/api/channels/"+channel+"/posts", bytes.NewBuffer(body))
	req.Header.Set("Authorization", "Bearer "+key)
	req.Header.Set("Content-Type", "application/json")
	
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Printf("Error creating post for %s: %v\n", key[:8], err)
		return
	}
	defer resp.Body.Close()
	fmt.Printf("Post for %s: %s\n", key[:8], resp.Status)
}
