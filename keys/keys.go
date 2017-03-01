package keys

import "github.com/veandco/go-sdl2/sdl"

// Key is a specific key.
type Key int

const (
	// Space is the space bar.
	Space = Key(sdl.K_SPACE)
	// // Apostrophe is the apostrophe.
	// Apostrophe = Key(glfw.KeyApostrophe)
	// Comma ist the comma.
	Comma = Key(sdl.K_COMMA)
	// Minus is the minus key.
	Minus = Key(sdl.K_MINUS)
	// Period is the period key.
	Period = Key(sdl.K_PERIOD)
	// Slash is the slash key.
	Slash = Key(sdl.K_SLASH)
	// N0 is the number Key 0.
	N0 = Key(sdl.K_0)
	// N1 is the number Key 1.
	N1 = Key(sdl.K_1)
	// N2 is the number Key 2.
	N2 = Key(sdl.K_2)
	// N3 is the number Key 3.
	N3 = Key(sdl.K_3)
	// N4 is the number Key 4.
	N4 = Key(sdl.K_4)
	// N5 is the number Key 5.
	N5 = Key(sdl.K_5)
	// N6 is the number Key 6.
	N6 = Key(sdl.K_6)
	// N7 is the number Key 7.
	N7 = Key(sdl.K_7)
	// N8 is the number Key 8.
	N8 = Key(sdl.K_8)
	// N9 is the number Key 9.
	N9 = Key(sdl.K_9)
	// Semicolon is the semicolon.
	Semicolon = Key(sdl.K_SEMICOLON)
	// Equal is the equal key.
	Equal = Key(sdl.K_EQUALS)
	// A key.
	A = Key(sdl.K_a)
	// B key.
	B = Key(sdl.K_b)
	// C key.
	C = Key(sdl.K_c)
	// D key.
	D = Key(sdl.K_d)
	// E key.
	E = Key(sdl.K_e)
	// F key.
	F = Key(sdl.K_f)
	// G key.
	G = Key(sdl.K_g)
	// H key.
	H = Key(sdl.K_h)
	// I key.
	I = Key(sdl.K_i)
	// J key.
	J = Key(sdl.K_j)
	// K key.
	K = Key(sdl.K_k)
	// L key.
	L = Key(sdl.K_l)
	// M key.
	M = Key(sdl.K_m)
	// N key.
	N = Key(sdl.K_n)
	// O key.
	O = Key(sdl.K_o)
	// P key.
	P = Key(sdl.K_p)
	// Q key.
	Q = Key(sdl.K_q)
	// R key.
	R = Key(sdl.K_r)
	// S key.
	S = Key(sdl.K_s)
	// T key.
	T = Key(sdl.K_t)
	// U key.
	U = Key(sdl.K_u)
	// V key.
	V = Key(sdl.K_v)
	// W key.
	W = Key(sdl.K_w)
	// X key.
	X = Key(sdl.K_x)
	// Y key.
	Y = Key(sdl.K_y)
	// Z key.
	Z = Key(sdl.K_z)
	// LeftBracket is the opeening brakcet.
	LeftBracket = Key(sdl.K_LEFTBRACKET)
	// Backslash is the backslash key.
	Backslash = Key(sdl.K_BACKSLASH)
	// RightBracket is the closing brakcet.
	RightBracket = Key(sdl.K_RIGHTBRACKET)
	// // GraveAccent is the grave accent.
	// GraveAccent = Key(sdl.K_)
	// Esc ist the escape key.
	Esc = Key(sdl.K_ESCAPE)
	// Enter ist the enter key.
	Enter = Key(sdl.K_KP_ENTER)
	// Tab ist the tabulator.
	Tab = Key(sdl.K_TAB)
	// Backspace ist the backspace.
	Backspace = Key(sdl.K_BACKSPACE)
	// Insert ist the insert key.
	Insert = Key(sdl.K_INSERT)
	// Delete ist the delete key.
	Delete = Key(sdl.K_DELETE)
	// Right ist the right cursor.
	Right = Key(sdl.K_RIGHT)
	// Left ist the left cursor.
	Left = Key(sdl.K_LEFT)
	// Down ist the down cursor.
	Down = Key(sdl.K_DOWN)
	// Up ist the up cursor.
	Up = Key(sdl.K_UP)
	// PageUp ist the page up cursor.
	PageUp = Key(sdl.K_PAGEUP)
	// PageDown ist the page down cursor.
	PageDown = Key(sdl.K_PAGEDOWN)
	// Home ist the home key.
	Home = Key(sdl.K_HOME)
	// End ist the end key.
	End = Key(sdl.K_END)
	// CapsLock ist the caps lock key.
	CapsLock = Key(sdl.K_CAPSLOCK)
	// ScrollLock ist the scroll lock key.
	ScrollLock = Key(sdl.K_SCROLLLOCK)
	// NumLock ist the num lock key.
	NumLock = Key(sdl.K_NUMLOCKCLEAR)
	// PrintScreen ist the print screen key.
	PrintScreen = Key(sdl.K_PRINTSCREEN)
	// Pause ist the pause key.
	Pause = Key(sdl.K_PAUSE)
	// F1 ist the F1 key.
	F1 = Key(sdl.K_F1)
	// F2 ist the F2 key.
	F2 = Key(sdl.K_F2)
	// F3 ist the F3 key.
	F3 = Key(sdl.K_F3)
	// F4 ist the F4 key.
	F4 = Key(sdl.K_F4)
	// F5 ist the F5 key.
	F5 = Key(sdl.K_F5)
	// F6 ist the F6 key.
	F6 = Key(sdl.K_F6)
	// F7 ist the F7 key.
	F7 = Key(sdl.K_F7)
	// F8 ist the F8 key.
	F8 = Key(sdl.K_F8)
	// F9 ist the F9 key.
	F9 = Key(sdl.K_F9)
	// F10 ist the F10 key.
	F10 = Key(sdl.K_F10)
	// F11 ist the F11 key.
	F11 = Key(sdl.K_F11)
	// F12 ist the F12 key.
	F12 = Key(sdl.K_F12)
	// F13 ist the F13 key.
	F13 = Key(sdl.K_F13)
	// F14 ist the F14 key.
	F14 = Key(sdl.K_F14)
	// F15 ist the F15 key.
	F15 = Key(sdl.K_F15)
	// F16 ist the F16 ky.
	F16 = Key(sdl.K_F16)
	// F17 ist the F17 key.
	F17 = Key(sdl.K_F17)
	// F18 ist the F19 key.
	F18 = Key(sdl.K_F18)
	// F19 ist the F19 key.
	F19 = Key(sdl.K_F19)
	// F20 ist the F20 key.
	F20 = Key(sdl.K_F20)
	// F21 ist the F21 key.
	F21 = Key(sdl.K_F21)
	// F22 ist the F22 key.
	F22 = Key(sdl.K_F22)
	// F23 ist the F23 key.
	F23 = Key(sdl.K_F23)
	// F24 ist the F24 key.
	F24 = Key(sdl.K_F24)
	// KP0 ist the key 0 on the key pad.
	KP0 = Key(sdl.K_KP_0)
	// KP1 ist the key 1 on the key pad.
	KP1 = Key(sdl.K_KP_1)
	// KP2 ist the key 2 on the key pad.
	KP2 = Key(sdl.K_KP_2)
	// KP3 ist the key 3 on the key pad.
	KP3 = Key(sdl.K_KP_3)
	// KP4 ist the key 4 on the key pad.
	KP4 = Key(sdl.K_KP_4)
	// KP5 ist the key 5 on the key pad.
	KP5 = Key(sdl.K_KP_5)
	// KP6 ist the key 6 on the key pad.
	KP6 = Key(sdl.K_KP_6)
	// KP7 ist the key 7 on the key pad.
	KP7 = Key(sdl.K_KP_7)
	// KP8 ist the key 8 on the key pad.
	KP8 = Key(sdl.K_KP_8)
	// KP9 ist the key 9 on the key pad.
	KP9 = Key(sdl.K_KP_9)
	// KPDecimal ist the decimal key on the key pad.
	KPDecimal = Key(sdl.K_KP_DECIMAL)
	// KPDivide ist the divide key on the key pad.
	KPDivide = Key(sdl.K_KP_DIVIDE)
	// KPMuliply ist the multiply key on the key pad.
	KPMuliply = Key(sdl.K_KP_MULTIPLY)
	// KPSubtract ist the subtract key on the key pad.
	KPSubtract = Key(sdl.K_KP_MINUS)
	// KPAdd ist the add key on the key pad.
	KPAdd = Key(sdl.K_KP_PLUS)
	// LeftShift is the keft shift key.
	LeftShift = Key(sdl.K_LSHIFT)
	// LeftControl is the keft shift key.
	LeftControl = Key(sdl.K_LCTRL)
	// LeftAlt is the keft shift key.
	LeftAlt = Key(sdl.K_LALT)
	// // LeftSuper is the keft shift key.
	// LeftSuper = Key(sdl.K_)
	// RightShift is the keft shift key.
	RightShift = Key(sdl.K_RSHIFT)
	// RightControl is the keft shift key.
	RightControl = Key(sdl.K_RCTRL)
	// RightAlt is the keft shift key.
	RightAlt = Key(sdl.K_RALT)
	// // RightSuper is the keft shift key.
	// RightSuper = Key(glfw.KeyRightSuper)
	// Menu Key.
	Menu = Key(sdl.K_MENU)
)
