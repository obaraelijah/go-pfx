#ifndef PFX_WINDOWS_H
#define PFX_WINDOWS_H

#undef WINVER
#undef _WIN32_WINNT

// Target Windows 10 upwards
#define WINVER 0x0A00
#define _WIN32_WINNT 0x0A00

#include <windows.h>
#include <stdint.h>

#define PFX_CLASS_ERROR (-1002)
#define PFX_CALL_ERROR (-1001)
#define PFX_MODULE_ERROR (-1000)
#define PFX_SEE_ERROR (-10)
#define PFX_SUCCESS 1

int pfx_windows_init();

void pfx_windows_init_callback();

int pfx_windows_new_window(
        uint64_t wid,
        LPCWSTR title,
        int width,
        int height
);

#endif //PFX_WINDOWS_H