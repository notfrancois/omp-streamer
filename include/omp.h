#ifndef GOMP_H
#define GOMP_H

#include <stdbool.h>
#include <stdint.h>
#include <stddef.h>

typedef struct {
    const char* buf;
    size_t length;
} String;

typedef struct {
	void** buf;
	size_t length;
} Array;

typedef struct {
    float x;
    float y;
    float z;
    float w;
} Vector4;

typedef struct {
    float x;
    float y;
    float z;
} Vector3;

typedef struct {
    float x;
    float y;
} Vector2;

#ifdef __cplusplus
extern "C" {
#endif

    #ifndef OMP_FUNCTIONS_DEFINED
    #define OMP_FUNCTIONS_DEFINED
    void loadSdk();
    void unloadSdk();

    inline void* findFunc(const char* name);
    #endif

#ifdef __cplusplus
}

#include <string>
#include <unordered_map>

#ifndef OMP_VARIABLES_DEFINED
#define OMP_VARIABLES_DEFINED
extern void* libHandle;
extern std::unordered_map<std::string, void*> funcs;
#endif

template <typename R, typename... Args>
R call(const std::string& funcName, Args... args)
{
    auto it = funcs.find(funcName);
    void* funcAddr = nullptr;

    if (it == funcs.end()) {
        funcAddr = findFunc(funcName.c_str());
        funcs.emplace(funcName, funcAddr);
    } else {
        funcAddr = it->second;
    }

    // R ret;
    // if funcAddr == nullptr {
    //     return ret;
    // }

    typedef R (* FuncType)(Args...);

    FuncType func = (FuncType)funcAddr;

    return (*func)(std::forward<Args>(args)...);
}
#endif

#endif // GOMP_H
