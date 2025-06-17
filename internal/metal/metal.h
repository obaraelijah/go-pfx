#include <stdint.h>

#if defined(__OBJC__)

#import <objc/objc.h>

#else
typedef void *id;
#endif

#define PFX_SUCCESS 1

int pfx_mtl_open(id *res, id *res_queue);

int pfx_mtl_configure_surface(id device, id layer);

void pfx_mtl_acquire_texture(id layer, id *res);

void pfx_mtl_present_texture(id queue, id text);

void pfx_mtl_discard_surface_texture(id text);