package ui

import "app/dom"

var input string
var orden = true

var Control = dom.NewComp(`

    <div class='boxSearch'>
      <input class='search' type='text'></input>
      <button class='filter'>Go</button>
      <button class='reset'>↺</button>
      <button class='reverse'>≚</button>
    </div>

`,dom.Args{
            "@Click=>.filter":   handleFilter,
            "@Click=>.reset":    handleReset,
            "@Click=>.reverse" : handleReverse,
            "@Link=>.search":    &input,
})



    
