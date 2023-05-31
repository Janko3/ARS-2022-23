package idempotency

var (
	idempotencyMap map[string]string
)

func GetIdempotencyMap() *map[string]string {
	if idempotencyMap == nil {
		idempotencyMap = make(map[string]string)
	}
	return &idempotencyMap
}
