#include <syslog.h>

void syslogWrapper(void* message) { syslog(LOG_ERR, "%s", (char*)message); }
