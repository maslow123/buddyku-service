package services

import (
	"context"
	"database/sql"
	"log"
	"net/http"
	"time"

	"github.com/maslow123/buddyku-users/pkg/pb"
	"github.com/maslow123/buddyku-users/pkg/utils"
)

func (s *Server) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.GenericResponse, error) {
	if req.Name == "" {
		return genericResponse(http.StatusBadRequest, "invalid-name")
	}
	if req.Username == "" {
		return genericResponse(http.StatusBadRequest, "invalid-username")
	}
	if req.Password == "" {
		return genericResponse(http.StatusBadRequest, "invalid-password")
	}
	if req.Level > int32(1) || req.Level < int32(0) {
		return genericResponse(http.StatusBadRequest, "invalid-level")
	}

	hashedPassword := utils.HashPassword(req.Password)

	// Check username already exists
	q := `SELECT username FROM users where username = $1`

	row := s.DB.QueryRowContext(ctx, q, req.Username)
	var username string

	_ = row.Scan(&username)

	if username != "" {
		return genericResponse(http.StatusBadRequest, "username-already-exists")
	}

	q = `
		INSERT INTO users (name, username, password, level)
		VALUES ($1, $2, $3, $4)
		RETURNING id;
	`
	row = s.DB.QueryRowContext(ctx, q,
		&req.Name,
		&req.Username,
		&hashedPassword,
		&req.Level,
	)

	var lastInsertedId int64

	err := row.Scan(&lastInsertedId)
	if err != nil {
		log.Println(err)
		return genericResponse(http.StatusInternalServerError, err.Error())
	}

	return genericResponse(http.StatusCreated, "")
}

func (s *Server) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	if req.Username == "" {
		return genericLoginResponse(http.StatusBadRequest, "invalid-username")
	}
	if req.Password == "" {
		return genericLoginResponse(http.StatusBadRequest, "invalid-password")
	}

	var userPass string
	var userID int32
	q := `
		SELECT id, password
		FROM users
		WHERE username = $1
		LIMIT 1
	`
	row := s.DB.QueryRowContext(ctx, q, req.Username)

	err := row.Scan(
		&userID,
		&userPass,
	)

	if err != nil {
		log.Println(err)
		if err == sql.ErrNoRows {
			return genericLoginResponse(http.StatusNotFound, "user-not-found")
		}
		return genericLoginResponse(http.StatusInternalServerError, err.Error())
	}

	match := utils.CheckPasswordHash(req.Password, userPass)
	if !match {
		return genericLoginResponse(http.StatusUnauthorized, "password-not-match")
	}

	token, _ := s.Jwt.GenerateToken(userID)
	resp := &pb.LoginResponse{
		Status: http.StatusOK,
		Error:  "",
		Token:  token,
	}

	return resp, nil
}

func (s *Server) Validate(ctx context.Context, req *pb.ValidateRequest) (*pb.ValidateResponse, error) {
	claims, err := s.Jwt.ValidateToken(req.Token)

	if err != nil {
		return &pb.ValidateResponse{
			Status: http.StatusBadRequest,
			Error:  err.Error(),
		}, nil
	}

	return &pb.ValidateResponse{
		Status: http.StatusOK,
		UserId: claims.UserID,
	}, nil
}

func (s *Server) ShowRegisterUser(ctx context.Context, req *pb.ShowUserRequest) (*pb.ShowUserResponse, error) {
	if req.Limit == 0 {
		return genericShowRegisterResponse(http.StatusBadRequest, "invalid-limit")
	}
	if req.Page == 0 {
		return genericShowRegisterResponse(http.StatusBadRequest, "invalid-page")
	}
	// only show publisher/user not company.
	q := `
		SELECT username, name, created_at, updated_at
		FROM users
		WHERE level = 0
		LIMIT $1 OFFSET $2
	`
	offset := (req.Page - 1) * req.Limit

	rows, err := s.DB.QueryContext(ctx, q, req.Limit, offset)
	if err != nil {
		log.Println(err)
		if err == sql.ErrNoRows {
			return genericShowRegisterResponse(http.StatusNotFound, "user-not-found")
		}
		return genericShowRegisterResponse(http.StatusInternalServerError, err.Error())
	}

	var users []*pb.User

	for rows.Next() {
		var user pb.User
		var createdAt, updatedAt time.Time
		if err := rows.Scan(
			&user.Username,
			&user.Name,
			&createdAt,
			&updatedAt,
		); err != nil {
			log.Println(err)
			return genericShowRegisterResponse(http.StatusInternalServerError, err.Error())
		}

		user.CreatedAt = int32(createdAt.Unix())
		user.UpdatedAt = int32(updatedAt.Unix())

		users = append(users, &user)
	}

	if err := rows.Close(); err != nil {
		return genericShowRegisterResponse(http.StatusInternalServerError, err.Error())
	}

	if err := rows.Err(); err != nil {
		return genericShowRegisterResponse(http.StatusInternalServerError, err.Error())
	}

	if len(users) == 0 {
		return genericShowRegisterResponse(http.StatusNotFound, "user-not-found")
	}

	resp, err := genericShowRegisterResponse(http.StatusOK, "")
	resp.Users = users

	return resp, err
}

func (s *Server) ShowAllUserPoint(ctx context.Context, req *pb.ShowUserRequest) (*pb.ShowUserResponse, error) {
	if req.Limit == 0 {
		return genericShowRegisterResponse(http.StatusBadRequest, "invalid-limit")
	}
	if req.Page == 0 {
		return genericShowRegisterResponse(http.StatusBadRequest, "invalid-page")
	}
	// only show publisher/user not company, also get each user point
	q := `
		SELECT u.username, u.name, u.created_at, u.updated_at, COALESCE(SUM(point),0) user_point
		FROM users u
		LEFT JOIN articles a 
			ON u.id = a.created_by
		WHERE u.level = 0 
		GROUP BY u.id
		LIMIT $1 OFFSET $2
	`
	offset := (req.Page - 1) * req.Limit

	rows, err := s.DB.QueryContext(ctx, q, req.Limit, offset)
	if err != nil {
		log.Println(err)
		if err == sql.ErrNoRows {
			return genericShowRegisterResponse(http.StatusNotFound, "user-not-found")
		}
		return genericShowRegisterResponse(http.StatusInternalServerError, err.Error())
	}

	var users []*pb.User

	for rows.Next() {
		var user pb.User
		var createdAt, updatedAt time.Time
		if err := rows.Scan(
			&user.Username,
			&user.Name,
			&createdAt,
			&updatedAt,
			&user.UserPoint,
		); err != nil {
			log.Println(err)
			return genericShowRegisterResponse(http.StatusInternalServerError, err.Error())
		}

		user.CreatedAt = int32(createdAt.Unix())
		user.UpdatedAt = int32(updatedAt.Unix())

		users = append(users, &user)
	}

	if err := rows.Close(); err != nil {
		return genericShowRegisterResponse(http.StatusInternalServerError, err.Error())
	}

	if err := rows.Err(); err != nil {
		return genericShowRegisterResponse(http.StatusInternalServerError, err.Error())
	}

	if len(users) == 0 {
		return genericShowRegisterResponse(http.StatusNotFound, "user-not-found")
	}

	resp, err := genericShowRegisterResponse(http.StatusOK, "")
	resp.Users = users

	return resp, err
}

func (s *Server) SetPoint(ctx context.Context, req *pb.SetPointRequest) (*pb.GenericResponse, error) {
	if req.ArticleId == 0 {
		return genericResponse(http.StatusBadRequest, "invalid-article-id")
	}
	if req.UserId == 0 {
		return genericResponse(http.StatusBadRequest, "invalid-user-id")
	}
	if req.Point == 0 {
		return genericResponse(http.StatusBadRequest, "invalid-point")
	}
	// check user id is found or not, also check level of user is company or not.
	count, err := s.checkCompanyIsExist(ctx, req.UserId)
	if err != nil {
		log.Println(err)
		return genericResponse(http.StatusInternalServerError, err.Error())
	}
	if count == 0 {
		return genericResponse(http.StatusNotFound, "user-not-found")
	}

	// check article already exists
	count, err = s.checkArticleIsExists(ctx, req.ArticleId)
	if err != nil {
		log.Println(err)
		return genericResponse(http.StatusInternalServerError, err.Error())
	}
	if count == 0 {
		return genericResponse(http.StatusNotFound, "article-not-found")
	}

	// set article point
	q := `
		UPDATE articles
		SET point = point + $1, updated_at = now()
		WHERE id = $2
	`

	_, err = s.DB.ExecContext(ctx, q, req.Point, req.ArticleId)
	if err != nil {
		log.Println(err)
		return genericResponse(http.StatusInternalServerError, err.Error())
	}

	resp, err := genericResponse(http.StatusOK, "")
	return resp, err
}

func (s *Server) checkCompanyIsExist(ctx context.Context, userID int32) (int32, error) {
	var count int32
	q := `
		SELECT count(1) count 
		FROM users
		WHERE level = 1 AND id = $1
	`
	row := s.DB.QueryRowContext(ctx, q, userID)
	err := row.Scan(&count)
	if err != nil {
		return 0, err
	}

	return count, err
}

func (s *Server) checkArticleIsExists(ctx context.Context, articleID int32) (int32, error) {
	var count int32
	q := `
		SELECT count(1) count 
		FROM articles
		WHERE id = $1
	`
	row := s.DB.QueryRowContext(ctx, q, articleID)
	err := row.Scan(&count)
	if err != nil {
		return 0, err
	}

	return count, err
}
