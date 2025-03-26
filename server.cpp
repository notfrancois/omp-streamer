#include "include/server.h"

extern "C" {
    void server_printLnU8(const char* fmt) {
        return call<void>("server_printLnU8", fmt);
    }

    void server_logLnU8(int logLevel, const char* fmt) {
        return call<void>("server_logLnU8", logLevel, fmt);
    }

    void server_setData(int type, String data) {
        return call<void>("server_setData", type, data);
    }
}
