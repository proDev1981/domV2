package ui

import (
	"app/dom"
	"app/model"
	"csv"
)

var database = dom.State(
  model.CsvToManoObra(
    csv.Open(`G:/Mi unidad/DB/src/operarios.db`).Get(),
  ))

  var App = dom.NewComp(
    dom.Args{ 
    	"Control":Control,
    	"Slice":Slice,
    	"Title":Title, 
    },
    `
      <div class='app'>
        </Title>
        </Control>
        </Slice>
      </div>
  `) 

