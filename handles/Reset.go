package handles

import "app/dom"
import "app/model"
import "csv"

// manejador boton reset
func Reset(e *dom.Events){
  dom.Selector(".search").SetAttribute("value","")
  model.Database.Set(model.CsvToManoObra(
      csv.Open(`G:/Mi unidad/DB/src/operarios.db`).Get(),
  ))
}

