#include <hb.h>

extern void drawMoveTo(void *draw_data, float to_x, float to_y);
extern void drawLineTo(void *draw_data, float to_x, float to_y);
extern void drawQuadraticTo(void *draw_data, float control_x, float control_y, float to_x, float to_y);
extern void drawCubicTo(void *draw_data, float control1_x, float control1_y, float control2_x, float control2_y, float to_x, float to_y);
extern void drawClosePath(void *draw_data);

void move_to(hb_draw_funcs_t *dfuncs, void *draw_data, hb_draw_state_t *st, float to_x, float to_y, void *user_data)
{
	drawMoveTo(draw_data, to_x, to_y);
}

void line_to(hb_draw_funcs_t *dfuncs, void *draw_data, hb_draw_state_t *st, float to_x, float to_y, void *user_data)
{
	drawLineTo(draw_data, to_x, to_y);
}

void quadratic_to(hb_draw_funcs_t *dfuncs, void *draw_data, hb_draw_state_t *st, float control_x, float control_y, float to_x, float to_y, void *user_data)
{
	drawQuadraticTo(draw_data, control_x, control_y, to_x, to_y);
}

void cubic_to(hb_draw_funcs_t *dfuncs, void *draw_data, hb_draw_state_t *st, float control1_x, float control1_y, float control2_x, float control2_y, float to_x, float to_y, void *user_data)
{
	drawCubicTo(draw_data, control1_x, control1_y, control2_x, control2_y, to_x, to_y);
}

void close_path(hb_draw_funcs_t *dfuncs, void *draw_data, hb_draw_state_t *st, void *user_data)
{
	drawClosePath(draw_data);
}