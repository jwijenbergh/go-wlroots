package wlroots

/*
 * This an unstable interface of wlroots. No guarantees are made regarding the
 * future consistency of this API.
 */

import "unsafe"

// #cgo pkg-config: wlroots wayland-server
// #cgo CFLAGS: -D_GNU_SOURCE -DWLR_USE_UNSTABLE
// #include <wlr/render/wlr_renderer.h>
import "C"

type Renderer struct {
	p *C.struct_wlr_renderer
}

func (r Renderer) Destroy() {
	C.wlr_renderer_destroy(r.p)
}

func (r Renderer) OnDestroy(cb func(Renderer)) {
	man.add(unsafe.Pointer(r.p), &r.p.events.destroy, func(unsafe.Pointer) {
		cb(r)
	})
}

func (r Renderer) InitDisplay(display Display) {
	C.wlr_renderer_init_wl_display(r.p, display.p)
}

func (r Renderer) Begin(output Output, width int, height int) {
	C.wlr_renderer_begin(r.p, C.uint(width), C.uint(height))
}

func (r Renderer) Clear(color *Color) {
	c := color.toC()
	C.wlr_renderer_clear(r.p, &c[0])
}

func (r Renderer) End() {
	C.wlr_renderer_end(r.p)
}

func (r Renderer) RenderTextureWithMatrix(texture Texture, matrix *Matrix, alpha float32) {
	m := matrix.toC()
	C.wlr_render_texture_with_matrix(r.p, texture.p, &m[0], C.float(alpha))
}

func (r Renderer) RenderRect(box *GeoBox, color *Color, projection *Matrix) {
	b := box.toC()
	c := color.toC()
	pm := projection.toC()
	C.wlr_render_rect(r.p, &b, &c[0], &pm[0])
}
