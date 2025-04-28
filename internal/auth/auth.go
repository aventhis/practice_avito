package auth

type AuthService struct {
	JWTSecret string
}

func NewAuthService(JWTSecret string) *AuthService {
	return &AuthService{
		JWTSecret: JWTSecret,
	}
}
