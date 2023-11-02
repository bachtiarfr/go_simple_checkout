package repositories

import (
	"database/sql"
	"errors"
	"simple-checkout-app/internal/entity"
	"time"
)

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{db}
}

func (r *userRepository) FindByEmail(email string) (*entity.User, error) {
	query := "SELECT id, email, password FROM users WHERE email = $1"
	var user entity.User

	err := r.db.QueryRow(query, email).Scan(&user.ID, &user.Email, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, err
		}
		return nil, err
	}

	return &user, nil
}

func (r *userRepository) FindAll() ([]entity.ListUser, error) {
	query := "SELECT id, fullname, email, age, phone_number FROM users"
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []entity.ListUser
	for rows.Next() {
		var user entity.ListUser
		err := rows.Scan(&user.ID, &user.Fullname, &user.Email, &user.Age, &user.PhoneNumber)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func (r *userRepository) CreateNewUser(user *entity.User) (entity.User, error) {
	query := "INSERT INTO users (fullname, email, age, password, phone_number) VALUES ($1, $2, $3, $4, $5) RETURNING id"

	err := r.db.QueryRow(query, user.Fullname, user.Email, user.Age, user.Password, user.PhoneNumber).Scan(&user.ID)
	if err != nil {
		return entity.User{}, err
	}

	return entity.User{}, err
}

func (r *userRepository) CreateOrUpdateRefreshToken(token string, userID int) error {
	var existingRefreshToken string
	err := r.db.QueryRow("SELECT refresh_token FROM refresh_tokens WHERE user_id = $1", userID).Scan(&existingRefreshToken)
	if err != nil {
		if err != sql.ErrNoRows {
			return err
		}
	}

	if existingRefreshToken != "" {
		_, err := r.db.Exec("UPDATE refresh_tokens SET refresh_token = $1 WHERE user_id = $2", token, userID)
		if err != nil {
			return err
		}
	} else {
		_, err = r.db.Exec("INSERT INTO refresh_tokens (user_id, refresh_token, expiration) VALUES ($1, $2, $3)", userID, token, time.Now().Add(time.Hour*24))
		if err != nil {
			return err
		}
	}

	return nil
}

func (r *userRepository) GetUserByRefreshToken(token string) (*entity.ListUser, error) {
	var user entity.ListUser
	var expiration time.Time

	err := r.db.QueryRow("SELECT users.id, users.fullname, users.email, users.age, users.phone_number, refresh_tokens.expiration FROM users "+
		"JOIN refresh_tokens ON users.id = refresh_tokens.user_id "+
		"WHERE refresh_tokens.refresh_token = $1", token).Scan(&user.ID, &user.Fullname, &user.Email, &user.Age, &user.PhoneNumber, &expiration)
	if err != nil {
		return nil, err
	}

	if expiration.Before(time.Now()) {
		return nil, errors.New("Refresh token expired")
	}

	return &user, nil
}
