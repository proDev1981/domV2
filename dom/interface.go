package dom

import "fmt"
import "strings"
import "reflect"
import "regexp"
import "strconv"
import "log"

type InterfaceComponent interface{
  Render()string
  parse()string
  GetIdentifier()string
}
type Comp struct{
  props       Args
  html        string
  body        string
  header      string
  foother     string
  identifier  string
  typeIdentifier string
}
// create new component implement interface 
func NewComp(props Args, html string)*Comp{
  ele := &Comp{ props:props , html:html }
  return ele
}
// Render component
func (this *Comp) Render()(res string){
  res = this.parse()
  if this.identifier != "" { this.upDateComp(res) }
  return
}
// convert new element to data old element
func (this *Comp) upDateComp(res string){
      var newEle *Element
      ele := Selector(this.identifier)
      key := ele.GetAttribute("key")
      if key == "0" { 
        Dom = []*Element{}
        newEle = StrToDom(res,true)
      }else{
        newEle = StrToDom(res,false)
      }
      newEle.SetAttribute("key",key)
      newEle.SetRef(Try(strconv.Atoi(key)))
      ele.copyEventLsisternerNode(newEle)
      ele = newEle
      ele.SetOuterHTML(ele.GetOuterHTML())
      ele.uploadEventListenerNode()
}
/////debugger////
func getPosition(class string)(res []int){
  for index,item := range Dom{
    if strings.Contains(strings.Join(item.Attributes," "),class){
      res = append(res, index)
    }
  }
  return
}
////////////////
// get indetifier save in component
func (this *Comp) GetIdentifier()string{
  return this.identifier
}
// parse componet with states and varibles
func (this *Comp) parse()(html string){
  html = this.html

  for key,value := range this.props {
    if key == "@Map"{
      html = this.parseMap(value)
    }
    if key == "@Escene"{
      value.(*_Escene_).Add(this)
    }
    if strings.Contains(key,"@Link=>"){
      addLink(key,value)
    }
    if strings.Contains(key,"=>") && !strings.Contains(key,"Link"){
      addEvent(key,value)
    }
    switch value.(type){
    case *_State_:

      s := value.(*_State_)
      s.Add(this)
      html = strings.ReplaceAll(html,"{{."+key+"}}", fmt.Sprint(s.Get()))
    case InterfaceComponent:

      s := value.(InterfaceComponent)
      html = strings.ReplaceAll(html,"</"+key+">", fmt.Sprint(s.Render()))
    case func()string:

      s := value.(func()string)
      html = strings.ReplaceAll(html,"{{."+key+"}}", fmt.Sprint(s()))
    default:

      html = strings.ReplaceAll(html,"{{."+key+"}}", fmt.Sprint(value))
    }
  }
  go this.getIdentifier()
  return 
}
// parse component model map
func (this *Comp) parseMap(value any)string{
  var rv reflect.Value
  var res string
  this.descomposeTemplate()
  src := this.body
  varTemplate := this.getVarsNames()

  switch value.(type){
  case *_State_:

    d := value.(*_State_)
    rv = reflect.ValueOf(d.Get())

  default:
    fmt.Println("tipo distinto de state")
  }

  index := 0 ; for  index < rv.Len() {
      item := rv.Index(index)

      for _, name := range varTemplate{
        src = strings.ReplaceAll(src,"{{."+name+"}}",fmt.Sprint(item.FieldByName(name)))
      }
      res += fmt.Sprint("<div class='item'>",src,"</div>")
      src = this.body
      index++
  }
  return this.header + res + this.foother 
}
// get var names in template
func (this *Comp) getVarsNames()[]string{
  reg := Try(regexp.Compile(`\{\{\.[\w_]+\}\}`))
  match := reg.FindAllString(this.body,-1)

  for index := range match{
    match[index] = match[index][3:len(match[index])-2]
  }
  return match
}
// descompose container of template map
func (this *Comp) descomposeTemplate(){
  slice := strings.Split(this.html,">")
  this.header = strings.Join(slice[:1],">")+">"
  this.foother = strings.Join(slice[len(slice)-2:],">")
  this.body = strings.Join(slice[1:len(slice)-2],">")+">"
}
// get indentifier component for search in dom html
func (this *Comp) getIdentifier(){
  class := Try(regexp.Compile(`class\s*=\s*'[\w\d]+'`))
  id := Try(regexp.Compile(`id\s*=\s*'[\w\d]+'`))

  identifier := strings.Split(this.html,">")[0]
  identifier = strings.TrimSpace(identifier)
  identifier = strings.Join(strings.Split(identifier," ")[1:],"")

  switch {
  case strings.Contains(identifier,"class"):

    if !strings.Contains(identifier,"'"){ panic("Error syntax in attributes need `'`")}
    strClass := strings.TrimSpace(class.FindAllString(identifier,1)[0])
    strClass = fmt.Sprint(".",strings.Split(strClass,"'")[1])
    this.identifier = strClass
    this.typeIdentifier = "class"
  case strings.Contains(identifier, "id"):

    if !strings.Contains(identifier,"'"){ panic("Error syntax in attributes need `'`")}
    strId := strings.TrimSpace(id.FindAllString(identifier,1)[0])
    strId = fmt.Sprint("#",strings.Split(strId,"'")[1])
    this.identifier = strId
    this.typeIdentifier = "id"
  default:
    panic("Component need identifier!!")
  }
}
// add link in element
func addLink(name string, value any){
  switch value.(type){
  case *string:
    name = strings.Split(name,"=>")[1]
    Action(func(){
      Selector(name).LinkVar(value.(*string))
      log.Println("interface.go: ","link in ",name," value =>",value.(*string))
    })
  default:
    log.Println("Link event need string direction")
  }
}
// add event in element
func addEvent(name string,value any){
  switch value.(type){
  case func(*Events):
    Action(func(){
      log.Println("event name ",name, "=>",value.(func(*Events)))
    })
  default:
    log.Println("Error event need func and event direction")
  }
}
