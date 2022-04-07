package services

import "VoteMe_BE_TX/models"

func GetCandidatesDB() (candidates []models.Candidates) {
	db.Find(&candidates)
	return
}
func GetOneCandidateDB(name string) (candidate models.Candidates, err error) {
	err = db.Where("name = ?", name).Find(&candidate).Error
	return
}
