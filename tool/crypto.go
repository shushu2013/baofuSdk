package tool

import (
	"bytes"
	"crypto"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/x509"
	"encoding/hex"
	"encoding/pem"
	"errors"
	"fmt"
	"math/big"
	"os"
	"sort"
	"strings"
	"time"

	"golang.org/x/crypto/pkcs12"
)

const SIGNATURE_SHA1_WITH_RSA_ALGORITHM = "SHA1withRSA"
const SIGNATURE_SHA256_WITH_RSA_ALGORITHM = "SHA256withRSA"

// 将字节数组转换为十六进制字符串
func byte2Hex(srcBytes []byte) string {
	return hex.EncodeToString(srcBytes)
}

// 将十六进制字符串转换为字节数组
func hex2Bytes(source string) ([]byte, error) {
	return hex.DecodeString(source)
}

func formatSerialNumber(serialNumber *big.Int) string {
	// 判断是否需要前导 0x00：
	// 如果最高位字节的最高位为 1，则需加 0x00 避免被解释为负数
	if serialNumber.BitLen() > 0 && serialNumber.Bit(0) == 1 {
		return fmt.Sprintf("00%x", serialNumber)
	}
	return fmt.Sprintf("%x", serialNumber)
}

// 从证书文件读取公钥的序列号
func GetPublicKeySNFromFile(pubCerPath string) (string, error) {
	// 读取证书文件
	certData, err := os.ReadFile(pubCerPath)
	if err != nil {
		return "", fmt.Errorf("公钥文件读取失败: %v", err)
	}

	// 解析PEM格式
	block, _ := pem.Decode(certData)
	if block == nil {
		return "", errors.New("无效的PEM格式")
	}

	// 解析证书
	cert, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		return "", fmt.Errorf("解析证书失败: %v", err)
	}

	// 获取证书序列号
	if cert.SerialNumber == nil {
		return "", errors.New("证书序列号为空")
	}

	// 将证书序列号转换为十六进制字符串（小写）
	return formatSerialNumber(cert.SerialNumber), nil
}

// 从证书文件读取公钥
func GetPublicKeyFromFile(pubCerPath string) (*rsa.PublicKey, error) {
	// 读取证书文件
	certData, err := os.ReadFile(pubCerPath)
	if err != nil {
		return nil, fmt.Errorf("公钥文件读取失败: %v", err)
	}

	return getPublicKeyByText(string(certData))
}

// 根据公钥文本读取公钥
func getPublicKeyByText(pubKeyText string) (*rsa.PublicKey, error) {
	// 清理PEM格式
	pubKeyText = strings.TrimSpace(pubKeyText)

	// 尝试解析PEM格式
	block, _ := pem.Decode([]byte(pubKeyText))
	if block == nil {
		return nil, errors.New("无效的PEM格式")
	}

	// 如果是证书格式
	if block.Type == "CERTIFICATE" {
		cert, err := x509.ParseCertificate(block.Bytes)
		if err != nil {
			return nil, fmt.Errorf("解析证书失败: %v", err)
		}
		publicKey, ok := cert.PublicKey.(*rsa.PublicKey)
		if !ok {
			return nil, errors.New("不是RSA公钥")
		}
		return publicKey, nil
	}

	// 如果是公钥格式
	if block.Type == "PUBLIC KEY" || block.Type == "RSA PUBLIC KEY" {
		publicKey, err := x509.ParsePKIXPublicKey(block.Bytes)
		if err != nil {
			// 尝试PKCS1格式
			publicKey, err := x509.ParsePKCS1PublicKey(block.Bytes)
			if err != nil {
				return nil, fmt.Errorf("解析公钥失败: %v", err)
			}
			return publicKey, nil
		}
		rsaPublicKey, ok := publicKey.(*rsa.PublicKey)
		if !ok {
			return nil, errors.New("不是RSA公钥")
		}
		return rsaPublicKey, nil
	}

	return nil, errors.New("不支持的PEM类型: " + block.Type)
}

// 从PFX文件读取私钥的序列号
func GetPrivateKeySNFromFile(pfxPath, priKeyPass string) (string, error) {
	// 读取PFX文件
	pfxData, err := os.ReadFile(pfxPath)
	if err != nil {
		return "", fmt.Errorf("私钥文件读取失败: %v", err)
	}

	// 解析PFX文件，获取证书和私钥
	privateKey, certificate, err := pkcs12.Decode(pfxData, priKeyPass)
	if err != nil {
		return "", fmt.Errorf("解析PFX文件失败: %v", err)
	}

	// 验证私钥类型
	_, ok := privateKey.(*rsa.PrivateKey)
	if !ok {
		return "", errors.New("文件不是PFX格式的RSA私钥")
	}

	// 从证书中获取序列号
	if certificate == nil {
		return "", errors.New("PFX文件中未找到证书")
	}

	// 将证书序列号转换为十六进制字符串（大端序）
	serialNumber := certificate.SerialNumber
	if serialNumber == nil {
		return "", errors.New("证书序列号为空")
	}

	return formatSerialNumber(serialNumber), nil
}

// 从PFX文件读取私钥
func GetPrivateKeyFromFile(pfxPath, priKeyPass string) (*rsa.PrivateKey, error) {
	pfxData, err := os.ReadFile(pfxPath)
	if err != nil {
		return nil, fmt.Errorf("私钥文件读取失败: %v", err)
	}

	return GetPrivateKeyByStream(pfxData, priKeyPass)
}

// 根据PFX字节流读取私钥
func GetPrivateKeyByStream(pfxBytes []byte, priKeyPass string) (*rsa.PrivateKey, error) {
	// 使用 pkcs12 库解析私钥
	privateKey, _, err := pkcs12.Decode(pfxBytes, priKeyPass)

	if err != nil {
		return nil, fmt.Errorf("解析私钥失败: %v", err)
	}

	rsaPrivateKey, ok := privateKey.(*rsa.PrivateKey)
	if !ok {
		return nil, errors.New("文件不是 PFX 格式的私钥")
	}

	return rsaPrivateKey, nil
}

// 验签
func verifySignature(pubCerPath, encryptStr, signature string, algorithm string) (bool, error) {

	publicKey, err := GetPublicKeyFromFile(pubCerPath)
	if err != nil {
		return false, err
	}

	return Verify(encryptStr, publicKey, signature, algorithm)
}

// 签名
func encryptByRSA(encryptStr, pfxPath, priKeyPass string, algorithm string) (string, error) {
	privateKey, err := GetPrivateKeyFromFile(pfxPath, priKeyPass)
	if err != nil {
		return "", err
	}

	return Sign(encryptStr, privateKey, algorithm)
}

// 校验数字签名
func Verify(data string, publicKey *rsa.PublicKey, sign string, algorithm string) (bool, error) {
	// 转换签名
	signature, err := hex2Bytes(sign)
	if err != nil {
		return false, fmt.Errorf("签名格式错误: %v", err)
	}

	var hashed []byte
	var hashMethod crypto.Hash

	if algorithm == SIGNATURE_SHA1_WITH_RSA_ALGORITHM {
		hashMethod = crypto.SHA1

		// 计算哈希
		hash := sha1.New()
		hash.Write([]byte(data))
		hashed = hash.Sum(nil)
	} else if algorithm == SIGNATURE_SHA256_WITH_RSA_ALGORITHM {
		hashMethod = crypto.SHA256

		// 计算哈希
		hash := sha256.Sum256([]byte(data))
		hashed = hash[:]
	} else {
		return false, fmt.Errorf("不支持的签名算法: %s", algorithm)
	}

	// 验证签名
	err = rsa.VerifyPKCS1v15(publicKey, hashMethod, hashed, signature)
	if err != nil {
		return false, nil // 验证失败，但不返回错误
	}

	return true, nil
}

// 用私钥对信息生成数字签名
func Sign(data string, privateKey *rsa.PrivateKey, algorithm string) (string, error) {

	var hashed []byte
	var hashMethod crypto.Hash

	if algorithm == SIGNATURE_SHA1_WITH_RSA_ALGORITHM {
		hashMethod = crypto.SHA1

		// 计算哈希
		hash := sha1.New()
		hash.Write([]byte(data))
		hashed = hash.Sum(nil)
	} else if algorithm == SIGNATURE_SHA256_WITH_RSA_ALGORITHM {
		hashMethod = crypto.SHA256

		// 计算哈希
		hash := sha256.Sum256([]byte(data))
		hashed = hash[:]
	} else {
		return "", fmt.Errorf("不支持的签名算法: %s", algorithm)
	}

	// 生成签名
	signature, err := rsa.SignPKCS1v15(rand.Reader, privateKey, hashMethod, hashed)
	if err != nil {
		return "", fmt.Errorf("签名生成失败: %v", err)
	}

	return byte2Hex(signature), nil
}

// 用公钥加密数据
func EncryptByPublicKey(data string, publicKey *rsa.PublicKey) (string, error) {
	// 加密数据
	encryptedData, err := rsa.EncryptPKCS1v15(rand.Reader, publicKey, []byte(data))
	if err != nil {
		return "", fmt.Errorf("加密数据失败: %v", err)
	}

	return byte2Hex(encryptedData), nil
}

// 用私钥解密数据
func DecryptByPrivateKey(encryptedData string, privateKey *rsa.PrivateKey) (string, error) {
	// 将加密数据转换为字节数组
	encryptedBytes, err := hex2Bytes(encryptedData)
	if err != nil {
		return "", fmt.Errorf("加密数据格式错误: %v", err)
	}

	// 解密数据
	decryptedData, err := rsa.DecryptPKCS1v15(rand.Reader, privateKey, encryptedBytes)
	if err != nil {
		return "", fmt.Errorf("解密数据失败: %v", err)
	}

	return string(decryptedData), nil
}

// RsaEncryptByPrivateKey 使用私钥进行 RSA 加密，分段处理
// src: 源数据字符串
// privateKey: RSA 私钥对象
// 返回处理后的十六进制字符串或错误
func RsaEncryptByPrivateKey(src string, privateKey *rsa.PrivateKey) (string, error) {
	srcData := []byte(src)
	// RSA 密钥长度 (如 2048bit = 256 byte)
	keySize := privateKey.N.BitLen() / 8
	// PKCS1v15填充长度为11字节，分段明文必须小于 keySize - 11
	maxBlockSize := keySize - 11

	result := make([]byte, 0, len(srcData))
	for len(srcData) > 0 {
		var chunk []byte
		if len(srcData) > maxBlockSize {
			chunk = srcData[:maxBlockSize]
			srcData = srcData[maxBlockSize:]
		} else {
			chunk = srcData
			srcData = nil
		}

		// 使用私钥加密 (私钥加密需注意标准，这里为了兼容使用SignPKCS1v15作为替代方案)
		// 实际上，RSA私钥加密等同于签名，但如果专门指"加密"，通常指SignPKCS1v15
		// 注意: 使用 rsa.EncryptPKCS1v15 传入私钥是不对的，需正确区分加签和加密

		encryptedChunk, err := rsa.SignPKCS1v15(rand.Reader, privateKey, 0, chunk) // 这里的0表示未指定哈希
		if err != nil {
			return "", err
		}
		result = append(result, encryptedChunk...)
	}

	return byte2Hex(result), nil
}

// 用公钥解密数据（支持分段解密）
func DecryptByPublicKey(encryptedData string, publicKey *rsa.PublicKey) (string, error) {
	// 将加密数据转换为字节数组
	ciphertext, err := hex2Bytes(encryptedData)
	if err != nil {
		return "", fmt.Errorf("加密数据格式错误: %v", err)
	}

	keySize := publicKey.N.BitLen() / 8
	var decryptedBytes []byte

	// 分段解密，每段长度为keySize
	for i := 0; i < len(ciphertext); i += keySize {
		end := i + keySize
		if end > len(ciphertext) {
			end = len(ciphertext)
		}

		// 解密当前段
		decryptedBlock, err := rsaDecryptWithPublicKey(ciphertext[i:end], publicKey)
		if err != nil {
			return "", fmt.Errorf("解密数据块失败 (offset=%d): %v", i, err)
		}

		decryptedBytes = append(decryptedBytes, decryptedBlock...)
	}

	return string(decryptedBytes), nil
}

// rsaDecryptWithPublicKey 使用公钥解密的内部实现
func rsaDecryptWithPublicKey(ciphertext []byte, publicKey *rsa.PublicKey) ([]byte, error) {
	// 将密文转换为大数
	c := new(big.Int).SetBytes(ciphertext)

	// RSA解密：m = c^e mod n
	m := new(big.Int).Exp(c, big.NewInt(int64(publicKey.E)), publicKey.N)

	// 转换为字节数组，需要补齐到密钥长度
	keySize := publicKey.N.BitLen() / 8
	plaintext := m.Bytes()

	// 如果结果长度小于密钥长度，需要在前面补0
	if len(plaintext) < keySize {
		padded := make([]byte, keySize)
		copy(padded[keySize-len(plaintext):], plaintext)
		plaintext = padded
	}

	// PKCS#1 v1.5 有两种格式：
	// Type 1 (签名): 0x00 0x01 [0xFF填充] 0x00 [数据]
	// Type 2 (加密): 0x00 0x02 [随机非零字节] 0x00 [数据]

	if len(plaintext) < 11 {
		return nil, fmt.Errorf("无效的PKCS#1 v1.5 padding: 长度=%d", len(plaintext))
	}

	if plaintext[0] != 0x00 {
		return nil, fmt.Errorf("无效的PKCS#1 v1.5 padding: 首字节=0x%02x (预期0x00)", plaintext[0])
	}

	var dataStart int

	if plaintext[1] == 0x01 {
		// Type 1: 签名格式 - 查找 0x00 分隔符（前面是 0xFF 填充）
		i := 2
		for i < len(plaintext) && plaintext[i] == 0xFF {
			i++
		}
		if i >= len(plaintext) || plaintext[i] != 0x00 {
			return nil, fmt.Errorf("无效的PKCS#1 v1.5 Type 1 padding: 未找到分隔符")
		}
		dataStart = i + 1
	} else if plaintext[1] == 0x02 {
		// Type 2: 加密格式 - 查找 0x00 分隔符（前面是随机非零字节）
		i := 2
		for i < len(plaintext) && plaintext[i] != 0x00 {
			i++
		}
		if i >= len(plaintext) {
			return nil, fmt.Errorf("无效的PKCS#1 v1.5 Type 2 padding: 未找到分隔符")
		}
		dataStart = i + 1
	} else {
		return nil, fmt.Errorf("无效的PKCS#1 v1.5 padding: 次字节=0x%02x (预期0x01或0x02)", plaintext[1])
	}

	// 返回实际数据
	return plaintext[dataStart:], nil
}

// aes加密-128位
func AesEncrypt(data, key string) (string, error) {
	// 检查密钥长度，AES-128需要16字节密钥
	keyBytes := []byte(key)
	if len(keyBytes) != 16 {
		return "", fmt.Errorf("AES-128加密需要16字节密钥，当前密钥长度为%d字节", len(keyBytes))
	}

	// 将数据转换为字节数组
	dataBytes := []byte(data)

	// 创建AES加密块
	block, err := aes.NewCipher(keyBytes)
	if err != nil {
		return "", fmt.Errorf("创建AES加密块失败: %v", err)
	}

	// 填充数据到块大小的倍数
	blockSize := block.BlockSize()
	dataBytes = pkcs7Padding(dataBytes, blockSize)

	// 创建加密模式（CBC模式）
	cipherText := make([]byte, len(dataBytes))
	iv := keyBytes[:blockSize] // 使用密钥的前16字节作为IV
	mode := cipher.NewCBCEncrypter(block, iv)

	// 执行加密
	mode.CryptBlocks(cipherText, dataBytes)

	// 返回十六进制字符串
	return byte2Hex(cipherText), nil
}

// PKCS7填充
func pkcs7Padding(data []byte, blockSize int) []byte {
	dataLen := len(data)
	if dataLen%blockSize == 0 {
		return data
	}

	padding := blockSize - dataLen%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(data, padText...)
}

// aes解密-128位
func AesDecrypt(encryptedData, key string) (string, error) {
	// 检查密钥长度，AES-128需要16字节密钥
	keyBytes := []byte(key)
	if len(keyBytes) != 16 {
		return "", fmt.Errorf("AES-128解密需要16字节密钥，当前密钥长度为%d字节", len(keyBytes))
	}

	// 将加密数据从十六进制字符串转换为字节数组
	encryptedBytes, err := hex2Bytes(encryptedData)
	if err != nil {
		return "", fmt.Errorf("加密数据格式错误: %v", err)
	}

	// 创建AES解密块
	block, err := aes.NewCipher(keyBytes)
	if err != nil {
		return "", fmt.Errorf("创建AES解密块失败: %v", err)
	}

	// 检查加密数据长度是否为块大小的倍数
	blockSize := block.BlockSize()
	if len(encryptedBytes)%blockSize != 0 {
		return "", errors.New("加密数据长度不是块大小的倍数")
	}

	// 创建解密模式（CBC模式）
	decryptedData := make([]byte, len(encryptedBytes))
	iv := keyBytes[:blockSize] // 使用密钥的前16字节作为IV
	mode := cipher.NewCBCDecrypter(block, iv)

	// 执行解密
	mode.CryptBlocks(decryptedData, encryptedBytes)

	// 返回解密后的字符串
	return string(decryptedData), nil
}

// sha1计算后进行16进制转换
func Sha1X16(data string) string {
	hash := sha1.Sum([]byte(data))
	return byte2Hex(hash[:])
}

// coverMap2String 将 map[string]string 按键排序后拼接成 "k1=v1&k2=v2" 形式的字符串
// 空值（仅含空白字符视为空）的键将被忽略
func CoverMap2String(data map[string]string) string {
	if len(data) == 0 {
		return ""
	}

	// 取出所有 key 并排序
	keys := make([]string, 0, len(data))
	for k := range data {
		keys = append(keys, k)
	}
	// 按字典序排序
	sort.Strings(keys)

	var sb strings.Builder
	for _, k := range keys {
		v := data[k]
		// 去掉首尾空白后判断是否为空
		trimmed := strings.TrimSpace(v)
		if trimmed == "" {
			continue
		}
		if sb.Len() > 0 {
			sb.WriteString("&")
		}
		sb.WriteString(fmt.Sprintf("%s=%s", k, trimmed))
	}
	return sb.String()
}

// 创建AES密钥
func CreateAeskey(i int) string {
	// 生成随机数+时间戳+随机数的组合
	randKey := fmt.Sprintf("%d%d%d", RandomInt(4), GetTimeMilliseconds(time.Now()), RandomInt(4))

	// 计算MD5值
	md5Key := MD5(randKey)

	// 处理参数i的范围
	if i < 1 {
		i = 1
	}

	// 根据i的值返回相应长度的密钥
	if i <= 32 {
		return strings.ToUpper(md5Key[:i])
	} else {
		return strings.ToUpper(md5Key)
	}
}

// 获取请求流水号
func GetTransSerialNo() string {
	currentTime := time.Now()
	strTime := currentTime.Format("20060102150405.000")
	// 移除毫秒前的小数点，使其成为连续的数字
	// 为了降低重复的概率，后面再加10位随机数
	return strings.Replace(strTime, ".", "", 1) + Int64ToString(RandomInt(10))
}

func GetRequestNo() string {
	return GetTransSerialNo()
}

func GetMsgId() string {
	return GetTransSerialNo()
}
