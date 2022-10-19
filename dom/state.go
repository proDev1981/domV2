package dom

import "log"
import "math/rand"
import "fmt"


type _State_ struct{
  name string
  value any
  subcribed []InterfaceComponent
}
// create _State_
func State(initial any)*_State_{
  return &_State_{ value:initial, name:fmt.Sprint(rand.Intn(100)) }
}
// getter value
func (this *_State_) Get()any{
  return this.value
}
// setter value
func (this *_State_) Set(v any){
  if fmt.Sprint(this.value) != fmt.Sprint(v) {
    this.value = v
    //log.Println("state.go: ","Change state value =>",this.value)
    for _,item := range this.subcribed{
      Render(item)
    }
  }
}
// add subcribed
func (this *_State_) Add(c InterfaceComponent){
  log.Println("state.go: ","add state",this.name," in =>",c.GetIdentifier() )
  this.subcribed = append(this.subcribed,c)
}
