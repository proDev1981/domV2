package model

type ManoObra struct{
  Name      string
  Fecha     string
  Obra      string
}

/// covert [][]string to ManoObra struct
func CsvToManoObra (data [][]string)([]ManoObra){
  var this []ManoObra
  for _,item := range data[1:]{
    if len(item) > 2{
      item[2] = item[2][:len(item[2])-1]// borro en caracter controlM de unix
      this = append(this,ManoObra{Name:item[1],Fecha:item[2],Obra:item[0]}) 
    }
  }
  return this
}

