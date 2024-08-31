package server

func Init() {
	r := NetRouter()
	r.Run()
}
