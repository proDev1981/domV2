package ui

import "app/dom"


var Prueba = dom.NewComp(
  dom.Args{ "@Escene":escenePrueba },
  `
    <div class='prueba'>
      <h1>Hola tios</h1>
    </div>

  `,
)

