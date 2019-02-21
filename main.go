package main

import (
    "os"
    "os/exec"
    "fmt"
    "log"
    "time"
    "strings"
    "go/build"
    "net/http"
    "io/ioutil"
    "encoding/json"
    "github.com/atotto/clipboard"
    "github.com/thoongnv/tzutils/config"
)

// Project object mapped with Gitlab project
type Project struct {
    Name            string      `json:"name"`
    WebURL          string      `json:"web_url"`
}

func main() {
    // read config from file
    gopath := os.Getenv("GOPATH")
    if gopath == "" {
        gopath = build.Default.GOPATH
    }

    v, err := config.ReadConfig("tzutils", gopath + "/src/github.com/thoongnv/tzutils")
    if err != nil {
        fmt.Println(err)
    }

    // save config params
    gitlabAPIURL := v.GetString("gitlab.api_url")
    gitlabPrivateToken := v.GetString("gitlab.private_token")

    // build request url
    reqURL := gitlabAPIURL + "/projects?private_token=" + gitlabPrivateToken
    client := &http.Client{Timeout: 10 * time.Second}

    // get all projects
    resp, err := client.Get(reqURL)
    if err != nil {
        log.Fatal(err)
    }
    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)

    // mapping gitlab projects
    var projects []Project
    err = json.Unmarshal(body, &projects)
    if err != nil {
        log.Fatal(err)
    }

    // get last folder name of current directory
    dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
    }
    folders := strings.Split(dir, "/")
    folder := folders[len(folders)-1]

    var found = false
    for _, p := range projects {
		if p.Name == folder {
            found = true
            // get last commit hash
            cmd := exec.Command("git", "rev-parse", "HEAD")
            commit, err := cmd.Output()
            if err != nil {
                log.Fatal(err)
            }
            commitURL := "[COMMIT](" + p.WebURL + "/commit/" + strings.TrimSuffix(string(commit), "\n") + ")"

            // copy to clipboard
            clipboard.WriteAll(commitURL)
            fmt.Println("Let's pasteeeeeeeeee ...")
		}
    }

    if found == false {
        fmt.Println("Failed to generate commit URL ...")
    }
}
