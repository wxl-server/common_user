package domain

type UserDO struct {
	ID        int64
	Email     string
	Password  string
	Extra     *string
	CreatedAt int64
	UpdatedAt int64
}
