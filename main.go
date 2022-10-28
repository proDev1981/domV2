package main

import "app/dom"
import "app/ui"


func main(){
  dom.New(
    ui.App.Render(),
    dom.NewWindow().
      SetTitle("pruebaV2"),
  )
  dom.OnWait()
}
