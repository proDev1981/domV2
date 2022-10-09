package dom

import "fmt"
import "os/exec"
import "strings"
import "runtime"

type size struct{
	Width int
	Height int
}
type pos struct {
	PosX int
	PosY int
}
type Window struct {
	size size
	pos pos
	icon string
	title string
}

// window 
func NewWindow()*Window{
	return &Window{Size(300,300),Pos(100,100),"",""}
}
func (win *Window)SetSize(w int , h int )*Window{
	win.size.Width = w
	win.size.Height = h
	return win
}
func (w *Window)SetPosition(p pos)*Window{
	w.pos = p
	return w
}
func (w *Window)SetTitle(t string)*Window{
	w.title = t
	return w
}
func (w *Window)SetIcon(i string)*Window{
	w.icon = i
	return w
}
func (w *Window)PositionCenter() *Window{
	w.pos = Center()
	return w
}
func Size(w int , h int)size{
	return size{Width: w,Height: h}
}
func (p pos) dividir(c int)pos{
	p.PosX = p.PosX/c
	p.PosY = p.PosY/c
	return p
} 
func Pos(x int , y int)pos{
	return pos{PosX: x,PosY: y}
}
func Center()pos{
	return Pos(1300,800).dividir(4)
}
func SizeDefault()size{
	return Size(300,300)
}
func isWindows(w *Window) bool {

	cmd := exec.Command(
		"C:/Program Files (x86)/Microsoft/Edge/Application/msedge.exe", 
		"--app="+fmt.Sprintf("%s", rootserv),
		"--window-size="+fmt.Sprint(w.size.Width)+","+fmt.Sprint(w.size.Height),
		"--window-position="+fmt.Sprint(w.pos.PosX)+","+fmt.Sprint(w.pos.PosY),
	)
	if err := cmd.Start(); err != nil { // Ejecutar comando
		cmd := exec.Command(
			"c:/Program Files (x86)/Google/Chrome/Application/./chrome",
			"--app="+fmt.Sprintf("%s", rootserv),
			"--window-size="+fmt.Sprint(w.size.Width)+","+fmt.Sprint(w.size.Height),
			"--window-position="+fmt.Sprint(w.pos.PosX)+","+fmt.Sprint(w.pos.PosY),
		)
		if err := cmd.Start(); err != nil { // Ejecutar comando

			panic(err)
			return false
		}
	}
	return true
}
// control window
func New(content Component , w *Window){

	var strHTML string
	app := StrToDom(Build(content.model()))

	if !strings.Contains(app.GetOuterHTML(),"<html>") || !strings.Contains(app.GetOuterHTML(),"<body>"){
		strHTML = `<!DOCTYPE html><meta name="theme-color" content="#872e4e"><html><body>`+ app.GetOuterHTML() +`</html></body>`
	}
	
	strHTML = strings.Replace(strHTML,"<html>","<html><head><title>"+ w.title +"</title><link rel='icon' href='"+ w.icon +"' sizes='16x16 32x32' type='image/png'></head>", 1)
	
	
	strHTML = strings.Replace( strHTML,"<body>","<body><style>"+ readCss() +"</style>", 1)
	
	
	// inyect js
	strHTML = strings.Replace(strHTML,"<body>" , "<body>"+ js() , 1)
	contenido = strHTML
	writeConten(contenido)

	if runtime.GOOS == "windows" {
		isWindows(w)
	}
	// start server and window
	go newServer()
	go onWindowLoad(func(){
		// ejecutar action de todos los componentes
		for _,caller := range actions{
			go caller()
		}
		//uploadValues(replaceVar(simplyConten))
	})
}
