package ui

import "app/dom"
import "app/model"
import "app/handles"

var FormGastos = dom.NewComp(`

    <div class='form gatos'>
      <input class='gastos_obra' placeholder='Obra..'></input>
      <input class='gastos_proveedor' placeholder='Proveedor..'></input>
      <input class='gastos_descripcion' placeholder='Descripcion..'></input>
      <input class='gastos_precio' placeholder='Precio..'></input>
      <div class='controler'>
        <button class='save'>Grabar</button>
        <button>Cancelar</button>
      </div>
    </div>

    `,dom.Args{
        "@Link=>.gastos_obra":&model.Gastos.Obra,
        "@Link=>.gastos_proveedor":&model.Gastos.Proveedor,
        "@Link=>.gastos_descripcion":&model.Gastos.Descripcion,
        "@Link=>.gastos_precio":&model.Gastos.Precio,
        "@Click=>.save": handles.Save,
    })

