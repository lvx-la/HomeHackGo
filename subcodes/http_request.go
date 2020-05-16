package subcodes

import (
    "net/http"
    "fmt"
    "io/ioutil"
    "os"
    "log"
)

func SaveJson(jsonbody []uint8){
    file, err := os.Create("./data/data.json")
    if err != nil{
        fmt.Println("Failed to create file")
    }
    defer file.Close()
    //encodedjson, _ := json.Marshal(&body)
    file.Write(([]byte)(string(jsonbody)))
    //fmt.Println("Done")
}

func GetDataNSave() {
    /*
        URLにアクセスしてJsonを特定のディレクトリに保存する。
        issues
        セキュリティ的にトークンを分離した方がいい
    */

    client := new(http.Client)
    var url string
    url = "https://api.nature.global/1/devices"

    req, err := http.NewRequest("GET", url, nil)
    if err != nil {
        fmt.Println("Error")
        log.Fatal(err)
    }
    naturekey := os.Getenv("NATUREAPI")

    req.Header.Set("Accept", "application/json")
    req.Header.Set("Authorization", naturekey)

    resp, err := client.Do(req)
    if err != nil {
        fmt.Println("Dead")
    }
    defer resp.Body.Close()

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        fmt.Println("Dead")
    }
    /*
    v := reflect.ValueOf(body)
    fmt.Printf("body's Type is %s", v.Type())
    fmt.Println(v.Kind())
    */


    //fmt.Println(string(body))
    SaveJson(body)
}

/*
func main () {
    GetDataNSave()
}
*/
