package vulkan

/*
#include <vulkan/vulkan.h>
#include "vulkan.h"
*/
import "C"
import (
	"log"
	"unsafe"
)

//export pfx_vk_debug_callback
func pfx_vk_debug_callback(
	messageSeverity C.VkDebugUtilsMessageSeverityFlagBitsEXT,
	messageTypes C.VkDebugUtilsMessageTypeFlagsEXT,
	pCallbackData *C.VkDebugUtilsMessengerCallbackDataEXT,
	pUserData unsafe.Pointer,
) C.VkBool32 {
	log.Println("pfx_vk_debug_callback")

	return 0
}
