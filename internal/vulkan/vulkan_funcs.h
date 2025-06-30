/* --- instance/device --- */

PFX_FUNC(VkResult, vkCreateInstance,
    (const VkInstanceCreateInfo*, pCreateInfo),
    (const VkAllocationCallbacks*, pAllocator),
    (VkInstance*, pInstance)
);

PFX_FUNC(VkResult, vkEnumeratePhysicalDevices,
    (VkInstance, instance),
    (uint32_t*, pPhysicalDeviceCount),
    (VkPhysicalDevice*, pPhysicalDevices)
);

PFX_FUNC(void, vkGetPhysicalDeviceProperties,
    (VkPhysicalDevice, physicalDevice),
    (VkPhysicalDeviceProperties*, pProperties)
);

PFX_FUNC(void, vkGetPhysicalDeviceQueueFamilyProperties,
    (VkPhysicalDevice, physicalDevice),
    (uint32_t*, pQueueFamilyPropertyCount),
    (VkQueueFamilyProperties*, pQueueFamilyProperties)
);

PFX_FUNC(VkResult, vkEnumerateDeviceExtensionProperties,
    (VkPhysicalDevice, physicalDevice),
    (const char*, pLayerName),
    (uint32_t*, pPropertyCount),
    (VkExtensionProperties*, pProperties)
);

PFX_FUNC(VkResult, vkCreateDevice,
    (VkPhysicalDevice, physicalDevice),
    (const VkDeviceCreateInfo*, pCreateInfo),
    (const VkAllocationCallbacks*, pAllocator),
    (VkDevice*, pDevice)
);

PFX_FUNC(void, vkGetDeviceQueue,
    (VkDevice, device),
    (uint32_t, queueFamilyIndex),
    (uint32_t, queueIndex),
    (VkQueue*, pQueue)
);

PFX_FUNC(VkResult, vkCreateDebugUtilsMessengerEXT,
    (VkInstance, instance),
    (const VkDebugUtilsMessengerCreateInfoEXT*, pCreateInfo),
    (const VkAllocationCallbacks*, pAllocator),
    (VkDebugUtilsMessengerEXT*, pMessenger)
);

/* --- surface --- */

PFX_FUNC(VkResult, vkCreateMetalSurfaceEXT,
    (VkInstance, instance),
    (const VkMetalSurfaceCreateInfoEXT*, pCreateInfo),
    (const VkAllocationCallbacks*, pAllocator),
    (VkSurfaceKHR*, pSurface)
);

PFX_FUNC(VkResult, vkCreateWin32SurfaceKHR,
    (VkInstance, instance),
    (const VkWin32SurfaceCreateInfoKHR*, pCreateInfo),
    (const VkAllocationCallbacks*, pAllocator),
    (VkSurfaceKHR*, pSurface)
);

PFX_FUNC(VkResult, vkGetPhysicalDeviceSurfaceCapabilitiesKHR,
    (VkPhysicalDevice, physicalDevice),
    (VkSurfaceKHR, surface),
    (VkSurfaceCapabilitiesKHR*, pSurfaceCapabilities)
);

PFX_FUNC(VkResult, vkGetPhysicalDeviceSurfaceFormatsKHR,
    (VkPhysicalDevice, physicalDevice),
    (VkSurfaceKHR, surface),
    (uint32_t*, pSurfaceFormatCount),
    (VkSurfaceFormatKHR*, pSurfaceFormats)
);

PFX_FUNC(VkResult, vkCreateSwapchainKHR,
    (VkDevice, device),
    (const VkSwapchainCreateInfoKHR*, pCreateInfo),
    (const VkAllocationCallbacks*, pAllocator),
    (VkSwapchainKHR*, pSwapchain)
);

PFX_FUNC(VkResult, vkGetSwapchainImagesKHR,
    (VkDevice, device),
    (VkSwapchainKHR, swapchain),
    (uint32_t*, pSwapchainImageCount),
    (VkImage*, pSwapchainImages)
);

PFX_FUNC(VkResult, vkAcquireNextImageKHR,
    (VkDevice, device),
    (VkSwapchainKHR, swapchain),
    (uint64_t, timeout),
    (VkSemaphore, semaphore),
    (VkFence, fence),
    (uint32_t*, pImageIndex)
);

PFX_FUNC(VkResult, vkQueuePresentKHR,
    (VkQueue, queue),
    (const VkPresentInfoKHR*, pPresentInfo)
);

/* --- sync --- */

PFX_FUNC(VkResult, vkCreateSemaphore,
    (VkDevice, device),
    (const VkSemaphoreCreateInfo*, pCreateInfo),
    (const VkAllocationCallbacks*, pAllocator),
    (VkSemaphore*, pSemaphore)
);

PFX_FUNC(VkResult, vkCreateFence,
    (VkDevice, device),
    (const VkFenceCreateInfo*, pCreateInfo),
    (const VkAllocationCallbacks*, pAllocator),
    (VkFence*, pFence)
);

PFX_FUNC(VkResult, vkWaitForFences,
    (VkDevice, device),
    (uint32_t, fenceCount),
    (const VkFence*, pFences),
    (VkBool32, waitAll),
    (uint64_t, timeout)
);

PFX_FUNC(VkResult, vkResetFences,
    (VkDevice, device),
    (uint32_t, fenceCount),
    (const VkFence*, pFences)
);

/* --- command --- */

PFX_FUNC(VkResult, vkCreateCommandPool,
    (VkDevice, device),
    (const VkCommandPoolCreateInfo*, pCreateInfo),
    (const VkAllocationCallbacks*, pAllocator),
    (VkCommandPool*, pCommandPool)
);

PFX_FUNC(VkResult, vkAllocateCommandBuffers,
    (VkDevice, device),
    (const VkCommandBufferAllocateInfo*, pAllocateInfo),
    (VkCommandBuffer*, pCommandBuffers)
);

PFX_FUNC(VkResult, vkBeginCommandBuffer,
    (VkCommandBuffer, commandBuffer),
    (const VkCommandBufferBeginInfo*, pBeginInfo)
);

PFX_FUNC(void, vkCmdBindPipeline,
    (VkCommandBuffer, commandBuffer),
    (VkPipelineBindPoint, pipelineBindPoint),
    (VkPipeline, pipeline)
);

PFX_FUNC(void, vkCmdDraw,
    (VkCommandBuffer, commandBuffer),
    (uint32_t, vertexCount),
    (uint32_t, instanceCount),
    (uint32_t, firstVertex),
    (uint32_t, firstInstance)
);

PFX_FUNC(VkResult, vkEndCommandBuffer,
    (VkCommandBuffer, commandBuffer)
);

/* --- graphics --- */

PFX_FUNC(VkResult, vkCreateShaderModule,
    (VkDevice, device),
    (const VkShaderModuleCreateInfo*, pCreateInfo),
    (const VkAllocationCallbacks*, pAllocator),
    (VkShaderModule*, pShaderModule)
);

PFX_FUNC(VkResult, vkCreatePipelineLayout,
    (VkDevice, device),
    (const VkPipelineLayoutCreateInfo*, pCreateInfo),
    (const VkAllocationCallbacks*, pAllocator),
    (VkPipelineLayout*, pPipelineLayout)
);

PFX_FUNC(VkResult, vkCreateGraphicsPipelines,
    (VkDevice, device),
    (VkPipelineCache, pipelineCache),
    (uint32_t, createInfoCount),
    (const VkGraphicsPipelineCreateInfo*, pCreateInfos),
    (const VkAllocationCallbacks*, pAllocator),
    (VkPipeline*, pPipelines)
);

PFX_FUNC(void, vkCmdBeginRenderingKHR,
    (VkCommandBuffer, commandBuffer),
    (const VkRenderingInfo*, pRenderingInfo)
);

PFX_FUNC(void, vkCmdEndRenderingKHR,
	(VkCommandBuffer, commandBuffer)
);

PFX_FUNC(VkResult, vkQueueSubmit2KHR,
    (VkQueue, queue),
    (uint32_t, submitCount),
    (const VkSubmitInfo2*, pSubmits),
    (VkFence, fence)
);

PFX_FUNC(void, vkCmdSetViewportWithCountEXT,
	(VkCommandBuffer, commandBuffer),
    (uint32_t, viewportCount),
    (const VkViewport*, pViewports)
);

PFX_FUNC(void, vkCmdSetScissorWithCountEXT,
	(VkCommandBuffer, commandBuffer),
    (uint32_t, scissorCount),
    (const VkRect2D*, pScissors)
);

PFX_FUNC(void, vkCmdPipelineBarrier2KHR,
	(VkCommandBuffer, commandBuffer),
    (const VkDependencyInfo*, pDependencyInfo)
);

/* --- resources --- */

PFX_FUNC(VkResult, vkCreateImageView,
    (VkDevice, device),
    (const VkImageViewCreateInfo*, pCreateInfo),
    (const VkAllocationCallbacks*, pAllocator),
    (VkImageView*, pView)
);