#if defined(WIN32) || defined(_WIN32) || defined(__WIN32__)
	#include <Windows.h>
#else
	#include <dlfcn.h>
#endif

#include "include/omp.h"
#include <iostream>
#include <string>

static void* libHandle = nullptr;
static std::unordered_map<std::string, void*> funcs;

namespace omp_internal {
    void* getLibHandle() {
        return libHandle;
    }

    void setLibHandle(void* handle) {
        libHandle = handle;
    }

    std::unordered_map<std::string, void*>& getFuncs() {
        return funcs;
    }
}

extern "C" {
    bool isLibraryAlreadyLoaded() {
#if defined(WIN32) || defined(_WIN32) || defined(__WIN32__)
        // Check if the library is already loaded in memory
        HMODULE hModule = GetModuleHandle("Go.dll");
        return (hModule != nullptr);
#else
        // On Unix-like systems, we can try to open with RTLD_NOLOAD
        // which only returns a handle if the library is already loaded
        void* handle = dlopen("./components/Go.so", RTLD_NOLOAD);
        if (handle != nullptr) {
            dlclose(handle);
            return true;
        }
        return false;
#endif
    }

    void loadSdk() {
        if (libHandle != nullptr) {
            return;
        }

        if (isLibraryAlreadyLoaded()) {
#if defined(WIN32) || defined(_WIN32) || defined(__WIN32__)
            libHandle = GetModuleHandle("Go.dll");
            if (libHandle == nullptr) {
                std::cerr << "Failed to get handle to already loaded Go.dll" << std::endl;
            }
#else
            libHandle = dlopen("./components/Go.so", RTLD_GLOBAL | RTLD_NOW);
            if (libHandle == nullptr) {
                std::cerr << "Failed to get handle to already loaded Go.so: " << dlerror() << std::endl;
            }
#endif
            return;
        }

        // Component not loaded anywhere, load it normally
#if defined(WIN32) || defined(_WIN32) || defined(__WIN32__)
        libHandle = LoadLibrary("./components/Go.dll");
        if (libHandle == nullptr) {
            std::cerr << "Failed to load Go.dll. Error code: " << GetLastError() << std::endl;
        }
#else
        libHandle = dlopen("./components/Go.so", RTLD_GLOBAL | RTLD_NOW);
        if (libHandle == nullptr) {
            std::cerr << "Failed to load Go.so: " << dlerror() << std::endl;
        }
#endif
    }

    void unloadSdk() {
        if (libHandle == nullptr) {
            return;
        }
        
        // Clear function pointer cache
        funcs.clear();
        
#if defined(WIN32) || defined(_WIN32) || defined(__WIN32__)
        FreeLibrary((HMODULE)libHandle);
#else
        dlclose(libHandle);
#endif
        libHandle = nullptr;
    }

    void* findFunc(const char* name) {
        if (libHandle == nullptr) {
            std::cerr << "Cannot find function. Library not loaded." << std::endl;
            return nullptr;
        }
        
        // Check if we've already looked up this function
        std::string funcName(name);
        auto it = funcs.find(funcName);
        if (it != funcs.end()) {
            return it->second;
        }
        
        // Look up the function and cache it
        void* funcPtr = nullptr;
#if defined(WIN32) || defined(_WIN32) || defined(__WIN32__)
        funcPtr = (void*)GetProcAddress((HMODULE)libHandle, name);
#else
        funcPtr = dlsym(libHandle, name);
#endif
        
        if (funcPtr != nullptr) {
            funcs[funcName] = funcPtr;
        } else {
#if defined(WIN32) || defined(_WIN32) || defined(__WIN32__)
            std::cerr << "Failed to find function " << name << ". Error code: " << GetLastError() << std::endl;
#else
            std::cerr << "Failed to find function " << name << ": " << dlerror() << std::endl;
#endif
        }
        
        return funcPtr;
    }
}
