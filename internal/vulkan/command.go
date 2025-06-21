package vulkan

import (
	"runtime"
	"unsafe"

	"github.com/obaraelijah/go-pfx/hal"
)

/*
#include <stdlib.h>
#include <vulkan/vulkan.h>
#include "vulkan.h"

VkClearValue pfx_vk_clear_color(float r, float g, float b, float a) {
	VkClearValue col;
	col.color.float32[0] = r;
	col.color.float32[1] = g;
	col.color.float32[2] = b;
	col.color.float32[3] = a;
	return col;
}

void pfx_vkCmdBeginRenderingKHR(
    VkInstance                                  instance,
	VkCommandBuffer                             commandBuffer,
    const VkRenderingInfo*                      pRenderingInfo
) {
	PFX_VK_EXT_FUNC(vkCmdBeginRenderingKHR, commandBuffer, pRenderingInfo);
}

void pfx_vkCmdEndRenderingKHR(
    VkInstance                                  instance,
	VkCommandBuffer                             commandBuffer
) {
	PFX_VK_EXT_FUNC(vkCmdEndRenderingKHR, commandBuffer);
}

VkResult pfx_vkQueueSubmit2KHR(
    VkInstance                                  instance,
	VkQueue                                     queue,
    uint32_t                                    submitCount,
    const VkSubmitInfo2*                        pSubmits,
    VkFence                                     fence
) {
	return PFX_VK_EXT_FUNC(vkQueueSubmit2KHR, queue, submitCount, pSubmits, fence);
}


void pfx_vkCmdSetViewportWithCountEXT(
    VkInstance                                  instance,
	VkCommandBuffer                             commandBuffer,
    uint32_t                                    viewportCount,
    const VkViewport*                           pViewports
) {
	PFX_VK_EXT_FUNC(vkCmdSetViewportWithCountEXT, commandBuffer, viewportCount, pViewports);
}

void pfx_vkCmdSetScissorWithCountEXT(
    VkInstance                                  instance,
	VkCommandBuffer                             commandBuffer,
    uint32_t                                    scissorCount,
    const VkRect2D*                             pScissors
) {
	PFX_VK_EXT_FUNC(vkCmdSetScissorWithCountEXT, commandBuffer, scissorCount, pScissors);
}

*/
import "C"

type CommandBuffer struct {
	graphics      *Graphics
	frame         *SurfaceFrame
	commandBuffer C.VkCommandBuffer
}

func (f *SurfaceFrame) CreateCommandBuffer() hal.CommandBuffer {
	var allocInfo C.VkCommandBufferAllocateInfo
	allocInfo.sType = C.VK_STRUCTURE_TYPE_COMMAND_BUFFER_ALLOCATE_INFO
	allocInfo.level = C.VK_COMMAND_BUFFER_LEVEL_PRIMARY
	allocInfo.commandPool = f.entry.commandPool
	allocInfo.commandBufferCount = 1

	var commandBuffer C.VkCommandBuffer

	if err := mapError(C.vkAllocateCommandBuffers(f.graphics.device, &allocInfo, &commandBuffer)); err != nil {
		panic(err)
	}

	var beginInfo C.VkCommandBufferBeginInfo
	beginInfo.sType = C.VK_STRUCTURE_TYPE_COMMAND_BUFFER_BEGIN_INFO
	beginInfo.flags = C.VK_COMMAND_BUFFER_USAGE_ONE_TIME_SUBMIT_BIT

	if err := mapError(C.vkBeginCommandBuffer(commandBuffer, &beginInfo)); err != nil {
		panic(err)
	}

	return &CommandBuffer{
		graphics:      f.graphics,
		frame:         f,
		commandBuffer: commandBuffer,
	}
}

func (c *CommandBuffer) BeginRenderPass(description hal.RenderPassDescriptor) {
	pinner := new(runtime.Pinner)
	defer pinner.Unpin()

	var cAttachs []C.VkRenderingAttachmentInfo

	for _, c := range description.ColorAttachments {
		tv, ok := c.View.(*TextureView)
		if !ok {
			panic("unexpected view type")
		}
		var colorAttachment C.VkRenderingAttachmentInfo
		colorAttachment.sType = C.VK_STRUCTURE_TYPE_RENDERING_ATTACHMENT_INFO
		colorAttachment.imageView = tv.view
		colorAttachment.imageLayout = C.VK_IMAGE_LAYOUT_COLOR_ATTACHMENT_OPTIMAL

		if c.Load {
			colorAttachment.loadOp = C.VK_ATTACHMENT_LOAD_OP_LOAD
		} else {
			colorAttachment.loadOp = C.VK_ATTACHMENT_LOAD_OP_CLEAR
		}

		if c.Discard {
			// TODO: change discard, as this may corrupt data?
			colorAttachment.loadOp = C.VK_ATTACHMENT_STORE_OP_NONE
		} else {
			colorAttachment.storeOp = C.VK_ATTACHMENT_STORE_OP_STORE
		}

		colorAttachment.clearValue = C.pfx_vk_clear_color(
			C.float(c.ClearColor.R),
			C.float(c.ClearColor.G),
			C.float(c.ClearColor.B),
			C.float(c.ClearColor.A),
		)

		cAttachs = append(cAttachs, colorAttachment)
	}

	var renderingInfo C.VkRenderingInfo
	renderingInfo.sType = C.VK_STRUCTURE_TYPE_RENDERING_INFO
	renderingInfo.renderArea.offset.x = 0
	renderingInfo.renderArea.offset.y = 0
	renderingInfo.renderArea.extent.width = C.uint32_t(c.frame.img.width)
	renderingInfo.renderArea.extent.height = C.uint32_t(c.frame.img.height)
	renderingInfo.layerCount = 1
	renderingInfo.colorAttachmentCount = C.uint32_t(len(cAttachs))
	renderingInfo.pColorAttachments = unsafe.SliceData(cAttachs)
	pinner.Pin(renderingInfo.pColorAttachments)

	C.pfx_vkCmdBeginRenderingKHR(c.graphics.instance, c.commandBuffer, &renderingInfo)

	var viewport C.VkViewport
	viewport.x = C.float(0)
	viewport.y = C.float(0)
	viewport.width = C.float(c.frame.img.width)
	viewport.height = C.float(c.frame.img.height)
	viewport.minDepth = C.float(0)
	viewport.maxDepth = C.float(1)

	C.pfx_vkCmdSetViewportWithCountEXT(c.graphics.instance, c.commandBuffer, 1, &viewport)

	var scissor C.VkRect2D
	scissor.offset.x = 0
	scissor.offset.y = 0
	scissor.extent.width = C.uint32_t(c.frame.img.width)
	scissor.extent.height = C.uint32_t(c.frame.img.height)

	C.pfx_vkCmdSetScissorWithCountEXT(c.graphics.instance, c.commandBuffer, 1, &scissor)
}

func (c *CommandBuffer) SetRenderPipeline(pipeline hal.RenderPipeline) {
	p, ok := pipeline.(*RenderPipeline)
	if !ok {
		panic("unexpected type")
	}

	C.vkCmdBindPipeline(c.commandBuffer, C.VK_PIPELINE_BIND_POINT_GRAPHICS, p.pipeline)
}

func (c *CommandBuffer) SetVertexBuffer(data hal.Buffer) {
	//TODO implement me
	panic("implement me")
}

func (c *CommandBuffer) Draw(start int, count int) {
	C.vkCmdDraw(c.commandBuffer, C.uint32_t(count), 1, C.uint32_t(start), 0)
}

func (c *CommandBuffer) EndRenderPass() {
	C.pfx_vkCmdEndRenderingKHR(c.graphics.instance, c.commandBuffer)
}

func (c *CommandBuffer) Submit() {
	pinner := new(runtime.Pinner)
	defer pinner.Unpin()

	C.vkEndCommandBuffer(c.commandBuffer)

	var cmdinfo C.VkCommandBufferSubmitInfo
	cmdinfo.sType = C.VK_STRUCTURE_TYPE_COMMAND_BUFFER_SUBMIT_INFO
	cmdinfo.pNext = nil
	cmdinfo.commandBuffer = c.commandBuffer
	cmdinfo.deviceMask = 0
	var waitInfo C.VkSemaphoreSubmitInfo
	waitInfo.sType = C.VK_STRUCTURE_TYPE_SEMAPHORE_SUBMIT_INFO
	waitInfo.pNext = nil
	waitInfo.semaphore = c.frame.entry.imgSem
	waitInfo.stageMask = C.VK_PIPELINE_STAGE_2_COLOR_ATTACHMENT_OUTPUT_BIT_KHR
	waitInfo.deviceIndex = 0
	waitInfo.value = 1

	var signalInfo C.VkSemaphoreSubmitInfo
	signalInfo.sType = C.VK_STRUCTURE_TYPE_SEMAPHORE_SUBMIT_INFO
	signalInfo.pNext = nil
	signalInfo.semaphore = c.frame.entry.completeSem
	signalInfo.stageMask = C.VK_PIPELINE_STAGE_2_ALL_GRAPHICS_BIT
	signalInfo.deviceIndex = 0
	signalInfo.value = 1

	var submitInfo C.VkSubmitInfo2
	submitInfo.sType = C.VK_STRUCTURE_TYPE_SUBMIT_INFO_2
	submitInfo.pNext = nil

	submitInfo.waitSemaphoreInfoCount = 1
	submitInfo.pWaitSemaphoreInfos = &waitInfo
	pinner.Pin(submitInfo.pWaitSemaphoreInfos)

	submitInfo.signalSemaphoreInfoCount = 1
	submitInfo.pSignalSemaphoreInfos = &signalInfo
	pinner.Pin(submitInfo.pSignalSemaphoreInfos)

	submitInfo.commandBufferInfoCount = 1
	submitInfo.pCommandBufferInfos = &cmdinfo
	pinner.Pin(submitInfo.pCommandBufferInfos)

	if err := mapError(C.pfx_vkQueueSubmit2KHR(c.graphics.instance, c.graphics.graphicsQueue, 1, &submitInfo, c.frame.entry.fence)); err != nil {
		panic(err)
	}

	// TODO: free buffer
}
