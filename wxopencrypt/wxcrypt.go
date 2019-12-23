package wxopencrypt

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha1"
	"encoding/base64"
	"encoding/binary"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"sort"
)

type WXBizMsgCrypt struct {
	m_sToken          string
	m_sEncodingAESKey string
	m_sAppid          string
	m_sKey            []byte
	m_sIv             []byte
}

func NewWXBizMsgCrypt() *WXBizMsgCrypt {
	wbmc := &WXBizMsgCrypt{}
	return wbmc
}

// var GWXBizMsgCrypt = newWXBizMsgCrypt()

func (this *WXBizMsgCrypt) Init(sToken string, sEncodingAESKey string, sAppid string) {
	this.m_sToken = sToken
	this.m_sEncodingAESKey = sEncodingAESKey
	this.m_sAppid = sAppid

	key, err := base64.StdEncoding.DecodeString(this.m_sEncodingAESKey + "=")
	if err != nil || len(key) != kAesKeySize {
		panic(1)
	}
	this.m_sKey = nil
	this.m_sKey = append(this.m_sKey, key...)
	this.m_sIv = this.m_sKey[:kAesIVSize]
}

func (this *WXBizMsgCrypt) DecryptMsg(sMsgSignature string, sTimeStamp string, sNonce string, sPostData string) (ret int, sMsg string) {
	//1.validate xml format
	ret1, sEncryptMsg := this.GetEncryptMsg(sPostData)
	if 0 != ret1 {
		ret = WXBizMsgCrypt_ParseXml_Error
		return
	}

	//2.validate signature
	if 0 != this.ValidateSignature(sMsgSignature, sTimeStamp, sNonce, sEncryptMsg) {
		ret = WXBizMsgCrypt_ValidateSignature_Error
		return
	}

	//3.decode base64
	sAesData, err := base64.StdEncoding.DecodeString(sEncryptMsg)
	if err != nil {
		ret = WXBizMsgCrypt_DecodeBase64_Error
		return
	}

	//4.decode aes
	c, err := aes.NewCipher(this.m_sKey)
	if err != nil {
		ret = WXBizMsgCrypt_IllegalAesKey
		return
	}

	cbc := cipher.NewCBCDecrypter(c, this.m_sIv)
	cbc.CryptBlocks(sAesData, sAesData)

	// fmt.Println(string(sAesData))
	sNoEncryptData := DecodeInPKCS7(sAesData)
	// fmt.Println(string(sNoEncryptData))

	// 5. remove kRandEncryptStrLen str
	if len(sNoEncryptData) <= (kRandEncryptStrLen + kMsgLen) {
		ret = WXBizMsgCrypt_IllegalBuffer
		return
	}

	buf := bytes.NewBuffer(sNoEncryptData[kRandEncryptStrLen : kRandEncryptStrLen+kMsgLen])
	var iMsgLen int32
	binary.Read(buf, binary.BigEndian, &iMsgLen)

	if len(sNoEncryptData) <= (kRandEncryptStrLen + kMsgLen + int(iMsgLen)) {
		ret = WXBizMsgCrypt_IllegalBuffer
		return
	}

	//6. validate appid
	sAppid := string(sNoEncryptData[kRandEncryptStrLen+kMsgLen+iMsgLen:])
	if sAppid != this.m_sAppid {
		ret = WXBizMsgCrypt_ValidateAppid_Error
		return
	}
	sMsg = string(sNoEncryptData[kRandEncryptStrLen+kMsgLen : kRandEncryptStrLen+kMsgLen+iMsgLen])
	ret = WXBizMsgCrypt_OK
	return
}

func (this *WXBizMsgCrypt) EncryptMsg(sReplyMsg string, sTimeStamp string, sNonce string) (ret int, sEncryptMsg string) {
	if 0 == len(sReplyMsg) {
		ret = WXBizMsgCrypt_ParseXml_Error
		return
	}

	//1.add rand str ,len, appid
	// var sNeedEncrypt = this.GenNeedEncryptData(sReplyMsg)
	// sNeedEncrypt += ""
	buf := new(bytes.Buffer)
	if err := binary.Write(buf, binary.BigEndian, int32(len(sReplyMsg))); err != nil {
		ret = -1
		return
	}
	iMsgLen := buf.Bytes()
	randBytes := make([]byte, kRandEncryptStrLen)
	if _, err := io.ReadFull(rand.Reader, randBytes); err != nil {
		ret = -1
		return
	}
	sNeedEncrypt := bytes.Join([][]byte{randBytes, iMsgLen, []byte(sReplyMsg), []byte(this.m_sAppid)}, nil)

	//2. AES Encrypt
	// var sAesData string
	// var sAesKey string
	// if 0 != GenAesKeyFromEncodingKey(this.m_sEncodingAESKey, sAesKey) {
	// 	ret = WXBizMsgCrypt_IllegalAesKey
	// 	return
	// }
	// if 0 != AES_CBCEncrypt(sNeedEncrypt, sAesKey, &sAesData) {
	// 	ret = WXBizMsgCrypt_EncryptAES_Error
	// 	return
	// }
	sAesData := EncodeInPKCS7(sNeedEncrypt)
	c, err := aes.NewCipher(this.m_sKey)
	if err != nil {
		ret = WXBizMsgCrypt_EncryptAES_Error
		return
	}
	cbc := cipher.NewCBCEncrypter(c, this.m_sIv)
	cbc.CryptBlocks(sAesData, sAesData)

	//3. base64Encode
	sBase64Data := base64.StdEncoding.EncodeToString(sAesData)

	//4. compute signature
	ret4, sSignature := this.ComputeSignature(this.m_sToken, sTimeStamp, sNonce, sBase64Data)
	if 0 != ret4 {
		ret = WXBizMsgCrypt_ComputeSignature_Error
		return
	}

	//5. Gen xml
	ret, sEncryptMsg = this.GenReturnXml(sBase64Data, sSignature, sTimeStamp, sNonce)
	if 0 != ret {
		ret = WXBizMsgCrypt_GenReturnXml_Error
		return
	}
	ret = WXBizMsgCrypt_OK
	return
}

func (this *WXBizMsgCrypt) ComputeSignature(sToken string, sTimeStamp string, sNonce string, sMessage string) (ret int, sSignature string) {
	if 0 == len(sToken) || 0 == len(sNonce) || 0 == len(sMessage) || 0 == len(sTimeStamp) {
		ret = -1
		return
	}

	//sort
	var vecStr []string
	vecStr = append(vecStr, sToken)
	vecStr = append(vecStr, sTimeStamp)
	vecStr = append(vecStr, sNonce)
	vecStr = append(vecStr, sMessage)
	// std::sort( vecStr.begin(), vecStr.end() );
	sort.Strings(vecStr)
	sStr := vecStr[0] + vecStr[1] + vecStr[2] + vecStr[3]

	//compute
	h := sha1.New()
	h.Write([]byte(sStr))
	output := h.Sum(nil)

	// to hex
	for i := 0; i < len(output); i++ {
		// fmt.Sprintln(,)
		hexStr := fmt.Sprintf("%02x", 0xff&output[i])
		sSignature = sSignature + hexStr
	}
	return
}

func (this *WXBizMsgCrypt) ValidateSignature(sMsgSignature string, sTimeStamp string, sNonce string, sEncryptMsg string) int {
	ret, sSignature := this.ComputeSignature(this.m_sToken, sTimeStamp, sNonce, sEncryptMsg)
	if 0 != ret {
		return -1
	}

	if sMsgSignature != sSignature {
		return -1
	}

	return 0
}

func (this *WXBizMsgCrypt) GetEncryptMsg(sPostData string) (ret int, sEncryptMsg string) {
	var vt VerifyTicketEncrypt
	xml.Unmarshal([]byte(sPostData), &vt)

	if len(vt.AppId) <= 0 || len(vt.Encrypt) <= 0 {
		ret = -1
		return
	}

	sEncryptMsg = vt.Encrypt
	return
}

/*
void WXBizMsgCrypt::GenRandStr(std::string & sRandStr, uint32_t len)
{
uint32_t idx = 0;
srand((unsigned)time(NULL));
char tempChar = 0;
sRandStr.clear();

while(idx < len)
{
tempChar = rand()%128;
if(isprint(tempChar))
{
sRandStr.append(1, tempChar);
++idx;
}
}
}
*/

/*
int WXBizMsgCrypt::SetOneFieldToXml(tinyxml2::XMLDocument * pDoc, tinyxml2::XMLNode* pXmlNode, const char * pcFieldName,
const std::string & value, bool bIsCdata)
{
if(!pDoc || !pXmlNode || !pcFieldName)
{
return -1;
}

tinyxml2::XMLElement * pFiledElement = pDoc->NewElement(pcFieldName);
if(NULL == pFiledElement)
{
return -1;
}

tinyxml2::XMLText * pText = pDoc->NewText(value.c_str());
if(NULL == pText)
{
return -1;
}

pText->SetCData(bIsCdata);
pFiledElement->LinkEndChild(pText);

pXmlNode->LinkEndChild(pFiledElement);
return 0;
}
*/

func (this *WXBizMsgCrypt) GenReturnXml(sEncryptMsg string, sSignature string, sTimeStamp string, sNonce string) (ret int, sResult string) {

	var msg MsgEncrypt
	msg.Encrypt = sEncryptMsg
	msg.MsgSignature = sSignature
	msg.TimeStamp = sTimeStamp
	msg.Nonce = sNonce
	sr, err := json.Marshal(msg)
	if err != nil {
		ret = -1
		return
	}

	sResult = string(sr)
	return
}
