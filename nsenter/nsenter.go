package nsenter

/*
#define _GNU_SOURCE
#include <errno.h>
#include <assert.h>
#include <stdio.h>
#include <fcntl.h>
#include <sched.h>
#include <stdlib.h>
#include <string.h>
#include <unistd.h>

__attribute__((constructor)) void enter_mnt() {
	char* pid = getenv("CT_PID");
	if (pid) {
		char str[256] = {0};
		char *namespaces[] = { "uts", "pid", "mnt" };
		for (int i = 0; i < sizeof(namespaces); i++) {
			sprintf(str, "/proc/%s/ns/%s", pid, namespaces[i]);
			int fd = open(str, O_RDONLY);
			assert(fd >= 0);
			if (setns(fd, 0) == -1) {
				fprintf(stderr, "setns on %s namespace failed: %s\n", namespaces[i], strerror(errno));
			} else {
				fprintf(stdout, "setns on %s namespace succeeded\n", namespaces[i]);
			}
			close(fd);
		}
	}
}
*/
import "C"
