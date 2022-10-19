package dom

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gorilla/websocket"
)
var(
	DEFAULT = Window{size:Size(300,300),pos:Pos(100,100)}
	upgrader = websocket.Upgrader{}
	port string = ":3333"
	rootws string = "/ws"
	contenido string  = "<html><head></head></html>"
	rootserv string = "http://localhost" + port + "/"
	com *websocket.Conn
	conection bool = false
	done bool = false
	methods []*method
	Event *Events
	Dom []*Element
	document *Element = &Element{ TagName: "document",}
	states []*_State_
	actions []func()
	AllEventListener []*EventListener
)
type method struct{
   name string
   pointer *method
   function func(*Events)
}
type smsJsons struct{
   Type string `json:"type"`
   Ref string `json:"ref"`
   Body string `json:"body"`
   Value string `json:"value"`
   Field string `json:"field"`
   Event string `json:"event"`
   Name string `json:"name"`
}
type Attrs struct{
	Type string
	Value string
}
// utils 
func Action(f func())(res func()){
	actions = append(actions,f)
	return f
}
func AddAction(funcs ...func())[]func(){
	for _,fn := range funcs{
		actions = append(actions, fn)
	}
	return actions
}
func Error(err error){
	if err != nil{
		fmt.Println(err)
	}
}
func Log( s ...interface{}){
	fmt.Println( s... )
}
func ToFirstUpperCase(str string)string{
	str = strings.ToTitle(str[:1]) + str[1:]
	return str
}
func Clean(s string)string{
	res := strings.ReplaceAll(strings.ReplaceAll(s , "\t",""),"\n","")
	if strings.Contains(res ,"> <"){
		res = strings.ReplaceAll(res ,"> <","><")
	}
	return res
}
// eventos
func onWindowLoad(call func()){
	for{
		if conection{
			call()
			return 
		}
	}
}
func OnWait(){
	for {
		if done == true {
			Log("Cerando...")
			os.Remove("./src/index.html")
			return 
		}
	}
}
func readCss()string{
	css,_ := os.ReadFile("./src/style.css")
	return string(css)
}
func writeConten(html string){
	file,_ := os.Create("./src/index.html")
	file.WriteString(html)
	file.Close()
}
func GetFile(path string)string{
	conten,err := (os.ReadFile(path))
	if err != nil{fmt.Println(err)}
	return Clean(string(conten))
}
// comunication
func send( sms string ){
	//tipo := "undefined"
	//if strings.Contains(sms,"bind"){ tipo = "is funtion" }else
	//if strings.Contains(sms,"addEventListener"){ tipo = "is event" }else
	//if strings.Contains(sms,"eval"){ tipo = "is evaluation" }
	for {
		if conection{
			_ = com.WriteMessage(1,[]byte(sms))
			//Log("message sent : " , sms  ,"   ", tipo)
			return
		}
	}
}
// falta implementar
func upload(s string ){
	print(s)
}
// functions server 
func reciver(w http.ResponseWriter, r *http.Request) {
	com, _ = upgrader.Upgrade(w, r, nil)
	defer com.Close()

	for {
		_, tempSMS, _ := com.ReadMessage()
		// Receive message
		evalOptions(string(tempSMS))
	}
}
func serv(w http.ResponseWriter, r *http.Request){
	w.Write([]byte(contenido))
	http.FileServer(http.Dir("./src/"))
}
func newServer() {
	
	http.HandleFunc(rootws, reciver)
	http.Handle("/" , http.FileServer(http.Dir("./src/")))
	log.Fatal(http.ListenAndServe(port, nil))
}
// add methods 
func AddMethod( name string , f func(*Events)){
	m := &method{ name:name,function:f }
	m.pointer = m
	methods = append(methods, m)
}
func evalMethods( sms string )bool{
	var Json smsJsons
	json.Unmarshal([]byte(sms), &Json)
	json.Unmarshal([]byte(Json.Event),&Event)
	for _,v := range methods{
		if v.name == Json.Name {
			v.function(Event)
			return true
		}
	}
	return false
}
// evaluation options reciver sw
func evalOptions(sms string){

	if sms == "ok"{
		ok()
		return
	}
	if sms == "close"{
		close( sms )
		return
	}
	if strings.Contains( sms , "upload"){
		upload( sms )
		return
	}
	evalMethods(sms)
}
func ok(){
	conection = true
	Log("conection is OK!")
}
func close( sms string ){
	done = true
}
// bind js
func Bind(name string , f func(*Events)){
	eval(`{"type":"bind","name":"`+ name +`"}`)
	AddMethod( name , f )
	return 
}
func eval( s string ){
	for {
		if conection{
			_ = com.WriteMessage(1,[]byte(s))
			//Log("message sent : " , s )
			return
		}
	}
}
// animation 
func getBoince()string{
	return `@keyframes bounceIn{0%{opacity: 0;transform: scale(0.3) translate3d(0,0,0);}50%{opacity: 0.9;transform: scale(1.1);}80%{opacity: 1;transform: scale(0.89);}100%{opacity: 1;transform: scale(1) translate3d(0,0,0);}}`
}
// selectors

func Delay(t time.Duration , callback func()){
	go func(){
		time.Sleep(t)
		callback()
	}()
}







