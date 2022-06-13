package service

import (
	"SimpleTikTok/commom"
	"SimpleTikTok/dao"
	"SimpleTikTok/model"
)

//func CommentsList(id int64) ([]model.Comments, error) {
//	commentDAO := dao.NewCommentDao()
//	//var comments []model.Comments
//	var comments = commentDAO.QueryCommentsById(id)
//	return comments, nil
//}
func CommentsList(id int64) ([]commom.Comment, error) {
	commentDAO := dao.NewCommentDaoByTHJ()
	//var comments []model.Comments
	var comments = commentDAO.QueryCommentsById(id)
	if comments != nil {

	}

	var commontsCommom []commom.Comment

	db := commom.GetDB()
	for i, comment := range comments {
		var commentCommon commom.Comment
		var userinfo model.Users
		commentCommon.Id = int64(comment.ID)
		commentCommon.CreateDate = comment.CreateDate.Format("01.02")
		commentCommon.Content = comment.Content
		db.Table("users").Where("id=?", comment.UserId).Find(&userinfo)
		commentCommon.User.Id = userinfo.ID
		commentCommon.User.Name = userinfo.Username
		commentCommon.User.FollowerCount = userinfo.FansCounts
		commentCommon.User.FollowCount = userinfo.FollowCounts
		commentCommon.User.IsFollow = false
		commontsCommom = append(commontsCommom, commentCommon)
		if i < 0 {

		}
		//commontsCommom[i]:=&commom.Comment{
		//	Id:        comment.ID,
		//	User: comment.UserId,
		//	Content: comment.Content,
		//	CreateDate: comment.CreateDate,
		//}
		//fmt.Println(comment)
		//fmt.Println(i)

	}
	//
	//commontsCommom:=&commom.Comment{
	//	Id:        comments[0].ID
	//}
	return commontsCommom, nil
}
