package repositories

import (
	"segmenta/src/configs"
	"segmenta/src/models"
)

// Explore Course Repositories

func GetAllAvailableCourses(userID uint) ([]models.ExploreCourse, error) {
	var courses []models.ExploreCourse
	result := configs.Database.Find(&courses)
	return courses, result.Error
}

func GetExploredCourseByID(id uint) (*models.ExploreCourse, error) {
	var course models.ExploreCourse
	result := configs.Database.First(&course, "explore_course_id = ?", id)
	return &course, result.Error
}

func SearchCourses(query string) ([]models.ExploreCourse, error) {
	var courses []models.ExploreCourse
	result := configs.Database.Where("title ILIKE ?", "%"+query+"%").Find(&courses)
	return courses, result.Error
}

func EnrollUserInCourse(userID uint, courseID uint) error {
	var templateCourse models.ExploreCourse
	result := configs.Database.First(&templateCourse, "explore_course_id = ?", courseID)
	if result.Error != nil {
		return result.Error
	}

	// Copy course data
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
		SourceVersion:        *templateCourse.Version,
	}

	if err := configs.Database.Create(&newCourse).Error; err != nil {
		return err
	}

	// Copy chapters from explore course to user course
	var exploreChapters []models.ExploreChapter
	configs.Database.Where("explore_course_id = ?", courseID).Order("\"order\" ASC").Find(&exploreChapters)

	for i, ec := range exploreChapters {
		chapter := models.Chapter{
			CourseID:  int(newCourse.CourseID),
			Title:     ec.Title,
			Position:  i + 1,
		}
		configs.Database.Create(&chapter)
	}

	return nil
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