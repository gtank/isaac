/*
------------------------------------------------------------------------------
isaac.go: an implementation of Bob Jenkins' random number generator ISAAC based on 'readable.c'
* 18 Aug 2014 -- direct port of readable.c to Go
------------------------------------------------------------------------------
*/

package isaac

/* external results */
var randrsl [256]uint32
var randcnt uint32

/* internal state */
var mm [256]uint32
var aa, bb, cc uint32 // zero by default

func Isaac() {
	var x, y uint32

	cc = cc + 1  /* cc just gets incremented once per 256 results */
	bb = bb + cc /* then combined with bb */

	for i := 0; i < 256; i++ {
		x = mm[i]

		switch i % 4 {
		case 0:
			aa = aa ^ (aa << 13)
			break
		case 1:
			aa = aa ^ (aa >> 6)
			break
		case 2:
			aa = aa ^ (aa << 2)
			break
		case 3:
			aa = aa ^ (aa >> 16)
			break
		}

		aa = mm[(i+128)%256] + aa
		y = mm[(x>>2)%256] + aa + bb
		bb = mm[(y>>10)%256] + x

		mm[i] = y
		randrsl[i] = bb
	}
}

/* if (flag!=0), then use the contents of randrsl[] to initialize mm[]. */
func mix(a, b, c, d, e, f, g, h uint32) (uint32, uint32, uint32, uint32, uint32, uint32, uint32, uint32) {
	a ^= b << 11
	d += a
	b += c
	b ^= c >> 2
	e += b
	c += d
	c ^= d << 8
	f += c
	d += e
	d ^= e >> 16
	g += d
	e += f
	e ^= f << 10
	h += e
	f += g
	f ^= g >> 4
	a += f
	g += h
	g ^= h << 8
	b += g
	h += a
	h ^= a >> 9
	c += h
	a += b
	return a, b, c, d, e, f, g, h
}

func Randinit(flag bool) {
	var a, b, c, d, e, f, g, h uint32
	aa, bb, cc = 0, 0, 0

	a, b, c, d, e, f, g, h = 0x9e3779b9, 0x9e3779b9, 0x9e3779b9, 0x9e3779b9, 0x9e3779b9, 0x9e3779b9, 0x9e3779b9, 0x9e3779b9

	for i := 0; i < 4; i++ {
		a, b, c, d, e, f, g, h = mix(a, b, c, d, e, f, g, h)
	}

	for i := 0; i < 256; i += 8 { /* fill mm[] with messy stuff */
		if flag { /* use all the information in the seed */
			a += randrsl[i]
			b += randrsl[i+1]
			c += randrsl[i+2]
			d += randrsl[i+3]
			e += randrsl[i+4]
			f += randrsl[i+5]
			g += randrsl[i+6]
			h += randrsl[i+7]
		}
		a, b, c, d, e, f, g, h = mix(a, b, c, d, e, f, g, h)
		mm[i] = a
		mm[i+1] = b
		mm[i+2] = c
		mm[i+3] = d
		mm[i+4] = e
		mm[i+5] = f
		mm[i+6] = g
		mm[i+7] = h
	}

	if flag { /* do a second pass to make all of the seed affect all of mm */
		for i := 0; i < 256; i += 8 {
			a += mm[i]
			b += mm[i+1]
			c += mm[i+2]
			d += mm[i+3]
			e += mm[i+4]
			f += mm[i+5]
			g += mm[i+6]
			h += mm[i+7]
			a, b, c, d, e, f, g, h = mix(a, b, c, d, e, f, g, h)
			mm[i] = a
			mm[i+1] = b
			mm[i+2] = c
			mm[i+3] = d
			mm[i+4] = e
			mm[i+5] = f
			mm[i+6] = g
			mm[i+7] = h
		}
	}

	Isaac()
	randcnt = 256
}

func Randcnt() uint32 {
	return randcnt
}

func Randrsl() [256]uint32 {
	return randrsl
}
