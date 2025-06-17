#import <Foundation/Foundation.h>
#import <AppKit/AppKit.h>
#import <QuartzCore/CAMetalLayer.h>
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

int pfx_ak_run() {
    @autoreleasepool {
        if (![NSThread isMainThread]) {
            return PFX_NOT_MAIN_THREAD;
        }

        [NSApplication sharedApplication];

        PfxApplicationDelegate *appDelegate = [[[PfxApplicationDelegate alloc] init] autorelease];

        // Ensure we are in multi-threading mode
        [NSThread detachNewThreadSelector:@selector(stubThread:)
                                 toTarget:appDelegate
                               withObject:nil];

        [NSApp setDelegate:appDelegate];

        [NSApp run];

        return PFX_SUCCESS;
    }
}

void pfx_ak_stop() {
    [NSApp stop:NSApp];
}

@class PfxWindow;
@class PfxWindowDelegate;
@class PfxView;

@interface PfxWindowContext : NSObject {
@public
    uint64_t wid;
    PfxWindow *window;
    PfxWindowDelegate *delegate;
    PfxView *view;
    CAMetalLayer *layer;
}

- (instancetype)initWithWID:(uint64_t)wid;
@end

@implementation PfxWindowContext
- (instancetype)initWithWID:(uint64_t)pwid {
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
    pfx_ak_window_closed_callback(context->wid);
}

@end

@interface PfxView : NSView <NSTextInputClient, CALayerDelegate> {
    PfxWindowContext *context;
}

- (instancetype)initWithContext:(PfxWindowContext *)ctx;

@end

@implementation PfxView
- (instancetype)initWithContext:(PfxWindowContext *)ctx {
    self = [super init];
    if (self != nil)
        context = ctx;

    [self setWantsLayer:YES];
    [self setLayerContentsRedrawPolicy:NSViewLayerContentsRedrawDuringViewResize];

    return self;
}

- (CALayer *)makeBackingLayer {
    context->layer = [CAMetalLayer layer];
    [context->layer setDelegate:self];
    return context->layer;
}

- (BOOL)canBecomeKeyView {
    return YES;
}

- (BOOL)acceptsFirstResponder {
    return YES;
}

- (BOOL)wantsUpdateLayer {
    return YES;
}

 - (void)updateLayer {
     pfx_ak_draw_callback(context->wid);
 }

- (void)displayLayer:(CALayer *)layer {
    pfx_ak_draw_callback(context->wid);
}

- (BOOL)canDrawSubviewsIntoLayer {
    return NO;
}

- (BOOL)hasMarkedText {
    return NO;
}

- (NSRange)markedRange {
    return NSMakeRange(NSNotFound, 0);
}

- (NSRange)selectedRange {
    return NSMakeRange(NSNotFound, 0);
}

- (NSRect)firstRectForCharacterRange:(NSRange)range actualRange:(nullable NSRangePointer)actualRange {
    return NSMakeRect(0, 0, 0, 0);
}

- (NSUInteger)characterIndexForPoint:(NSPoint)point {
    return 0;
}

- (NSArray<NSAttributedStringKey> *)validAttributesForMarkedText {
    return [NSArray array];
}

- (NSAttributedString *)attributedSubstringForProposedRange:(NSRange)range actualRange:(nullable NSRangePointer)actualRange {
    return nil;
}

- (void)insertText:(nonnull id)string replacementRange:(NSRange)replacementRange {
}


- (void)setMarkedText:(nonnull id)string selectedRange:(NSRange)selectedRange replacementRange:(NSRange)replacementRange {
}

- (void)unmarkText {
}

@end

int pfx_ak_new_window(uint64_t wid, int width, int height, id *res, id *res_wh) {
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
        [ctx->window setReleasedWhenClosed:NO];
        ctx->delegate = [[PfxWindowDelegate alloc] initWithContext:ctx];
        [ctx->window setDelegate:ctx->delegate];

        ctx->view = [[PfxView alloc] initWithContext:ctx];
        [ctx->window setContentView:ctx->view];
        [ctx->window makeFirstResponder:ctx->view];
        [ctx->view setNeedsDisplay:YES];

        [ctx->window setTitle:@"hello"];
        [ctx->window setRestorable:NO];
        [ctx->window setTabbingMode:NSWindowTabbingModeDisallowed];
        [ctx->window setCollectionBehavior:(NSWindowCollectionBehaviorFullScreenPrimary |
                                            NSWindowCollectionBehaviorManaged)];
        [ctx->window setAcceptsMouseMovedEvents:YES];

        [ctx->window center];
        [ctx->window makeKeyAndOrderFront:NSApp];
        [ctx->window orderFrontRegardless];

        *res_wh = ctx->layer;

        return PFX_SUCCESS;
    }
}

More actions
void pfx_ak_close_window(id w) {
    @autoreleasepool {
        PfxWindowContext *ctx = w;

        [ctx->window close];
    }
}

void pfx_ak_free_context(id w) {
    @autoreleasepool {
        PfxWindowContext *ctx = w;

        [ctx->window setDelegate:nil];
        [ctx->delegate release];
        [ctx->window setContentView:nil];
        [ctx->view release];
        [ctx->window release];
        [ctx release];
    }
}