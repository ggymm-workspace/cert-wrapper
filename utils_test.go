package main

import (
	"testing"
)

func Test_Base64Encode(t *testing.T) {
	t.Log(encode([]byte(`{"country":"CN","Organization":"CPO","CommonName":"JX_CB_097"}`)))
	t.Log(encode([]byte(`{"cert":"cert/certificate.pem"}`)))
}

func Test_Base64Decode(t *testing.T) {
	t.Log(string(decode("c8XPa9ywS0lGvqfVJ1hSOK2atoItAe/2Tf+6rmrxW8t5mIGqCZ9ubygUisaBmqOW7fmJETVooGJleziqWHiXPYAA/hlFAEKqyyrGi1wEs/jnqjXpSca0IoNF8DPwsMdNaz/wJePgv3E70TtHmT9122mksUkqES8nsjC6zz2iRq8FwzfuvdYy9AAvSE8MkD/IsLvGMg8nz8Ub8XjzslvfVWH6VoMOJzYCYpGIl/PClGG7wkg4tejCGKd9nYwyUmS9hSmOOCcwvgyrTu0muprQ+SPYoyXEo+7CYqND4n24JlpMXUbH0HYuFpDGhWRkatqO8G49HVqrRXc7RcPlYYdeZw==")))
}
