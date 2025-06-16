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
@end

@interface PfxWindow : NSWindow {}
@end

@implementation PfxWindow
- (BOOL)canBecomeKeyWindow {
    return YES;
}

- (BOOL)canBecomeMainWindow {
    return YES;
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
        [NSApp setDelegate:appDelegate];

        [NSApp run];

        return PFX_SUCCESS;
    }
}

PfxWindow *window;

int pfx_ak_new_window(int width, int height) {
    @autoreleasepool {
        NSRect contentRect = NSMakeRect(0, 0, width, height);
        NSUInteger styleMask = NSWindowStyleMaskMiniaturizable | NSWindowStyleMaskTitled | NSWindowStyleMaskClosable |
                               NSWindowStyleMaskResizable;

        window = [[PfxWindow alloc]
                initWithContentRect:contentRect
                          styleMask:styleMask
                            backing:NSBackingStoreBuffered
                              defer:NO
        ];

        // todo: setcontentview, makefirstresponder setDelegate

        [window setTitle:@"hello"];
        [window setRestorable:NO];
        [window setTabbingMode:NSWindowTabbingModeDisallowed];
        [window setCollectionBehavior:(NSWindowCollectionBehaviorFullScreenPrimary |
                                       NSWindowCollectionBehaviorManaged)];
        [window setAcceptsMouseMovedEvents:YES];

        [window center];
        [window makeKeyAndOrderFront:NSApp];
        [window orderFrontRegardless];

        return PFX_SUCCESS;
    }
}