package requests

type RefreshTokenRequest struct {
	RefreshToken string `json:"refreshToken" binding:"required"`
}
