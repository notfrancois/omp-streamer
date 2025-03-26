#if defined(WIN32) || defined(_WIN32) || defined(__WIN32__)
	#include <Windows.h>
#else
	#include <dlfcn.h>
#endif

#include "include/omp.h"

void* libHandle;
std::unordered_map<std::string, void*> funcs;

extern "C" {
    void loadComponent() {
        // Check if component is already loaded
        if (libHandle != nullptr) {
            return; // Component already loaded
        }

#if defined(WIN32) || defined(_WIN32) || defined(__WIN32__)
        libHandle = LoadLibrary("./components/Go.dll");
        if (libHandle == nullptr) {
            // Handle error - could log or throw exception
        }
#else
        libHandle = dlopen("./components/Go.so", RTLD_GLOBAL | RTLD_NOW);
        if (libHandle == nullptr) {
            // Handle error - could log error message from dlerror()
        }
#endif
    }

    void unloadComponent() {
#if defined(WIN32) || defined(_WIN32) || defined(__WIN32__)
        FreeLibrary((HMODULE)libHandle);
#else
        dlclose(libHandle);
#endif
    }

    void* findFunc(const char* name) {
#if defined(WIN32) || defined(_WIN32) || defined(__WIN32__)
        return (void*)GetProcAddress((HMODULE)libHandle, name);
#else
        return dlsym(libHandle, name);
#endif
    }
}
