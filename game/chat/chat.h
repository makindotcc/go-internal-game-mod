#include <stdlib.h>
#include <stdint.h>
#include <syslog.h>

typedef void (*fnPrintChat)(void* chatClient, char* text, unsigned int undefined1);

void printChat(long long chatClient, long long chatPrint, char* message) {
    fnPrintChat essa = (fnPrintChat)chatPrint;
    syslog(LOG_ERR, "makinsense | printing chat...............................");

    essa((void*)chatClient, message, 0);
}
