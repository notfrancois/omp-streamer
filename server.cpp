#include "include/server.h"

extern "C" {
    #ifndef SERVER_FUNCTIONS_DEFINED
    #define SERVER_FUNCTIONS_DEFINED
    
    inline void server_printLnU8(const char* fmt) {
        return call<void>("server_printLnU8", fmt);
    }

    inline void server_logLnU8(int logLevel, const char* fmt) {
        return call<void>("server_logLnU8", logLevel, fmt);
    }

    inline void server_setData(int type, String data) {
        return call<void>("server_setData", type, data);
    }
    
    #endif
}
