#include <stdint.h>

#if defined(__OBJC__)

#import <objc/objc.h>

#else
typedef void *id;
#endif

#define PFX_SUCCESS 1

int pfx_mtl_open(id *res);