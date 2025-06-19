#include "vulkan.h"
#include <vulkan/vulkan.h>

uint32_t pfx_vk_version(int a, int b, int c) {Add commentMore actions
    return VK_MAKE_VERSION(a, b, c);
}