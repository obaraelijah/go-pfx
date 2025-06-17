#import <Metal/MTLDevice.h>
#include "metal.h"

int pfx_mtl_open(id *res) {
    @autoreleasepool {
        id<MTLDevice> device = MTLCreateSystemDefaultDevice();

        *res = device;

        return PFX_SUCCESS;
    }
}