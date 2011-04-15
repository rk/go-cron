include $(GOROOT)/src/Make.inc

TARG=cron
GOFILES=\
	cron.go\

include $(GOROOT)/src/Make.pkg

format:
	gofmt -s -spaces=true -tabindent=false -tabwidth=2 -w cron.go
