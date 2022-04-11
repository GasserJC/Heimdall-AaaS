//https://go.dev/play/p/awjIIj8Kwms
package user

type UC struct {
	Key string
	Username string
	Password string
}

func (r *UC) Slice() []string {
	return []string{r.Key, r.Username, r.Password}
}

