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
func AddScoreToCandidate(name string) (err error) {
	var candidate models.Candidates
	db.Where("name = ?", name).Find(&candidate)
	var temp models.Candidates
	temp.Name = candidate.Name
	temp.Score = candidate.Score + 1
	db.Where("name = ?", name).Delete(&candidate)
	db.Save(&temp)
	return nil
}
