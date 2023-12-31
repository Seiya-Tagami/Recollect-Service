package user

import "github.com/Seiya-Tagami/Recollect-Service/api/domain/entity"

//go:generate mockgen -source=$GOFILE -destination=$GOPATH/Recollect-Service/api/mock/$GOPACKAGE/$GOFILE -package=mock_$GOPACKAGE
type Repository interface {
	Insert(user *entity.User) error
	GetAnalysisDataBySub(sub string) (AnalysisData, error)
	SelectBySub(user *entity.User, sub string) error
	UpdateBySub(user *entity.User, sub string) error
	DeleteBySub(sub string) error
	ExistsByEmail(email string) (bool, error)
	ExistsByUserID(userID string) (bool, error)
}

type AnalysisData struct {
	CardTitleString      string
	AnalysisResultString string
}
