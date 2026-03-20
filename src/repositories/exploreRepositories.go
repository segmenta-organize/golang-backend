package repositories

import (
	"segmenta/src/configs"
	"segmenta/src/models"
)

// Explore Course Repositories


func GetAllExploreCourses() ([]models.ExploreCourse, error) {
	var exploreCourses []models.ExploreCourse
	result := configs.Database.Find(&exploreCourses)
	return exploreCourses, result.Error
}

func GetExploreCourseByID(id uint) (*models.ExploreCourse, error) {
	var targetCourse models.ExploreCourse
	result := configs.Database.First(&targetCourse, "explore_course_id = ?", id)
	return &targetCourse, result.Error
}

func SearchExploreCourses(query string) ([]models.ExploreCourse, error) {
	var exploreCourses []models.ExploreCourse
	result := configs.Database.Where("title ILIKE ?", "%"+query+"%").Find(&exploreCourses)
	return exploreCourses, result.Error
}

func GetAllCoursesByCategoryForExplore(category string) ([]models.ExploreCourse, error) {
	var exploreCourses []models.ExploreCourse
	result := configs.Database.Where("category = ?", category).Find(&exploreCourses)
	return exploreCourses, result.Error
}

func EnrollUserInCourse(userID uint, courseID uint) error {
	var templateCourse models.ExploreCourse
	result := configs.Database.First(&templateCourse, "explore_course_id = ?", courseID)
	if result.Error != nil {
		return result.Error
	}

	sourceVersion := 1
	if templateCourse.Version != nil {
		sourceVersion = *templateCourse.Version
	}

	sourceID := int(courseID)
	newCourse := models.Course{
		UserID:               int(userID),
		Title:                templateCourse.Title,
		Description:          templateCourse.Description,
		Channel:              templateCourse.Channel,
		ChannelLink:          templateCourse.ChannelLink,
		VideoLink:            templateCourse.VideoLink,
		ThumbnailImageURL:    templateCourse.ThumbnailImageURL,
		SourcePublicCourseID: &sourceID,
		SourceVersion:        sourceVersion,
	}

	if errorHandler := configs.Database.Create(&newCourse).Error; errorHandler != nil {
		return errorHandler
	}

	var exploreChapters []models.ExploreChapter
	configs.Database.Where("explore_course_id = ?", courseID).Order("\"order\" ASC").Find(&exploreChapters)
	
	for i, exploreChapter := range exploreChapters {
		chapter := models.Chapter{
			CourseID: int(newCourse.CourseID),
			Title:    exploreChapter.Title,
			Position: i + 1,
		}
		configs.Database.Create(&chapter)
	}

	return nil
}

func EditPublicCourse(courseID uint, updatedCourse *models.ExploreCourse) error {
	return configs.Database.Model(&models.ExploreCourse{}).Where("explore_course_id = ?", courseID).Updates(updatedCourse).Error
}

func DeleteOnePublicCourseByID(courseID uint) error {
	if errorHandler := DeleteExploreChaptersByCourseID(courseID); errorHandler != nil {
		return errorHandler
	}
	return configs.Database.Where("explore_course_id = ?", courseID).Delete(&models.ExploreCourse{}).Error
}

func DeleteExploreChapterByUserID(userID uint) error {
	return configs.Database.Where("user_id = ?", userID).Delete(&models.ExploreChapter{}).Error
}

// Explore Chapter Repositories

func GetAllExploreChaptersByCourseID(courseID uint) ([]models.ExploreChapter, error) {
	var chapters []models.ExploreChapter
	result := configs.Database.Where("explore_course_id = ?", courseID).Order("\"order\" ASC").Find(&chapters)
	return chapters, result.Error
}

func GetOneExploreChapterByID(id uint) (*models.ExploreChapter, error) {
	var chapter models.ExploreChapter
	result := configs.Database.First(&chapter, "explore_chapter_id = ?", id)
	return &chapter, result.Error
}

func CreateExploreChapter(chapter *models.ExploreChapter) error {
	return configs.Database.Create(chapter).Error
}

func UpdateExploreChapterByID(id uint, chapter *models.ExploreChapter) error {
	return configs.Database.Model(&models.ExploreChapter{}).Where("explore_chapter_id = ?", id).Updates(chapter).Error
}

func DeleteExploreChapterByID(id uint) error {
	return configs.Database.Where("explore_chapter_id = ?", id).Delete(&models.ExploreChapter{}).Error
}

func DeleteExploreChaptersByCourseID(courseID uint) error {
	return configs.Database.Where("explore_course_id = ?", courseID).Delete(&models.ExploreChapter{}).Error
}

func DeleteExploreCourseByUserID(userID uint) error {
	return configs.Database.Where("creator_id = ?", userID).Delete(&models.ExploreCourse{}).Error
}