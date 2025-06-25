#include "helper.h"

#define PFX_WEVENT_INIT (WM_APP+1)

HMODULE pfx_win_module;
ATOM pfx_win_class;

LRESULT pfx_window_proc(HWND hwnd, UINT uMsg, WPARAM wParam, LPARAM lParam);

HMODULE GetCurrentModuleHandle() {
    HMODULE ImageBase;
    if (GetModuleHandleExW(
            GET_MODULE_HANDLE_EX_FLAG_FROM_ADDRESS | GET_MODULE_HANDLE_EX_FLAG_UNCHANGED_REFCOUNT,
            (LPCWSTR) &GetCurrentModuleHandle,
            &ImageBase
    )) {
        return ImageBase;
    }
    return 0;
}

int pfx_windows_init() {
    pfx_win_module = GetCurrentModuleHandle();
    if (!pfx_win_module) {
        return PFX_MODULE_ERROR;
    }

    SetProcessDpiAwarenessContext(DPI_AWARENESS_CONTEXT_PER_MONITOR_AWARE_V2);

    WNDCLASSEXW wc = {sizeof(wc)};
    wc.style = CS_OWNDC;
    wc.lpfnWndProc = (WNDPROC) pfx_window_proc;
    wc.hInstance = pfx_win_module;
    wc.hCursor = LoadCursorW(NULL, (LPTSTR) IDC_ARROW);
    wc.lpszClassName = L"PFXWindowClass";

    pfx_win_class = RegisterClassExW(&wc);
    if (!pfx_win_class) {
        return PFX_CLASS_ERROR;
    }

    DWORD main = GetCurrentThreadId();
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

int pfx_windows_new_window(
        uint64_t wid,
        LPCWSTR title,
        int width,
        int height
) {
    HWND hwnd = CreateWindowExW(
            WS_EX_OVERLAPPEDWINDOW | WS_EX_APPWINDOW,
            MAKEINTATOM(pfx_win_class),
            title,
            WS_CLIPSIBLINGS | WS_CLIPCHILDREN | WS_TILEDWINDOW,
            CW_USEDEFAULT, CW_USEDEFAULT, width, height,
            NULL, NULL,
            pfx_win_module,
            NULL
    );
    if (hwnd == NULL) {
        return PFX_CALL_ERROR;
    }Add commentMore actions

    ShowWindow(hwnd, SW_NORMAL);

    return PFX_SUCCESS;
}

LRESULT pfx_window_proc(HWND hwnd, UINT uMsg, WPARAM wParam, LPARAM lParam) {
    return DefWindowProcW(hwnd, uMsg, wParam, lParam);
}