package dom

import "strings"
import "log"
import "fmt"
/// reparar el renderiazado del padre e hijo segun el orden de renderiazado
/// se imprime el nombre de la variable no el valor

// cuando de renderiza un componente te que comprobar que el valor nuevo no es el mismo que el actual
// cuando renderizo un componente tengo que renderizar sus hijos 
// 
// types
type State struct {
	name string
	value any
	funcEleSubcribed func(...any)string
	subscribed []*Element
}
// states slice
func SliceState[T any](query string,data []T,f func(d ...any)string)(res *State){
  res = &State{ name:query, value:data, funcEleSubcribed:f }
  ele := Selector(query)
  ele.PushState(res)
  log.Println(ele)
  res.PushSubcribe(ele)
  states = append(states,res)
  return
}
/// states string
func NewState(name,value string)(res *State){
	res = &State{ name:name, value:value }
	for _,item :=range Dom{
		if 	strings.Contains(item.InnerText,"$"+name) || 
				strings.Contains(strings.Join(item.Attributes,""),"$"+name){
			res.PushSubcribe(item)
			item.PushState(res)
			res.uploadString(item)
		}
	}
	states = append(states, res)
	return
}
/// find state in array states by name
func FindState(name string)*State{
	for _,item := range states{
		if item.name == name { return item }
	}
	return nil
}
/// setter state
func (this *State) Set(value any)*State{
		this.value = value
		switch value.(type){
			case string:
				for _,item := range this.subscribed{
					this.uploadString(item)
				}
			default:
				for index,item := range this.subscribed{
					this.uploadSlice(item, index)
				}
		}
	return this
}
//
func SetState(name,value string){
	FindState(name).Set(value)
}
/// getter value state
func (this *State) Get()any{
	return this.value
}
//
func GetState(name,value string)any{
	return FindState(name).Get()
}
/// replace state in string httml
func Replace(ele *Element)(res string){
	for _,item := range ele.states{
		res = strings.ReplaceAll(ele.GetOuterHTML(),"$"+item.name,item.Get().(string))
	}
		return
}
/// comprovate if exist state by name
func stateExists(name string)bool{
	if FindState(name) != nil { return true }
	return false
}
/// push new subscribeded
func (this *State) PushSubcribe(n *Element){
	this.subscribed = append(this.subscribed,n)
}
/// upload html and event listener in element
func (this *State) uploadString(ele *Element){
	ele.SetOuterHTML(ele.GetOuterHTML())
	recursiveUpload(ele)
}
/// recursive upload events listeners
func recursiveUpload(ele *Element){
	ele.uploadEventListeners()
	for _,item := range ele.Children{
		recursiveUpload(item)
	} 
}
/// upload html and events listener in elements slice state
func (this *State) uploadSlice(ele *Element, pos int){
	if ele.TagName == ""{
		ele = Selector(this.name)
		this.subscribed[pos] = ele
		
	}
	res := strings.ReplaceAll(this.funcEleSubcribed(this.Get()),"\n","")
	res = fmt.Sprint("<",ele.TagName,ele.Attributes,">",res,"</",ele.TagName,">")
	ele.SetOuterHTML(res)
}
