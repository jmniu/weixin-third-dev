package wxopencrypt

const BLOCKSIZE = 32

func DecodeInPKCS7(buf_en []byte) []byte {
	tail := int(buf_en[len(buf_en)-1])

	if tail < 1 || tail > BLOCKSIZE {
		tail = 0
	}

	return buf_en[:len(buf_en)-tail]
}

func EncodeInPKCS7(buf_de []byte) []byte {
	tail := BLOCKSIZE - len(buf_de)%BLOCKSIZE

	for i := 0; i < tail; i++ {
		buf_de = append(buf_de, byte(tail))
	}

	return buf_de
}
