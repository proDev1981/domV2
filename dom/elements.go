package dom

import "fmt"
import "strings"
import "strconv"


type Element struct {
	ref 						int
	TagName    			string
	Value 					string
	Attributes 			[]string
	innerHtml  			string
	OuterHtml  			string
	InnerText 			string
	ParentNode 			*Element
	Children   			[]*Element
	eventlisteners 	[]*EventListener
	states 					[]*_State_
}
type EventListener struct{
	Name string
	Type string
	Action func(*Events)
	Parent *Element
}
/// push event listener in element
func (this *Element) pushEventListener(n,t string , f func(*Events)){
	this.eventlisteners = append(this.eventlisteners, &EventListener{ Name:n, Type:t, Action:f })
}
/// push state
func (this *Element) PushState(e *_State_){
	this.states = append(this.states,e)
}
/// search element for querySelector
func Selector(q string)(ele *Element){
	switch{
	case strings.Contains(q,"."):
		// class
		ele = SelectorClass(q)
	case strings.Contains(q,"#"):
		// id
		ele = SelectorId(q)
	case strings.Contains(q,"&"):
		// pos
		ele = SelectorPos(q)
	}
	return ele
}
/// search element by id
func SelectorId(id string)*Element{
	for _,item := range Dom{
		if strings.Contains(item.GetAttribute("id"),id[1:]){
			return item
		}
	}
	return &Element{}
}
/// search by class
func SelectorClass(class string)*Element{
	for _,item := range Dom{
		if strings.Contains(item.GetAttribute("class"),class[1:]){
			return item
		}
	}
	return &Element{}
}
/// search by pos
func SelectorPos(pos string)*Element{
	posInt,_ := strconv.Atoi(pos[1:]) 
	return Dom[posInt]
}
/// link value element frontend wicth backend
func (e *Element) LinkVar(value *string){
	e.AddEventListener("change",func(ev *Events){
		*value = ev.GetTarget().GetValue()
	})
}
/// link value element wicht event
func (e *Element) LinkValue(callback func(string)){
	e.AddEventListener("change",func(ev *Events){
		callback(ev.GetTarget().GetValue())
	})
	
}
/// setter attribute html
func (e *Element) SetAttribute(t, v string){
	existAttr := false
	for index,item := range e.Attributes{
		if strings.Contains(item,t){
			e.Attributes[index] = t + "='" + v + "'"
			existAttr = true
		}
	}
	if !existAttr { e.Attributes = append(e.Attributes, t + "='" + v + "'") }
	if t != "key" && t != "value"{

		js := "document.querySelector(`[key='"+ e.GetRef()+"']`).setAttribute('"+ t +"','"+ v +"')"
		eval (`{"type":"eval","js":"`+ js +`"}`)

	}else if t == "value"{

		js := "document.querySelector(`[key='"+ e.GetRef() +"']`).value = '"+ v +"'"
		eval(`{"type":"eval","js":"`+ js +`"}`)
	}
}
/// getter attribute html
func (e *Element) GetAttribute(n string)(string){
	if n == "value"{
		return e.Value
	}
	for _,item := range e.Attributes{
		if strings.Contains(item,n){
			return strings.ReplaceAll(
				strings.ReplaceAll(
					strings.Split(item,"=")[1],`"`,""),"'","")
		}
	}
	return "undefined"
}
/// getter ref element
func (e *Element) GetRef()string{
	return fmt.Sprint(e.ref)
}
/// setter ref element
func (e *Element) SetRef(v int){
	e.ref = v
}
/// getter inner html
func (e *Element) GetInnerHTML()string{
	return getInnerhtml(e)
}
/// setter inner html
func (e *Element) SetInnerHTML(html string){
	js := "document.querySelector(`[key='"+ e.GetRef() +"']`).innerHTML = `"+ html +"`"
	eval(`{"type":"eval","js":"`+ js +`"}`)
}
/// getter inner wicth tags element
func (e *Element) GetOuterHTML()(res string){
	res = getOutherhtml(e)
	for _,item := range e.states{
		res = strings.ReplaceAll(res,"$"+item.name,item.Get().(string))
	}
	return
}
/// setter outer element
func (e *Element) SetOuterHTML(html string){
	js := "document.querySelector(`[key='"+ e.GetRef() +"']`).outerHTML = `"+ html +"`"
	eval(`{"type":"eval","js":"`+ js +`"}`)
}
/// setter value element
func (e *Element) SetValue(v string){
	e.SetAttribute("value",v)
	js := "document.querySelector(`[key='"+ e.GetRef() +"']`).value = '"+ v +"'"
	eval(`{"type":"eval","js":"`+ js +`"}`)
}
/// setter innerText element
func (e *Element) SetInnerText(v string){
	e.InnerText = v
	js := "document.querySelector(`[key='"+ e.GetRef() +"']`).innerText = '"+ v +"'"
	eval(`{"type":"eval","js":"`+ js +`"}`)
}
/// getter value element
func (e *Element) GetValue()string{
	return e.Value
}
/// addd event listener element
func (e *Element) AddEventListener(t string, f func(*Events)){
	n := "method_"+fmt.Sprint(len(methods))
	e.pushEventListener(n,t,f)
	AllEventListener = append(AllEventListener, &EventListener{n,t,f,e})
	Bind(n,f)
	js := "document.querySelector(`[key='"+ e.GetRef()+"']`).addEventListener('"+ t +"',"+ n +")"
	eval (`{"type":"eval","js":"`+ js +`"}`)
}
/// upload event listener element
func (e *Element) uploadEventListeners(){
	for _,item := range e.eventlisteners{
		js := "document.querySelector(`[key='"+ e.GetRef()+"']`).addEventListener('"+ item.Type +"',"+ item.Name +")"
		eval (`{"type":"eval","js":"`+ js +`"}`)
	}
}
/// upload event listener the node
func (e *Element) uploadEventListenerNode(){
	e.Tree(func(item *Element){
		item.uploadEventListeners()
	})
}
func (e *Element) copyEventLsisternerNode(copy *Element){
	copy.eventlisteners = append(copy.eventlisteners, e.eventlisteners...)
	for _, item := range e.Children{
		item.copyEventLsisternerNode(copy)
	}
}
/// recorrer node y hacer algo
func (e *Element) Tree(f func(*Element)){
	f(e)
	for _,item := range e.Children{
		item.Tree(f)
	}
}
/// getter children
func (e *Element) GetChildren()[]*Element{
	return e.Children
}
/// DrawTree
func (e *Element)DrawElement(call func(int,*Element)){
  call(-1,e)
  for index,item := range e.GetChildren(){
    call(index,item)
  }
}
