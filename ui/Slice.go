package ui

import "app/dom"

var Slice = dom.NewComp(`

    <div class='slice'>
      <div class='fecha_name'>
        <span class='Name'>{{.Name}}</span>
        <span>{{.Fecha}}</span>
      </div>
      <span class='obra'>{{.Obra}}</span>
    </div>

`,dom.Args{ "@Map":database })

