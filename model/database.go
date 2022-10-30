package model

import "csv"
import "app/dom"


var Database = dom.State(
  CsvToManoObra(
    csv.Open(`G:/Mi unidad/DB/src/operarios.db`).Get(),
  ))
