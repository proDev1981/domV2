package ui

import (
	"app/dom"
	"app/model"
	"csv"
)

var database = dom.State(
  model.CsvToManoObra(
    csv.Open(`G:/Mi unidad/DB/src/operarios.csv`).Get(),
  ))

var escenePrueba = dom.Escene("prueba",false)

  var App = dom.NewComp(
    dom.Args{ 
    	"Control":Control,
    	"Slice":Slice,
    	"Title":Title, 
    	"Prueba":Prueba, 
    },
    `
      <div class='app'>
        </Title>
        </Prueba>
        </Control>
        </Slice>
      </div>
  `) 

