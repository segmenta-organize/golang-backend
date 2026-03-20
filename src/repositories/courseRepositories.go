package repositories

import (
	"segmenta/src/configs"
	"segmenta/src/models"
)

// Course Repositories

func GetAllEnrolledCourses(userID uint) ([]models.Course, error) {
	var courses []models.Course
	result := configs.Database.Where("user_id = ?", userID).Find(&courses)
	return courses, result.Error
}

func GetOneCourseByID(id uint) (*models.Course, error) {
	var course models.Course
	result := configs.Database.First(&course, "course_id = ?", id)
	return &course, result.Error
}

func CreateCourse(course *models.Course) error {
	return configs.Database.Create(course).Error
}

func DeleteAllChaptersInOneCourseByID(courseID uint) error {
	return configs.Database.Where("course_id = ?", courseID).Delete(&models.Chapter{}).Error
}

func UpdateCourseByID(id uint, course *models.Course) error {
	return configs.Database.Model(&models.Course{}).Where("course_id = ?", id).Updates(course).Error
}

func DeleteOneCourseByID(id uint) error {
	return configs.Database.Where("course_id = ?", id).Delete(&models.Course{}).Error
}

func CreatePublicCourseFromCourse(courseID uint, userID uint) error {
	version := 1
	var course models.Course
	if errorHandler := configs.Database.First(&course, "course_id = ?", courseID).Error; errorHandler != nil {
		return errorHandler
	}

	exploreCourse := models.ExploreCourse{
		CreatorID:         int(userID),
		Title:             course.Title,
		Description:       course.Description,
		Channel:           course.Channel,
		ChannelLink:       course.ChannelLink,
		VideoLink:         course.VideoLink,
		ThumbnailImageURL: course.ThumbnailImageURL,
		Version:           &version,
	}

	return configs.Database.Create(&exploreCourse).Error
}

func UpdatePublicCourseFromCourse(courseID uint, exploreCourseID uint) error {
	var course models.Course
	if errorHandler := configs.Database.First(&course, "course_id = ?", courseID).Error; errorHandler != nil {
		return errorHandler
	}

	return configs.Database.Model(&models.ExploreCourse{}).
		Where("explore_course_id = ?", exploreCourseID).
		Updates(map[string]interface{}{
			"title":               course.Title,
			"description":         course.Description,
			"channel":             course.Channel,
			"channel_link":        course.ChannelLink,
			"video_link":          course.VideoLink,
			"thumbnail_image_url": course.ThumbnailImageURL,
			"version":             configs.Database.Raw("version + 1"),
		}).Error
}

func CheckCourseLinkExists(userID uint, link string) (bool, error) {
	var count int64
	result := configs.Database.Model(&models.Course{}).Where("user_id = ? AND video_link = ?", userID, link).Count(&count)
	return count > 0, result.Error
}

func DeleteCourseByUserID(userID uint) error {
	return configs.Database.Where("user_id = ?", userID).Delete(&models.Course{}).Error
}

// Chapter Repositories

func GetAllChaptersByCourseID(courseID uint) ([]models.Chapter, error) {
	var chapters []models.Chapter
	result := configs.Database.Where("course_id = ?", courseID).Order("position ASC").Find(&chapters)
	return chapters, result.Error
}

func GetChapterByID(id uint) (*models.Chapter, error) {
	var chapter models.Chapter
	result := configs.Database.First(&chapter, "chapter_id = ?", id)
	return &chapter, result.Error
}

func CreateChapter(chapter *models.Chapter) error {
	return configs.Database.Create(chapter).Error
}

func UpdateChapterByID(id uint, chapter *models.Chapter) error {
	return configs.Database.Model(&models.Chapter{}).Where("chapter_id = ?", id).Updates(chapter).Error
}

func DeleteChapterByID(id uint) error {
	return configs.Database.Where("chapter_id = ?", id).Delete(&models.Chapter{}).Error
}

func DeleteChaptersByCourseID(courseID uint) error {
	return configs.Database.Where("course_id = ?", courseID).Delete(&models.Chapter{}).Error
}

// Chapter → ExploreChapter public functions

func CreatePublicChapterFromChapter(chapterID uint, exploreCourseID uint) error {
	var chapter models.Chapter
	if errorHandler := configs.Database.First(&chapter, "chapter_id = ?", chapterID).Error; errorHandler != nil {
		return errorHandler
	}

	// Count existing explore chapters to determine order
	var count int64
	configs.Database.Model(&models.ExploreChapter{}).Where("explore_course_id = ?", exploreCourseID).Count(&count)

	exploreChapter := models.ExploreChapter{
		ExploreCourseID: exploreCourseID,
		Title:           chapter.Title,
		Order:           int(count) + 1,
	}

	return configs.Database.Create(&exploreChapter).Error
}

func UpdatePublicChapterFromChapter(chapterID uint, exploreChapterID uint) error {
	var chapter models.Chapter
	if errorHandler := configs.Database.First(&chapter, "chapter_id = ?", chapterID).Error; errorHandler != nil {
		return errorHandler
	}

	return configs.Database.Model(&models.ExploreChapter{}).
		Where("explore_chapter_id = ?", exploreChapterID).
		Updates(map[string]interface{}{
			"title": chapter.Title,
		}).Error
}

func DeletePublicChapterFromChapter(exploreChapterID uint) error {
	return configs.Database.Where("explore_chapter_id = ?", exploreChapterID).Delete(&models.ExploreChapter{}).Error
}

func DeleteChapterByUserID(userID uint) error {
	return configs.Database.Where("user_id = ?", userID).Delete(&models.Chapter{}).Error
}
