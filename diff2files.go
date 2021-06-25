// Do "unset GOROOT"
package main

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "os"
)

type Diff struct {
    Id        int    `json:"id"`
    Type      string `json:"type"`
    Label     string `json:"label"`
    Title     string `json:"title"`
    Oid       int    `json:"oid"`
    Status    string `json:"status"`
    SourceDDL string `json:"source_ddl"`
    TargetDDL string `json:"target_ddl"`
    DiffDDL   string `json:"diff_ddl"`
    GroupName string `json:"group_name"`
}

func main() {
    jsonFile, err := os.Open("diff.json")
    if err != nil {
        fmt.Println(err)
    }

    defer jsonFile.Close()

    byteValue, _ := ioutil.ReadAll(jsonFile)

    var diffs []Diff

    json.Unmarshal(byteValue, &diffs)

    err = os.RemoveAll("database")
    if err != nil {
        panic(err)
    }

    err = os.Mkdir("database", 0755)
    if err != nil {
        panic(err)
    }

    for i := 0; i < len(diffs); i++ {
        f, err := os.Create("database/" + diffs[i].Label + ".sql")
        if err != nil {
            panic(err)
        }
        defer f.Close()
        if _, err = f.WriteString(diffs[i].DiffDDL); err != nil {
            panic(err)
        }
    }

  }
