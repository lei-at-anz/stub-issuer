package key

type KeyType string

const (
	KeyTypeES256 KeyType = "ES256"
	KeyTypeES384 KeyType = "ES384"
	KeyTypeES512 KeyType = "ES512"
)

type SigningKey struct {
	Key interface{}
	ID  string
}
