ifdef OS
   RM = del /Q
   FixPath = $(subst /,\,$1)
else
   ifeq ($(shell uname), Linux)
      RM = rm -f
      FixPath = $1
   endif
endif

jugador:
	go run Jugador/client/client.go

lider:
	go run Lider/server/server.go

namenode:
	go run NameNode/server/server.go

pozo:
	go run Pozo/server.go

clean:
	$(RM) *.txt
