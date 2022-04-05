up:
	./main
run-fiber:
	go test -benchmem -bench RequestFiber -run=^$
run-fasthttp:
	go test -benchmem -bench RequestFastHttpRouting -run=^$
run-gin:
	go test -benchmem -bench RequestGin -run=^$
run-echo:
	go test -benchmem -bench RequestEcho -run=^$
run-mux:
	go test -benchmem -bench RequestMux -run=^$
run-chi:
	go test -benchmem -bench RequestChi -run=^$