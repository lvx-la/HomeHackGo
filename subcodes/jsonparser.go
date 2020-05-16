package subcodes

import (
    //"fmt"
    "io/ioutil"
    "log"
    "encoding/json"
)

type Humidity struct {
    Hu_value    float64 `json:"val"`
}

type Illumination struct {
    Il_value    float64 `json:"val"`
}

type Temperature struct {
    Te_value    float64 `json:"val"`
}

type NewestEvents struct {
    Hu  Humidity `json:"hu"`
    Il  Illumination `json:il`
    Te  Temperature `json:te`
}

type RoomState struct {
    Name    string `json:"id"`
    NE      NewestEvents `json:"newest_events"`
}


func Jsonparse() [3]float64{
    bytes, err := ioutil.ReadFile("./data/data.json")

    /*
    v := reflect.ValueOf(bytes)
    fmt.Println(v.Type())
    */
    
    if err != nil {
        log.Fatal(err)
    }

    var roomstate []RoomState
    if err := json.Unmarshal(bytes, &roomstate); err != nil {
        log.Fatal(err)
    }

    var rs [3]float64
    for _, p := range roomstate {
        //fmt.Printf("%s, Temp %.2f, ilumination %.2f, Hu %.2f\n", p.Name, p.NE.Te.Te_value, p.NE.Il.Il_value, p.NE.Hu.Hu_value)
        rs = [3]float64{p.NE.Te.Te_value, p.NE.Il.Il_value, p.NE.Hu.Hu_value}
    }

    return rs
}
