package handles

import "app/dom"
import "app/model"

var Query string

// manejador boton filter
func Filter(e *dom.Events){
  res := dom.Filter(model.Database.Get(),func(item model.ManoObra)bool{
              return dom.ContainsWordInAny(Query, item.Fecha,item.Name,item.Obra)
          })
  model.Database.Set(res)
}
