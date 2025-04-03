all: hare

hare: hare.c json.o
	$(CC) -Wall -Werror $(CFLAGS) -o hare hare.c json.o $(LDFLAGS)

json.o: json.c json.h

test: hare
	PAM_TYPE=open_session \
	PAM_USER=jjolie \
	PAM_RHOST=this.host \
	PAM_TTY=console \
	PAM_SERVICE=sshd ./hare 127.0.0.1

clean:
	rm -f *.o

clobber: clean
	rm -f hare
