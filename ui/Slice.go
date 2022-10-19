package ui

import "app/dom"

var Slice = dom.NewComp(
  dom.Args{ 
    "@Map":database, 
  },
  `
    <div class='slice'>
      <div class='fecha_name'>
        <span class='Name'>{{.Name}}</span>
        <span>{{.Fecha}}</span>
      </div>
      <span class='obra'>{{.Obra}}</span>
    </div>
`)

