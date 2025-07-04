#include <stdint.h>
#include <stdbool.h>

#if defined(__OBJC__)

#import <objc/objc.h>

#else
typedef void *id;
#endif

#define PFX_SEE_ERROR (-10)
#define PFX_SUCCESS 1

int pfx_mtl_open(id *res, id *res_queue);

int pfx_mtl_configure_surface(id device, id layer, int pixelFormat);

void pfx_mtl_acquire_surface(id layer, id* res_draw, id* res_text);

void pfx_mtl_present_texture(id queue, id draw);

void pfx_mtl_discard_surface_texture(id draw);

int pfx_mtl_create_shader(id device, const void *src, int src_len, id *res_lib, char **res_err);

void pfx_mtl_get_shader_function(id lib, const void *name, int name_len, id *res);

void pfx_mtl_buffer_from_bytes(id device, const void *data, int data_len, id *res);

void pfx_mtl_create_command_buf(id queue, id *res);

typedef struct PipelineColorAttachment {
    int format;
} PipelineColorAttachment;

int pfx_mtl_create_render_pipeline(
        id device,
        id vertFunc,
        id fragFunc,
        const struct PipelineColorAttachment *colors,
        uint64_t colors_len,
        id *res_lib,
        char **res_err
);

typedef struct ColorAttachment {
    id view;
    bool load;
    bool store;
    double r;
    double g;
    double b;
    double a;
} ColorAttachment;

void pfx_mtl_begin_rpass(id buf, const struct ColorAttachment *colors, uint64_t colors_len, id *res);

void pfx_mtl_set_render_pipeline(id enc, id pipeline);

void pfx_mtl_set_vertex_buffer(id enc, id buffer);

void pfx_mtl_draw(id enc, int start, int count);

void pfx_mtl_end_rpass(id enc);

void pfx_mtl_cbuf_submit(id buffer);