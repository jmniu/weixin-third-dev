package wxopencrypt

type VerifyTicketEncrypt struct {
	AppId   string
	Encrypt string
}

type MsgEncrypt struct {
	Encrypt      string
	MsgSignature string
	TimeStamp    string
	Nonce        string
}
