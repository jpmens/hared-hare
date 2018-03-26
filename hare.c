#include <stdio.h>
#include <stdlib.h>
#include <stdbool.h>
#include <unistd.h>
#include <netdb.h>
#include <netinet/in.h>
#include <arpa/inet.h>
#include <memory.h>
#include <time.h>
#include <sys/socket.h>
#include "json.h"

/*
 * hare.c (C)2018 by Jan-Piet Mens <jp@mens.de>
 */

#define PORTNO  "8035"

#define env(K)	( getenv(K) ? getenv(K) : "<unknown>" )

int send_it(bool verbose, char *hostname, char *buf)
{
	struct addrinfo hints, *infoptr, *p;
	char host[256];
	int rc, sockfd;

	memset(&hints, 0, sizeof hints);
	hints.ai_family = AF_INET;
	hints.ai_socktype = SOCK_DGRAM;

	if ((rc = getaddrinfo(hostname, PORTNO, &hints, &infoptr)) != 0) {
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
	char *host = argv[1], myhostname[BUFSIZ];
	char *pamtype = env("PAM_TYPE");	/* Linux */
	char *pamsm = env("PAM_SM_FUNC");	/* FreeBSD */
	bool verbose = false;
	JsonNode *json;
	char *js;

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

	json = json_mkobject();
	json_append_member(json, "user", json_mkstring(env("PAM_USER")));
	json_append_member(json, "rhost", json_mkstring(env("PAM_RHOST")));
	json_append_member(json, "service", json_mkstring(env("PAM_SERVICE")));
	json_append_member(json, "hostname", json_mkstring(myhostname));
	json_append_member(json, "tst", json_mknumber(time(0)));

	if ((js = json_stringify(json, NULL)) != NULL) {
		send_it(verbose, host, js);
		free(js);
	}

	return (0);
}
