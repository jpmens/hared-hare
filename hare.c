#include <stdio.h>
#include <stdlib.h>
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

int main(int argc, char **argv)
{
	struct sockaddr_in servaddr;
	unsigned long addr;
	int sockfd;
	char *host = argv[1], buf[BUFSIZ], myhostname[BUFSIZ];
	char *pamtype = env("PAM_TYPE");	/* Linux */
	char *pamsm = env("PAM_SM_FUNC");	/* FreeBSD https://www.freebsd.org/cgi/man.cgi?query=pam_exec */

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

	addr = inet_addr(host);

	memset(&servaddr, 0, sizeof(servaddr));
	servaddr.sin_family = AF_INET;
	servaddr.sin_port = htons(PORTNO);
	memcpy((void *)&servaddr.sin_addr, (void *)&addr, sizeof(addr));

	sockfd = socket(AF_INET, SOCK_DGRAM, 0);

	if (sendto(sockfd, (void *)buf, strlen(buf), 0, (struct sockaddr *) &servaddr, sizeof(servaddr)) < 0) {
		perror("sendto");
	}
	return (0);
}
