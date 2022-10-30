package dom

import "strings"
import "fmt"
import "regexp"

var attrs01 = `[\w\d]+\s*=\s*"[\w\d()$'.,\s]+"`// regexp atributos con comillas dobles
var attrs02 = `[\w\d]+\s*=\s*'[\w\d()$".,\s]+'`// regexp atributos con comillas simples
var attrs03 = `\s+[\w\d]+\s*=\s*[\w\d$]+\s*`// regexp atributos sin comillas
var contRef int

// stack struct
type Stack struct{
  data []*Element
}
// push new element stack
func (e *Stack) Push(ele *Element){
  e.data = append(e.data, ele)
}
// extract latest element stack
func (e *Stack) Pop()*Element{
  res := e.Get()
  e.data = e.data[:len(e.data)-1]
  return res
}
// return latest element stack
func (e *Stack) Get()*Element{
  return e.data[len(e.data)-1]
}
// return len of Stack
func (e *Stack) Len()int{
  return len(e.data)
}
// return boolean value if value in str 
func In(str , value string)bool{
  return strings.Contains(str,value)
}
// insert marck in string for delimitation inner text best easy
func insertMark(str string)string{
  str = strings.ReplaceAll(str,">",">%")
  str = strings.ReplaceAll(str,"<","%<")
  return str
}
// clean items emptys
func cleanItemsEmptys(v []string)[]string{
  var newSlice []string
  for _,item := range v{
    item = strings.TrimSpace(item)
    if len(item) > 0 { newSlice = append(newSlice,item) }
  }
  return newSlice
}
// trim space  \n and \t
func trimSpace(str string)string{
  str = strings.TrimSpace(str)
  str = strings.ReplaceAll(str,"\n","")
  return strings.ReplaceAll(str,"\t","")
}
// extract atributes in array string
func extractAttrs(attr string)(match []string){

    op1,_ := regexp.Compile(attrs01)
    op2,_ := regexp.Compile(attrs02)
    op3,_ := regexp.Compile(attrs03)
    
    match_op1 := op1.FindAllString(attr,-1)
    match_op2 := op2.FindAllString(attr,-1)
    match_op3 := op3.FindAllString(attr,-1)

    match = append(match,match_op1...) 
    match = append(match,match_op2...)
    match = append(match,match_op3...)

    return
}
// fill attributes element
func fillAttrs(str string, ele *Element)*Element{
  var name string
  var attr string
  var match []string

  sep := strings.Index(str," ")
  if sep > 0 {
    name = str[1:sep]
    attr = str[sep:len(str)-1]
    // extraer atributos con regex
    match = extractAttrs(attr)    
  }else{
    name = str[1:len(str)-1]
  }
  ele.TagName = name
  ele.SetRef(contRef)
  ele.Attributes = append(match," key='"+fmt.Sprint(ele.ref)+"'")
  contRef++
  return ele
}
// get innerhtml of element
func getInnerhtml(ele *Element)string{
  inner := ele.InnerText
  for _,child := range ele.Children{
    inner += child.GetOuterHTML()
  }
  return inner
}
func getOutherhtml(ele *Element)string{
  return fmt.Sprint("<",ele.TagName," ",strings.Join(ele.Attributes," "),">",ele.GetInnerHTML(),"</",ele.TagName,">")
}
// convert string html in struct dom
func StrToDom(str string,save bool)*Element{

  var stack Stack
  var body *Element
  if save { contRef = 0 }

  str = insertMark(trimSpace(str))// inserto marcas para poder diferenciar el inner text
  cont := cleanItemsEmptys(strings.Split(str, "%"))
  for _,item := range cont{
    switch {

    case  !In(item ,"</") && In(item,">"):
      if stack.Len() == 0 {// si es 0 es el primer elemento lo inserto a la pila y al dom 
        ele := &Element{}
        stack.Push(fillAttrs(item,ele)) 
        body = stack.Get()
      }else if stack.Len() > 0 {// si es mayor inserto como hijo del actual stack en la pila y lo añado a la pila 
        ele := &Element{}
        ele = fillAttrs(item,ele)
        ele.ParentNode = stack.Get()
        stack.Get().Children = append(stack.Get().Children,ele)
        stack.Push(ele)
      }
      if save{
        Dom = append(Dom,stack.Get())// añadir elemento a lista Dom 
      }

    case  In(item,"</") && In(item,">"):
      // extraer ultima posicion del stack
      stack.Pop()

    default:
      // insertar inner en ultimo elemento del stack
      stack.Get().InnerText = item
    }
    
  }
  if stack.Len() != 0 { print("html mal formado , falta cierre o apertura ") }
  return body
}  


