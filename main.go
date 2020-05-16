package main

import(
    "fmt"
    "time"
    "./subcodes"
    "os/exec"
    "os"
    "strconv"
    "log"
)

// goルーチンで実行する関数
func asyncFunc(s string) {
    fmt.Println(s)
}

func stdout(temp float64, ilm float64, hum float64) {
    const layout = "2006-01-02 15:04:05"
    t := time.Now()
    fmt.Printf("%s,%.2f,%.2f,%.2f\n", t.Format(layout), temp, ilm, hum)

}

func jsonparse_n_savetext() {
    var rs [3]float64
    rs = subcodes.Jsonparse()
    stdout(rs[0], rs[1], rs[2])
    //fmt.Printf("%.2f,%.2f,%.2f\n", rs[0], rs[1], rs[2])

    //save text as tmux style
    out, err := exec.Command("ColoredTemp", strconv.FormatFloat(rs[0], 'f', 2, 64), strconv.FormatFloat(rs[2], 'f', 2, 64)).Output()
    if err != nil {
        log.Fatal(err)
    }

    result_file, err := os.Create("./Data/result.txt")
    if err != nil{
        fmt.Println("Failed to create file")
    }
    result_file.WriteString(string(out))
    defer result_file.Close()
}

func missioncontroller() {
    subcodes.GetDataNSave()
    jsonparse_n_savetext()
}

// メイン関数
func main() {
    // goルーチンの関数の実行
    /*
    for i := 0;; i++{
        str := fmt.Sprintf("Go routine (no: %v)", i)
        go asyncFunc(str)

        time.Sleep(1 * time.Second)
    }
    */
    for {
        go missioncontroller()
        time.Sleep(10 * time.Minute)
    }

}
