package dom

import "runtime"
import "strings"

type Component struct{
	name string 
	model func()string
}

// components
func RenderDom(str string)Component{
	_,caller,_,_ := runtime.Caller(1)
	caller = strings.TrimSpace(caller[strings.LastIndex(caller,"/")+1:len(caller)-3])
	return Component{model:func()string{return str} ,name:caller}
}
func AddChilds(childs ...*Component){
	childsApp = append(childsApp,childs...)
}
func (e *Component) AddChilds(childs ...*Component){
	AddChilds(childs...)
}
func (e *Component) SetName(n string){
	e.name = n
}
