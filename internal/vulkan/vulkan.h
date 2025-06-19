#ifndef GO_PFX_VULKAN_H
#define GO_PFX_VULKAN_H

#include <stdint.h>

uint32_t pfx_vk_version(int a, int b, int c);

#define PFX_VK_EXT_FUNC2(ID, ...) PFN_ ## ID func = (PFN_ ## ID)(vkGetInstanceProcAddr(instance, #ID)); func(__VA_ARGS__)

#define PFX_VK_EXT_FUNC(ID, ...) ((PFN_ ## ID)(vkGetInstanceProcAddr(instance, #ID)))(__VA_ARGS__)

#endif //GO_PFX_VULKAN_H