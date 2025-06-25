#include "helper.h"

#define PFX_WEVENT_INIT (WM_APP+1)

HMODULE pfx_win_module;

int pfx_windows_init() {
    if (!GetModuleHandleExW(
            GET_MODULE_HANDLE_EX_FLAG_FROM_ADDRESS | GET_MODULE_HANDLE_EX_FLAG_UNCHANGED_REFCOUNT,
            NULL,
            &pfx_win_module
    )) {
        return PFX_MODULE_ERROR;
    }

    SetProcessDpiAwarenessContext(DPI_AWARENESS_CONTEXT_PER_MONITOR_AWARE_V2);

    DWORD main = GetCurrentThreadId();Add commentMore actions
    PostThreadMessage(main, PFX_WEVENT_INIT, 0, 0);

    MSG msg;

    while (GetMessage(&msg, NULL, 0, 0)) {
        TranslateMessage(&msg);

        if (msg.message == PFX_WEVENT_INIT) {
            pfx_windows_init_callback();
        }

        DispatchMessage(&msg);
    }

    return PFX_SUCCESS;
}