package jwt

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"google.golang.org/grpc/metadata"
)

// Secret key để ký JWT (nên lưu trong biến môi trường)
var (
	jwtSecret        = []byte("your-secret-key")
	ErrNoMetadata    = errors.New("Error No Metadata")
	ErrEmptyMetadata = errors.New("Error Empty Metadata")
)

func Claim(ctx context.Context, outputHeader string) (string, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return "", ErrEmptyMetadata
	}

	data := md.Get(outputHeader)
	if len(data) <= 0 {
		return "", ErrEmptyMetadata
	}
	return data[0], nil
}

// Hàm tạo JWT
func GenerateJWT(id string) (string, error) {
	// Tạo payload cho JWT
	claims := jwt.MapClaims{
		"id":  id,                                    // Gán id vào JWT
		"exp": time.Now().Add(24 * time.Hour).Unix(), // Hạn sử dụng 24 giờ
	}

	// Tạo token mới với phương thức ký HS256
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Ký token bằng secret key
	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// Hàm xác minh và giải mã JWT
func ValidateJWT(tokenString string) (string, error) {
	// Parse token và xác minh chữ ký
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Xác minh rằng thuật toán ký là HS256
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return jwtSecret, nil
	})

	if err != nil {
		return "", err
	}

	// Lấy claims từ token
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// Trích xuất id từ claims
		id := claims["id"].(string)
		return id, nil
	}

	return "", fmt.Errorf("invalid token")
}
