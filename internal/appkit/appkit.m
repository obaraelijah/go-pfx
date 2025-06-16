#import <Foundation/Foundation.h>
#import <AppKit/AppKit.h>
#include "appkit.h"

@interface PfxApplicationDelegate : NSObject <NSApplicationDelegate>
@end

@implementation PfxApplicationDelegate
- (void)applicationDidFinishLaunching:(NSNotification *)notification {
    [NSApp setActivationPolicy:NSApplicationActivationPolicyRegular];

    pfx_ak_init_callback();
}

- (void)stubThread:(id)sender {
}
@end

PfxApplicationDelegate *appDelegate;

int pfx_ak_run() {
    @autoreleasepool {
        if (![NSThread isMainThread]) {
            return PFX_NOT_MAIN_THREAD;
        }

        [NSApplication sharedApplication];

        appDelegate = [[PfxApplicationDelegate alloc] init];

        // Ensure we are in multi-threading mode
        [NSThread detachNewThreadSelector:@selector(stubThread:)
                                 toTarget:appDelegate
                               withObject:nil];

        [NSApp setDelegate:appDelegate];

        [NSApp run];

        return PFX_SUCCESS;
    }
}

@class PfxWindow;

@interface PfxWindowContext : NSObject {
@public
    uint32_t wid;
    PfxWindow *window;
}

- (instancetype)initWithWID:(uint32_t)wid;
@end

@implementation PfxWindowContext
- (instancetype)initWithWID:(uint32_t)pwid {
    self = [super init];
    if (self != nil)
        self->wid = pwid;

    return self;
}
@end

@interface PfxWindow : NSWindow
@end

@implementation PfxWindow
- (BOOL)canBecomeKeyWindow {
    return YES;
}

- (BOOL)canBecomeMainWindow {
    return YES;
}
@end

@interface PfxWindowDelegate : NSObject <NSWindowDelegate> {
    PfxWindowContext *context;
}

- (instancetype)initWithContext:(PfxWindowContext *)ctx;
@end

@implementation PfxWindowDelegate
- (instancetype)initWithContext:(PfxWindowContext *)ctx {
    self = [super init];
    if (self != nil)
        context = ctx;

    return self;
}

- (BOOL)windowShouldClose:(NSWindow *)sender {
    pfx_ak_close_requested_callback(context->wid);

    return NO;
}

- (void)windowWillClose:(NSNotification *)notification {
    [context->window setDelegate:nil];

    pfx_ak_window_closed_callback(context->wid);

    [self release];
}

@end

int pfx_ak_new_window(uint32_t wid, int width, int height, id *res) {
    @autoreleasepool {
        PfxWindowContext *ctx = [[PfxWindowContext alloc] initWithWID:wid];
        *res = ctx;

        NSRect contentRect = NSMakeRect(0, 0, width, height);
        NSUInteger styleMask = NSWindowStyleMaskMiniaturizable | NSWindowStyleMaskTitled | NSWindowStyleMaskClosable |
                               NSWindowStyleMaskResizable;

        ctx->window = [[PfxWindow alloc]
                initWithContentRect:contentRect
                          styleMask:styleMask
                            backing:NSBackingStoreBuffered
                              defer:NO
        ];


        PfxWindowDelegate *delegate = [[PfxWindowDelegate alloc] initWithContext:ctx];
        [ctx->window setDelegate:delegate];

        // todo: setcontentview, makefirstresponder

        [ctx->window setTitle:@"hello"];
        [ctx->window setRestorable:NO];
        [ctx->window setTabbingMode:NSWindowTabbingModeDisallowed];
        [ctx->window setCollectionBehavior:(NSWindowCollectionBehaviorFullScreenPrimary |
                                            NSWindowCollectionBehaviorManaged)];
        [ctx->window setAcceptsMouseMovedEvents:YES];

        [ctx->window center];
        [ctx->window makeKeyAndOrderFront:NSApp];
        [ctx->window orderFrontRegardless];

        return PFX_SUCCESS;
    }
}