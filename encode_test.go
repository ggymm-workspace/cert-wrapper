package main

import "testing"

func Test_Base64Encode(t *testing.T) {
	t.Log(encode([]byte(`{"filepath":"cert/certificate.pem"}`)))
}

func Test_Base64Decode(t *testing.T) {
	t.Log(decode("eyJuZXciOnRydWUsImNvdW50cnkiOiJDTiIsIk9yZ2FuaXphdGlvbiI6IkNQTyIsIkNvbW1vbk5hbWUiOiJKWF9DQl8wOTcifQ=="))
}
