#include <stdint.h>

#if defined(__OBJC__)

#import <objc/objc.h>

#else
typedef void *id;
#endif

#define PFX_NOT_MAIN_THREAD (-10)
#define PFX_SUCCESS 1

int pfx_ak_run();

void pfx_ak_stop();

void pfx_ak_init_callback();

int pfx_ak_new_window(uint32_t wid, int width, int height, id *res);

void pfx_ak_close_requested_callback(uint32_t wid);

void pfx_ak_window_closed_callback(uint32_t wid);

void pfx_ak_close_window(id w);

void pfx_ak_free_context(id w);