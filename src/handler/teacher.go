package handler

import (
	db "db_course_design_backend/src/database"
	"db_course_design_backend/src/model"
	"db_course_design_backend/src/utils"
	"db_course_design_backend/src/utils/e"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"log"
	"net/http"
)

// 创建教师
func TeacherCreate(c *gin.Context) {
	teacher := model.TeacherInfo{}

	if c.ShouldBindBodyWith(&teacher, binding.JSON) != nil || teacher.TeaNo == "" || teacher.TeaName == "" {
		c.JSON(http.StatusOK, model.GetResultByCode(e.INVALID_PARAMS))
		return
	}

	//log.Println("before lock")
	tx := db.GetDB().Begin()
	//log.Println("after lock")
	if err := CreateUser(tx, teacher.TeaNo, model.USERTYPE_TEACHER); err != nil {
		log.Println("err one")
		tx.Rollback()
		c.JSON(http.StatusOK, model.GetResultByCode(e.ERROR_USER_EXIST))
		return
	}
	//log.Println("after c user")
	if err := tx.Create(&teacher).Error; err != nil {
		//log.Println("err two")
		tx.Rollback()
		c.JSON(http.StatusOK, model.GetResultByCode(e.ERROR_USER_EXIST))
		return
	}
	//log.Println("before commit")
	tx.Commit()
	c.JSON(http.StatusOK, model.GetResultByCode(e.SUCCESS))
}

// 更新教师
func TeacherUpdate(c *gin.Context) {
	teacher := model.TeacherInfo{}

	if c.ShouldBindBodyWith(&teacher, binding.JSON) != nil || teacher.TeaNo == "" || teacher.TeaName == "" {
		c.JSON(http.StatusOK, model.GetResultByCode(e.INVALID_PARAMS))
		return
	}

	if db.GetDB().Where(&model.TeacherInfo{TeaNo: teacher.TeaNo}).First(&model.TeacherInfo{}).RecordNotFound() {
		c.JSON(http.StatusOK, model.GetResultByCode(e.ERROR_USER_NOT_EXIST))
		return
	}

	db.GetDB().Model(&teacher).Where(&model.TeacherInfo{TeaNo: teacher.TeaNo}).Update(&teacher)

	c.JSON(http.StatusOK, model.GetResultByCode(e.SUCCESS))
	return
}

// 删除教师
func TeacherDelete(c *gin.Context) {
	teaNo := c.Query(e.KEY_TEA_NO)

	if teaNo == "" {
		c.JSON(http.StatusOK, model.GetResultByCode(e.INVALID_PARAMS))
		return
	}

	db.GetDB().Where(&model.User{
		UserId: teaNo,
	}).Delete(&model.User{})

	c.JSON(http.StatusOK, model.GetResultByCode(e.SUCCESS))
	return
}

// 教师查询
func TeacherQuery(c *gin.Context) {
	teaNo, teaNoExist := c.GetQuery(e.KEY_TEA_NO)
	page, pageExist := c.GetQuery(e.KEY_PAGE)
	var teachers []model.TeacherInfo

	query := db.GetDB().Model(&model.TeacherInfo{})
	if teaNoExist {
		query = query.Where(&model.TeacherInfo{TeaNo: teaNo})
	}

	if pageExist {
		result := model.GetResultByCode(e.SUCCESS)
		payload := utils.GenPagePayload(query, page, &teachers)
		result.Data = payload
		c.JSON(http.StatusOK, result)
	} else {
		query.Find(&teachers)
		result := model.GetResultByCode(e.SUCCESS)
		result.Data = teachers
		c.JSON(http.StatusOK, result)
	}
}
