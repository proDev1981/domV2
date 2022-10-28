package ui

import "app/dom"
import "app/model"


// manejador boton filter
func handleFilter(e *dom.Events){
  res := dom.Filter(database.Get(),func(item model.ManoObra)bool{
              return dom.ContainsWordInAny(input, item.Fecha,item.Name,item.Obra)
          })
  database.Set(res)
}
