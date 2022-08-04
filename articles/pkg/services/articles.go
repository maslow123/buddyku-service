package services

import (
	"context"
	"database/sql"
	"log"
	"net/http"
	"time"

	"github.com/maslow123/buddyku-users/pkg/pb"
)

func (s *Server) CreateArticle(ctx context.Context, req *pb.CreateArticleRequest) (*pb.GenericResponse, error) {
	if req.Title == "" {
		return genericResponse(http.StatusBadRequest, "invalid-title")
	}
	if req.Content == "" {
		return genericResponse(http.StatusBadRequest, "invalid-content")
	}
	if req.CreatedBy == 0 {
		return genericResponse(http.StatusBadRequest, "invalid-created-by")
	}

	// Insert to article table
	q := `
		INSERT INTO articles
		(title, content, created_by)
		VALUES
		($1, $2, $3)
		RETURNING id;
	`
	row := s.DB.QueryRowContext(ctx, q, &req.Title, &req.Content, &req.CreatedBy)
	var lastInsertedId int32

	err := row.Scan(&lastInsertedId)
	if err != nil {
		log.Println(err)
		return genericResponse(http.StatusInternalServerError, err.Error())
	}

	resp, err := genericResponse(http.StatusCreated, "")

	var article pb.Article
	article.Id = lastInsertedId

	resp.Article = &article

	return resp, err
}

func (s *Server) DetailArticle(ctx context.Context, req *pb.ArticleDetailRequest) (*pb.GenericResponse, error) {
	if req.Id == 0 {
		return genericResponse(http.StatusBadRequest, "invalid-article-id")
	}

	// update point + 1
	q := `
		UPDATE articles 
		SET point = point + 1
		WHERE id = $1
	`

	_, err := s.DB.ExecContext(ctx, q, req.Id)
	if err != nil {
		log.Println(err)
		return genericResponse(http.StatusInternalServerError, err.Error())
	}

	q = `
		SELECT id, title, content, point, view, created_by, created_at, updated_at
		FROM articles
		WHERE id = $1
	`

	var article pb.Article
	var createdAt, updatedAt time.Time

	row := s.DB.QueryRowContext(ctx, q, req.Id)
	err = row.Scan(
		&article.Id,
		&article.Title,
		&article.Content,
		&article.Point,
		&article.View,
		&article.CreatedBy,
		&createdAt,
		&updatedAt,
	)

	if err != nil {
		log.Println(err)
		if err == sql.ErrNoRows {
			return genericResponse(http.StatusNotFound, "article-not-found")
		}
		return genericResponse(http.StatusInternalServerError, err.Error())
	}

	article.CreatedAt = int32(createdAt.Unix())
	article.UpdatedAt = int32(updatedAt.Unix())

	resp, err := genericResponse(http.StatusOK, "")
	resp.Article = &article

	return resp, err
}

func (s *Server) GetUserArticlePoint(ctx context.Context, req *pb.UserArticlePointRequest) (*pb.UserArticlePointResponse, error) {
	if req.CreatedBy == 0 {
		return genericUserArticlePointResponse(http.StatusBadRequest, "invalid-created-by")
	}

	q := `SELECT SUM(point) total_point FROM articles WHERE created_by = $1`
	var totalPoint int32

	row := s.DB.QueryRowContext(ctx, q, req.CreatedBy)
	err := row.Scan(&totalPoint)
	if err != nil {
		log.Println(err)
		return genericUserArticlePointResponse(http.StatusNotFound, "article-not-found")
	}

	resp, err := genericUserArticlePointResponse(http.StatusOK, "")
	resp.TotalPoint = totalPoint

	return resp, err
}
