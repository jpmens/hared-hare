#include <stdio.h>
#include <stdlib.h>
#include <stdbool.h>
#include <unistd.h>
#include <netdb.h>
#include <netinet/in.h>
#include <arpa/inet.h>
#include <memory.h>
#include <sys/socket.h>

/*
 * hare.c (C)2018 by Jan-Piet Mens <jp@mens.de>
 */

#define PORTNO  8035

#define env(K)	( getenv(K) ? getenv(K) : "<unknown>" )

int send_it(bool verbose, char *hostname, char *buf)
{
	struct addrinfo hints, *infoptr, *p;
	char host[256];
	int rc, sockfd;

	memset(&hints, 0, sizeof hints);
	hints.ai_family = AF_INET;
	hints.ai_socktype = SOCK_DGRAM;

	if ((rc = getaddrinfo(hostname, "8035", &hints, &infoptr)) != 0) {
		fprintf(stderr, "getaddrinfo: %s\n", gai_strerror(rc));
		return (1);
	}

	for (p = infoptr; p != NULL; p = p->ai_next) {
		getnameinfo(p->ai_addr, p->ai_addrlen, host, sizeof(host), NULL, 0, NI_NUMERICHOST);
		if (verbose) fprintf(stderr, "-> %s\n", host);

		sockfd = socket(AF_INET, SOCK_DGRAM, 0);

		if (sendto(sockfd, (void *)buf, strlen(buf), 0, p->ai_addr, p->ai_addrlen) < 0) {
			perror("sendto");
		}
		return (0);
	}

	freeaddrinfo(infoptr);

	return 0;
}


int main(int argc, char **argv)
{
	char *host = argv[1], buf[BUFSIZ], myhostname[BUFSIZ];
	char *pamtype = env("PAM_TYPE");	/* Linux */
	char *pamsm = env("PAM_SM_FUNC");	/* FreeBSD */
	bool verbose = false;

	if (strcmp(pamtype, "open_session") != 0 && strcmp(pamsm, "pam_sm_open_session") != 0) {
		fprintf(stderr, "Neither PAM open_session nor pam_sm_open_session detected\n");
		return 0;
	}

	if (argc != 2) {
		fprintf(stderr, "Usage: %s address\n", *argv);
		exit(2);
	}

	if (gethostname(myhostname, sizeof(myhostname)) != 0)
		strcpy(myhostname, "?");

	snprintf(buf, sizeof(buf), "%s login to %s from %s via %s",
		env("PAM_USER"),
		myhostname,
		env("PAM_RHOST"),
		env("PAM_SERVICE"));

	send_it(verbose, host, buf);
	return (0);
}
