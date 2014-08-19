package tea

import "testing"

func TestEncrypt(t *testing.T) {
	plaintext := []uint32{0x01234567, 0x89abcdef}
	key := []uint32{0x00112233, 0x44556677, 0x8899aabb, 0xccddeeff}
	ciphertext := Encrypt(plaintext, key)

	if ciphertext[0] != 0x126c6b92 || ciphertext[1] != 0xc0653a3e {
		t.Errorf("ciphertext: %x, %x\n", ciphertext[0], ciphertext[1])
	}
}

func TestDecrypt(t *testing.T) {
	ciphertext := []uint32{0x126c6b92, 0xc0653a3e}
	key := []uint32{0x00112233, 0x44556677, 0x8899aabb, 0xccddeeff}

	plaintext := Decrypt(ciphertext, key)

	if plaintext[0] != 0x01234567 || plaintext[1] != 0x89abcdef {
		t.Errorf("plaintext: %x, %x\n", plaintext[0], plaintext[1])
	}
}
