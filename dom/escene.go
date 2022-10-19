package dom

type _Escene_ struct{
  name string
  components []*Comp
  state bool
}
// contructor
func Escene(name string, initial bool)*_Escene_{
  return &_Escene_{ name:name,state:initial }
}
// getter slice components
func (this *_Escene_) AddAll(components ...*Comp)*_Escene_{
  this.components = append(this.components, components...)
  return this
}
// getter slice component
func (this *_Escene_) Add(component *Comp)*_Escene_{
  this.components = append(this.components, component)
  go this.AwaitCreateComp(this.reaction)
  return this
}
// change state this escene
func (this *_Escene_) Change()*_Escene_{
  this.state = !this.state 
  this.reaction()
  return this
}
// uploadDate of components
func (this *_Escene_) reaction(){
  for _,item := range this.components{
    comp := Selector(item.identifier)
    if this.state{
      comp.SetAttribute("style","display:initial")
    }else{
      comp.SetAttribute("style","display:none")
    }
  }
}
// getter state
func (this *_Escene_) GetState()bool{
  return this.state
}
// await if comp is create
func (this *_Escene_) AwaitCreateComp(f func()){
  Action(f)
}
