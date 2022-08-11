// Package password
// @Description: AES加密解密包
package password

import (
	"crypto/aes"
	"crypto/rand"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"crypto/subtle"
	"encoding/base64"
	"errors"
	"fmt"
	"golang.org/x/crypto/pbkdf2"
	"golang.org/x/crypto/scrypt"
	"strconv"
	"strings"
)

/**
 * @Author: h.li
 * @Date: 2022/04/29
 * @Description: 数据库加密方法，同步网阅系统的加密方法
 * @param src 加密前密码（byte）
 * @param key 密钥
 * @return []byte
 * @return error
 */
func AesEncryptECB(src []byte, key []byte) ([]byte, error) {
	key, err := AesSha1prng(key, 128)
	if err != nil {
		return nil, err
	}
	cipher, _ := aes.NewCipher(generateKey(key))
	length := (len(src) + aes.BlockSize) / aes.BlockSize
	plain := make([]byte, length*aes.BlockSize)
	copy(plain, src)
	pad := byte(len(plain) - len(src))
	for i := len(src); i < len(plain); i++ {
		plain[i] = pad
	}
	encrypted := make([]byte, len(plain))
	// 分组分块加密
	for bs, be := 0, cipher.BlockSize(); bs <= len(src); bs, be = bs+cipher.BlockSize(), be+cipher.BlockSize() {
		cipher.Encrypt(encrypted[bs:be], plain[bs:be])
	}

	return encrypted, nil
}

/**
 * @Author: h.li
 * @Date: 2022/04/29
 * @Description: 解密算法，同步网阅
 * @param encrypted 加密后的密码（byte）
 * @param key 密钥
 * @return []byte
 * @return error
 */
func AesDecryptECB(encrypted []byte, key []byte) ([]byte, error) {
	key, err := AesSha1prng(key, 128)
	if err != nil {
		return nil, err
	}

	cipher, _ := aes.NewCipher(generateKey(key))
	decrypted := make([]byte, len(encrypted))

	for bs, be := 0, cipher.BlockSize(); bs < len(encrypted); bs, be = bs+cipher.BlockSize(), be+cipher.BlockSize() {
		cipher.Decrypt(decrypted[bs:be], encrypted[bs:be])
	}

	trim := 0
	if len(decrypted) > 0 {
		trim = len(decrypted) - int(decrypted[len(decrypted)-1])
	}
	return decrypted[:trim], nil
}

// 模拟 java SHA1PRNG 处理
func AesSha1prng(keyBytes []byte, encryptLength int) ([]byte, error) {
	hashs := Sha1(Sha1(keyBytes))
	maxLen := len(hashs)
	realLen := encryptLength / 8
	if realLen > maxLen {
		return nil, errors.New("invalid length")
	}

	return hashs[0:realLen], nil
}

func Sha1(data []byte) []byte {
	h := sha1.New()
	h.Write(data)
	return h.Sum(nil)
}

func generateKey(key []byte) (genKey []byte) {
	genKey = make([]byte, 16)
	copy(genKey, key)
	for i := 16; i < len(key); {
		for j := 0; j < 16 && i < len(key); j, i = j+1, i+1 {
			genKey[j] ^= key[i]
		}
	}
	return genKey
}

/* The PBKDF2_* and SCRYPT_* constants may be changed without breaking existing stored hashes. */
const (
	// PBKDF2_HASH_ALGORITHM can be set to sha1, sha224, sha256, sha384 or sha512 as the underlying hashing mechanism to be used by the PBKDF2 function
	PBKDF2_HASH_ALGORITHM string = "sha512"
	// PBKDF2_ITERATIONS sets the amount of iterations used by the PBKDF2 hashing algorithm
	PBKDF2_ITERATIONS int = 15000
	// SCRYPT_N is a CPU/memory cost parameter, which must be a power of two greater than 1
	SCRYPT_N int = 32768
	// SCRYPT_R is the block size parameter
	SCRYPT_R int = 8
	// SCRYPT_P is the parallelization parameter, a positive integer less than or equal to ((2^32-1) * 32) / (128 * r)
	SCRYPT_P int = 1

	// SALT_BYTES sets the amount of bytes for the salt used in the PBKDF2 / scrypt hashing algorithm
	SALT_BYTES int = 64
	// HASH_BYTES sets the amount of bytes for the hash output from the PBKDF2 / scrypt hashing algorithm
	HASH_BYTES int = 64
)

/* altering the HASH_* constants breaks existing stored hashes */
const (
	// HASH_SECTIONS identifies the expected amount of parameters encoded in a hash generated and/or tested in this package
	HASH_SECTIONS int = 4
	// HASH_ALGORITHM_INDEX identifies the position of the hash algorithm identifier in a hash generated and/or tested in this package
	HASH_ALGORITHM_INDEX int = 0
	// HASH_ITERATION_INDEX identifies the position of the iteration count used by PBKDF2 in a hash generated and/or tested in this package
	HASH_ITERATION_INDEX int = 1
	// HASH_SALT_INDEX identifies the position of the used salt in a hash generated and/or tested in this package
	HASH_SALT_INDEX int = 2
	// HASH_PBKDF2_INDEX identifies the position of the actual password hash in a hash generated and/or tested in this package
	HASH_PBKDF2_INDEX int = 3
	// HASH_SCRYPT_R_INDEX identifies the position of the scrypt block size parameter in a hash generated and/or tested in this package
	HASH_SCRYPT_R_INDEX int = 4
	// HASH_SCRYPT_R_INDEX identifies the position of the scrypt parallelization parameter in a hash generated and/or tested in this package
	HASH_SCRYPT_P_INDEX int = 5
)

// CreateHash creates a salted cryptographic hash with key stretching (PBKDF2), suitable for storage and usage in password authentication mechanisms.
func Encrypt(password string) (string, error) {
	salt := make([]byte, SALT_BYTES)
	_, err := rand.Read(salt)
	if err != nil {
		return "", err
	}
	var hash []byte
	switch PBKDF2_HASH_ALGORITHM {
	default:
		return "", errors.New("invalid hash algorithm selected")
	case "sha1":
		hash = pbkdf2.Key([]byte(password), salt, PBKDF2_ITERATIONS, HASH_BYTES, sha1.New)
	case "sha224":
		hash = pbkdf2.Key([]byte(password), salt, PBKDF2_ITERATIONS, HASH_BYTES, sha256.New224)
	case "sha256":
		hash = pbkdf2.Key([]byte(password), salt, PBKDF2_ITERATIONS, HASH_BYTES, sha256.New)
	case "sha384":
		hash = pbkdf2.Key([]byte(password), salt, PBKDF2_ITERATIONS, HASH_BYTES, sha512.New384)
	case "sha512":
		hash = pbkdf2.Key([]byte(password), salt, PBKDF2_ITERATIONS, HASH_BYTES, sha512.New)
	case "scrypt":
		hash, err = scrypt.Key([]byte(password), salt, SCRYPT_N, SCRYPT_R, SCRYPT_P, HASH_BYTES)
		if err != nil {
			return "", err
		}
		/* format: algorithm:cpu/mem cost:salt:hash:R(blocksize):P(parallelization) */
		return fmt.Sprintf(
			"%s:%d:%s:%s:%d:%d", PBKDF2_HASH_ALGORITHM, SCRYPT_N,
			base64.StdEncoding.EncodeToString(salt), base64.StdEncoding.EncodeToString(hash),
			SCRYPT_R, SCRYPT_P,
		), err
	}

	/* format: algorithm:iterations:salt:hash */
	return fmt.Sprintf(
		"%s:%d:%s:%s", PBKDF2_HASH_ALGORITHM, PBKDF2_ITERATIONS,
		base64.StdEncoding.EncodeToString(salt), base64.StdEncoding.EncodeToString(hash),
	), err
}

// ValidatePassword hashes a password according to the setup found in the correct hash string and does a constant time compare on the correct hash and calculated hash.
func Validate(password string, correctHash string) bool {
	params := strings.Split(correctHash, ":")
	if len(params) < HASH_SECTIONS {
		return false
	}
	it, err := strconv.Atoi(params[HASH_ITERATION_INDEX])
	if err != nil {
		return false
	}
	salt, err := base64.StdEncoding.DecodeString(params[HASH_SALT_INDEX])
	if err != nil {
		return false
	}
	hash, err := base64.StdEncoding.DecodeString(params[HASH_PBKDF2_INDEX])
	if err != nil {
		return false
	}

	var testHash []byte
	switch params[HASH_ALGORITHM_INDEX] {
	default:
		return false
	case "sha1":
		testHash = pbkdf2.Key([]byte(password), salt, it, len(hash), sha1.New)
	case "sha224":
		testHash = pbkdf2.Key([]byte(password), salt, it, len(hash), sha256.New224)
	case "sha256":
		testHash = pbkdf2.Key([]byte(password), salt, it, len(hash), sha256.New)
	case "sha384":
		testHash = pbkdf2.Key([]byte(password), salt, it, len(hash), sha512.New384)
	case "sha512":
		testHash = pbkdf2.Key([]byte(password), salt, it, len(hash), sha512.New)
	case "scrypt":
		r, err := strconv.Atoi(params[HASH_SCRYPT_R_INDEX])
		if err != nil {
			return false
		}
		p, err := strconv.Atoi(params[HASH_SCRYPT_P_INDEX])
		if err != nil {
			return false
		}
		testHash, err = scrypt.Key([]byte(password), salt, it, r, p, len(hash))
	}
	return subtle.ConstantTimeCompare(hash, testHash) == 1
}
