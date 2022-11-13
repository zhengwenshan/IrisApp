package main

//noinspection GoInvalidPackageImport
import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/jwt"
	"time"
)

var (
	sigKey = []byte("signature_hmac_secret_shared_key")
	encKey = []byte("GCM_AES_256_secret_shared_key_32")
)

type fooClaims struct {
	Foo string `json:"foo"`
}

func main() {

	signer := jwt.NewSigner(jwt.HS256, sigKey, 10*time.Minute)

	verifier := jwt.NewVerifier(jwt.HS256, sigKey)
	verifier.WithDefaultBlocklist()

	verify := verifier.Verify(func() interface{} {
		return new(fooClaims)
	})

	app := iris.New()
	protectedAPI := app.Party("/protected")
	protectedAPI.Use(verify)
	println(signer)
	//signer.Sign()
}
