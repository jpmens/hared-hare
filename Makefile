CFLAGS=-Wall -Werror

all: hare

hare: hare.c
	$(CC) $(CFLAGS) -o hare hare.c $(LDFLAGS)

test: hare
	PAM_TYPE=open_session \
	PAM_USER=username \
	PAM_HOST=this.host \
	PAM_SERVICE=test ./hare 192.168.1.130
