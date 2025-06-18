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

int pfx_ak_new_window(uint64_t wid, int width, int height, id *res, id *res_wh);

int pfx_ak_new_window(uint64_t wid, int width, int height, id *res);

void pfx_ak_close_requested_callback(uint64_t wid);

void gfx_ak_draw_callback(uint64_t wid, id drawable);

void gfx_ak_resize_callback(uint64_t wid, double width, double height);

void pfx_ak_draw_callback(uint64_t wid);

void pfx_ak_close_window(id w);

void pfx_ak_free_context(id w);