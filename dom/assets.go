package dom

import "os"
import "log"
import "strings"

/// foreach slice any
func ForEach[T any](data []T,f func(item T)string)(res string){
  for _,i := range data{
    res += f(i)
  }
  return
}
/// filter slice any
func Filter[T any](data any, f func(T)bool)(newdata []T){
  d := data.([]T)
  for _,i := range d{
    if f(i){
      newdata = append(newdata,i)
    }
  }
  return newdata
}
/// get svg files
func GetSVG(path string)string{
  res,err := os.ReadFile(path)
  if err != nil { log.Println(err) }
  return string(res)
}
/// reverse slice
func Reverse[T any](data any)[]T{
  d := data.([]T)
  newdata := []T{}
  max := len(d)-1
  for index := range d{
    newdata = append(newdata,d[max-index])
  }
  return newdata
}
/// contains word in any strings
func ContainsWordInAny(s string, words ...string)bool{
  for _,word := range words{
    if strings.Contains(word, s){
      return true
    }
  }
  return false
}
/// debugger tool and states
func DebuggerTool[T any](tool T){
}
