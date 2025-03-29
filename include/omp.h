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

    void* findFunc(const char* name);
    #endif

#ifdef __cplusplus
}

#include <string>
#include <unordered_map>

namespace omp_internal {
    void* getLibHandle();
    void setLibHandle(void* handle);
    std::unordered_map<std::string, void*>& getFuncs();
}

template <typename R, typename... Args>
R call(const std::string& funcName, Args... args)
{
    auto& funcsMap = omp_internal::getFuncs();
    auto it = funcsMap.find(funcName);
    void* funcAddr = nullptr;

    if (it == funcsMap.end()) {
        funcAddr = findFunc(funcName.c_str());
        funcsMap.emplace(funcName, funcAddr);
    } else {
        funcAddr = it->second;
    }

    typedef R (* FuncType)(Args...);
    FuncType func = (FuncType)funcAddr;

    return (*func)(std::forward<Args>(args)...);
}
#endif

#endif // GOMP_H
