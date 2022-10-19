package dom

import "fmt"
import "reflect"
import "strings"

//***** Args *****//
type Arg struct{
  Name string
  Value interface{}
}
// type slice of Arg
type Args map[string]any
// getter Arg by name
func GetArg[T any](data Args ,name string)(res T){
  return data["name"].(T)
}
//***** Render *****//
// render in stdout
func Render(ele InterfaceComponent){
  fmt.Println(ele.Render())
}
//***** Map *****//
// map slice and draw string 
func Map[T any](data any, html string )func()string{
  var d []T
  return func()(res string){
                switch data.(type){
                case *_State_:
                  d = data.(*_State_).Get().([]T)
                default:
                  d = data.([]T)
                }
                var rem string
                if len(d) > 0 {
                  for _,item := range d{
                    m := Entries(item)
                    rem = html
                    for key, value := range m {
                      if strings.Contains(html, "{{."+fmt.Sprint(key)+"}}"){
                        rem = strings.ReplaceAll(rem,"{{."+fmt.Sprint(key)+"}}",fmt.Sprint(value)) 
                      }
                    }
                    res += rem
                  }
                }
                return
          }
}
//***** Entries *****//
// get map entries the struct
func Entries(obj any )map[string]any{
  entries := make(map[string]any)
  t := reflect.TypeOf(obj)
  v := reflect.ValueOf(obj)
  _len := t.NumField()
  count := 0
  for count < _len{
    name := t.Field(count).Name 
    value := v.Field(count)
    entries[name] = value
   count++
  }
  return entries
}
//***** Try *****//
// handle error
func Try[T any](v T,e error)T{
  if e != nil { panic(e) }
  return v
}
