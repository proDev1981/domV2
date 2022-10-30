package ui

import "app/dom"
import "app/handles"


var Control = dom.NewComp(`

    <div class='boxSearch'>
      <input class='search' type='text'></input>
      <button class='filter'>Go</button>
      <button class='reset'>↺</button>
      <button class='reverse'>≚</button>
    </div>

`,dom.Args{
            "@Click=>.filter":   handles.Filter,
            "@Click=>.reset":    handles.Reset,
            "@Click=>.reverse" : handles.Reverse,
            "@Link=>.search":    &handles.Query,
})



    
