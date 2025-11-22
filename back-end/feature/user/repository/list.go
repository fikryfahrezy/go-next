package repository

import (
	"context"
)

func (r *userRepository) List(ctx context.Context, limit, offset int) ([]User, int64, error) {
	usersMap, ok := r.db.V["users"]
	if !ok {
		return []User{}, 0, nil
	}

	var users []User
	for _, row := range usersMap {
		user, ok := row.(User)
		if !ok {
			continue
		}
		users = append(users, user)
	}

	return users, int64(len(users)), nil
}
