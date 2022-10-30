package handles

import "app/dom"
import "app/model"

var orden = true

// manejador boton reset
func Reverse(e *dom.Events){
  reverse := e.GetTarget()
  if orden{
      reverse.SetAttribute("style","transform : rotate(180deg)")
      orden = false
  }else{
      reverse.SetAttribute("style","transform : rotate(0deg)")
      orden = true
  }
  model.Database.Set(dom.Reverse[model.ManoObra](model.Database.Get()))
}
