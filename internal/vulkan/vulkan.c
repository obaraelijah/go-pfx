#include "vulkan.h"
#include "volk.c"

uint32_t pfx_vk_version(int a, int b, int c) {
    return VK_MAKE_VERSION(a, b, c);
}

#undef PFX_FUNC
#define PFX_FUNC(RES, NAME, ...) RES pfx_ ## NAME ( MAP_LIST(PFX_PARAMS_PAIR, __VA_ARGS__) ) {  return NAME (  MAP_LIST(PFX_PARAM_NAMES_PAIR, __VA_ARGS__)   ); }

#include "vulkan_funcs.h"