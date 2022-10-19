package ui

import "app/dom"
import "app/model"
import "csv"

var input string
var orden = true

var Control = dom.NewComp(
  dom.Args{},
  `
    <div class='boxSearch'>
      <input class='search' type='text'></input>
      <button class='filter'>Go</button>
      <button class='reset'>↺</button>
      <button class='reverse'>≚</button>
    </div>
`)

func AControl(){
    filter := dom.Selector(".filter")
    reset := dom.Selector(".reset")
    search := dom.Selector(".search")
    reverse := dom.Selector(".reverse")
    // filter
    filter.AddEventListener("click",func(e *dom.Events){
        escenePrueba.Change()
        filter.SetInnerText("...")
        res := dom.Filter(database.Get(),func(item model.ManoObra)bool{
                    return dom.ContainsWordInAny(input, item.Fecha,item.Name,item.Obra)
                })
        database.Set(res)
    })
    // reset
    reset.AddEventListener("click",func(e *dom.Events){
        filter.SetInnerText("Go")
        search.SetAttribute("value","")
        database.Set(model.CsvToManoObra(
            csv.Open(`G:/Mi unidad/DB/src/operarios.csv`).Get(),
        ))
    })
    // reverse
    reverse.AddEventListener("click",func(e *dom.Events) {
        if orden{
            reverse.SetAttribute("style","transform : rotate(180deg)")
            orden = false
        }else{
            reverse.SetAttribute("style","transform : rotate(0deg)")
            orden = true
        }
        database.Set(dom.Reverse[model.ManoObra](database.Get()))
    })
    // search
    search.LinkVar(&input)

}


    
