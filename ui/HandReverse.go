package ui

import "app/dom"
import "app/model"


// manejador boton reset
func handleReverse(e *dom.Events){
  reverse := e.GetTarget()
  if orden{
      reverse.SetAttribute("style","transform : rotate(180deg)")
      orden = false
  }else{
      reverse.SetAttribute("style","transform : rotate(0deg)")
      orden = true
  }
  database.Set(dom.Reverse[model.ManoObra](database.Get()))
}
