package tea

func Encrypt(v, k []uint32) [2]uint32 {
	var delta uint32 = 0x9e3779b9
	var sum uint32
	y, z := v[0], v[1]
	k0, k1, k2, k3 := k[0], k[1], k[2], k[3]

	for i := 0; i < 32; i++ {
		sum += delta
		y += ((z << 4) + k0) ^ (z + sum) ^ ((z >> 5) + k1)
		z += ((y << 4) + k2) ^ (y + sum) ^ ((y >> 5) + k3)
	}

	return [2]uint32{y, z}
}

func Decrypt(v, k []uint32) [2]uint32 {
	var delta uint32 = 0x9e3779b9
	var sum uint32 = (delta << 5) & 0xFFFFFFFF
	y, z := v[0], v[1]
	k0, k1, k2, k3 := k[0], k[1], k[2], k[3]

	for i := 0; i < 32; i++ {
		z -= ((y << 4) + k2) ^ (y + sum) ^ ((y >> 5) + k3)
		y -= ((z << 4) + k0) ^ (z + sum) ^ ((z >> 5) + k1)
		sum -= delta
	}

	return [2]uint32{y, z}
}
