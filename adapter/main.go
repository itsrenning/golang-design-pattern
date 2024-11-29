package main

func main() {
	cl := CurrentLogger{}
	cl.LogInformation("test current logger")

	nl := &NewLogger{}
	adapter := NewToCurrentAdapter{nl}
	adapter.NewLogger.LogInfo("test adapter for new logger")
}
