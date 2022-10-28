package ui

import "app/dom"
import "app/model"
import "csv"

// manejador boton reset
func handleReset(e *dom.Events){
  dom.Selector(".search").SetAttribute("value","")
  database.Set(model.CsvToManoObra(
      csv.Open(`G:/Mi unidad/DB/src/operarios.db`).Get(),
  ))
}

