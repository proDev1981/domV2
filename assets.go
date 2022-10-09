package dom

/// foreach slice any
func ForEach[T any](data []T,f func(item T)string)(res string){
  for _,i := range data{
    res += f(i)
  }
  return
}
/// filter slice any
func Filter[T any](data []T, f func(T)bool)(newdata []T){
  for _,i := range data{
    if f(i){
      newdata = append(newdata,i)
    }
  }
  return newdata
}
