# Segmenta Backend API Documentation

> Base URL: `http://localhost:3344`
>
> Semua endpoint kecuali `/auth/register` dan `/auth/login` memerlukan header:
> ```
> Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImhpbG1pbXVzeWFmYTE3QGdtYWlsLmNvbSIsImV4cCI6MTc3MzQ5NDk1NywibmFtZSI6IkhpbG1pIE11c3lhZmEiLCJ1c2VyX2lkIjoxfQ.lcKcpOQPe8Aybg_34l4X7dHiMiNvt5tKR4r12kBXoRo
> ```

---

## 1. Authentication

### Register
```bash
curl -X POST http://localhost:3344/auth/register \
  -H "Content-Type: application/json" \
  -d '{
    "full_name": "Hilmi Musyafa",
    "email": "hilmi@example.com",
    "password": "password123"
  }'
```
**Response (200):**
```json
{
  "success": true,
  "message": "[REGISTER] User registered successfully",
  "data": {
    "token": eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
  }
}
```

### Login
```bash
curl -X POST http://localhost:3344/auth/login \
  -H "Content-Type: application/json" \
  -d '{
    "email": "hilmi@example.com",
    "password": "password123"
  }'
```
**Response (200):**
```json
{
  "success": true,
  "message": "[LOGIN] Login successful",
  "data": {
    "token": eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
  }
}
```

### Logout
```bash
curl -X POST http://localhost:3344/auth/logout \
  -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImhpbG1pbXVzeWFmYTE3QGdtYWlsLmNvbSIsImV4cCI6MTc3MzQ5NDk1NywibmFtZSI6IkhpbG1pIE11c3lhZmEiLCJ1c2VyX2lkIjoxfQ.lcKcpOQPe8Aybg_34l4X7dHiMiNvt5tKR4r12kBXoRo"
```
**Response (200):**
```json
{
  "success": true,
  "message": "[LOGOUT] Logout successful"
}
```

### Refresh Token
```bash
curl -X POST http://localhost:3344/auth/refresh \
  -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImhpbG1pbXVzeWFmYTE3QGdtYWlsLmNvbSIsImV4cCI6MTc3MzQ5NDk1NywibmFtZSI6IkhpbG1pIE11c3lhZmEiLCJ1c2VyX2lkIjoxfQ.lcKcpOQPe8Aybg_34l4X7dHiMiNvt5tKR4r12kBXoRo"
```
**Response (200):**
```json
{
  "success": true,
  "message": "[REFRESH] Token refreshed successfully",
  "data": {
    "token": eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
  }
}
```

---

## 2. Courses (User's Personal Courses)

### Get All Courses
```bash
curl -X GET http://localhost:3344/courses/ \
  -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImhpbG1pbXVzeWFmYTE3QGdtYWlsLmNvbSIsImV4cCI6MTc3MzQ5NDk1NywibmFtZSI6IkhpbG1pIE11c3lhZmEiLCJ1c2VyX2lkIjoxfQ.lcKcpOQPe8Aybg_34l4X7dHiMiNvt5tKR4r12kBXoRo"
```
**Response (200):**
```json
{
  "success": true,
  "message": "[GET COURSES] Courses fetched successfully",
  "data": {
    "courses": [
      {
        "course_id": 1,
        "user_id": 1,
        "title": "Belajar Go Lang",
        "description": "Tutorial lengkap Go dari dasar",
        "channel": "ProgrammerZaman",
        "channel_link": "https://youtube.com/@ProgrammerZaman",
        "video_link": "https://youtube.com/watch?v=abc123",
        "thumbnail_image_url": "https://i.ytimg.com/vi/abc123/maxresdefault.jpg",
        "progress": 45,
        "source_public_course_id": null,
        "source_version": 0,
        "update_check": false,
        "created_at": "2026-03-13T12:00:00Z",
        "updated_at": "2026-03-13T12:00:00Z"
      }
    ]
  }
}
```

### Get Course By ID
```bash
curl -X GET http://localhost:3344/courses/1 \
  -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImhpbG1pbXVzeWFmYTE3QGdtYWlsLmNvbSIsImV4cCI6MTc3MzQ5NDk1NywibmFtZSI6IkhpbG1pIE11c3lhZmEiLCJ1c2VyX2lkIjoxfQ.lcKcpOQPe8Aybg_34l4X7dHiMiNvt5tKR4r12kBXoRo"
```
**Response (200):**
```json
{
  "success": true,
  "message": "[GET COURSE BY ID] Course fetched successfully",
  "data": {
    "course": {
      "course_id": 1,
      "user_id": 1,
      "title": "Belajar Go Lang",
      "description": "Tutorial lengkap Go dari dasar",
      "channel": "ProgrammerZaman",
      "channel_link": "https://youtube.com/@ProgrammerZaman",
      "video_link": "https://youtube.com/watch?v=abc123",
      "thumbnail_image_url": "https://i.ytimg.com/vi/abc123/maxresdefault.jpg",
      "progress": 45,
      "source_public_course_id": null,
      "source_version": 0,
      "update_check": false,
      "created_at": "2026-03-13T12:00:00Z",
      "updated_at": "2026-03-13T12:00:00Z"
    }
  }
}
```

### Create Course
```bash
curl -X POST http://localhost:3344/courses/create \
  -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImhpbG1pbXVzeWFmYTE3QGdtYWlsLmNvbSIsImV4cCI6MTc3MzQ5NDk1NywibmFtZSI6IkhpbG1pIE11c3lhZmEiLCJ1c2VyX2lkIjoxfQ.lcKcpOQPe8Aybg_34l4X7dHiMiNvt5tKR4r12kBXoRo" \
  -H "Content-Type: application/json" \
  -d '{
    "video_link": "https://youtube.com/watch?v=abc123"
  }'
```
**Response (200):**
```json
{
  "success": true,
  "message": "[CREATE COURSE] Course created successfully",
  "data": {
    "course": {
      "course_id": 2,
      "user_id": 1,
      "title": "",
      "video_link": "https://youtube.com/watch?v=abc123",
      "progress": 0,
      "source_version": 0,
      "update_check": false,
      "created_at": "2026-03-13T12:00:00Z",
      "updated_at": "2026-03-13T12:00:00Z"
    }
  }
}
```

### Update Course
```bash
curl -X PUT http://localhost:3344/courses/edit/1 \
  -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImhpbG1pbXVzeWFmYTE3QGdtYWlsLmNvbSIsImV4cCI6MTc3MzQ5NDk1NywibmFtZSI6IkhpbG1pIE11c3lhZmEiLCJ1c2VyX2lkIjoxfQ.lcKcpOQPe8Aybg_34l4X7dHiMiNvt5tKR4r12kBXoRo" \
  -H "Content-Type: application/json" \
  -d '{
    "title": "Belajar Go Lang - Updated",
    "progress": 75
  }'
```
**Response (200):**
```json
{
  "success": true,
  "message": "[UPDATE COURSE] Course updated successfully"
}
```

### Delete Course
```bash
curl -X DELETE http://localhost:3344/courses/delete/1 \
  -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImhpbG1pbXVzeWFmYTE3QGdtYWlsLmNvbSIsImV4cCI6MTc3MzQ5NDk1NywibmFtZSI6IkhpbG1pIE11c3lhZmEiLCJ1c2VyX2lkIjoxfQ.lcKcpOQPe8Aybg_34l4X7dHiMiNvt5tKR4r12kBXoRo"
```
**Response (200):**
```json
{
  "success": true,
  "message": "[DELETE COURSE] Course deleted successfully"
}
```

---

## 3. Chapters (Under User's Course)

### Get All Chapters by Course ID
```bash
curl -X GET "http://localhost:3344/chapters/?course_id=1" \
  -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImhpbG1pbXVzeWFmYTE3QGdtYWlsLmNvbSIsImV4cCI6MTc3MzQ5NDk1NywibmFtZSI6IkhpbG1pIE11c3lhZmEiLCJ1c2VyX2lkIjoxfQ.lcKcpOQPe8Aybg_34l4X7dHiMiNvt5tKR4r12kBXoRo"
```
> Note: `course_id` dikirim via route param `/chapters/` dengan controller membaca `c.Param("course_id")`

**Response (200):**
```json
{
  "success": true,
  "message": "[GET CHAPTERS BY COURSE ID] Chapters fetched successfully",
  "data": {
    "chapters": [
      {
        "chapter_id": 1,
        "course_id": 1,
        "title": "Introduction",
        "video_timestamp": "0:00",
        "position": 1,
        "is_completed": false,
        "created_at": "2026-03-13T12:00:00Z",
        "updated_at": "2026-03-13T12:00:00Z"
      },
      {
        "chapter_id": 2,
        "course_id": 1,
        "title": "Setup Environment",
        "video_timestamp": "5:30",
        "position": 2,
        "is_completed": true,
        "created_at": "2026-03-13T12:00:00Z",
        "updated_at": "2026-03-13T12:00:00Z"
      }
    ]
  }
}
```

### Get One Chapter by ID
```bash
curl -X GET http://localhost:3344/chapters/1 \
  -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImhpbG1pbXVzeWFmYTE3QGdtYWlsLmNvbSIsImV4cCI6MTc3MzQ5NDk1NywibmFtZSI6IkhpbG1pIE11c3lhZmEiLCJ1c2VyX2lkIjoxfQ.lcKcpOQPe8Aybg_34l4X7dHiMiNvt5tKR4r12kBXoRo"
```
**Response (200):**
```json
{
  "success": true,
  "message": "[GET CHAPTER BY ID] Chapter fetched successfully",
  "data": {
    "chapter": {
      "chapter_id": 1,
      "course_id": 1,
      "title": "Introduction",
      "video_timestamp": "0:00",
      "position": 1,
      "is_completed": false,
      "created_at": "2026-03-13T12:00:00Z",
      "updated_at": "2026-03-13T12:00:00Z"
    }
  }
}
```

### Create Chapter
```bash
curl -X POST http://localhost:3344/chapters/create \
  -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImhpbG1pbXVzeWFmYTE3QGdtYWlsLmNvbSIsImV4cCI6MTc3MzQ5NDk1NywibmFtZSI6IkhpbG1pIE11c3lhZmEiLCJ1c2VyX2lkIjoxfQ.lcKcpOQPe8Aybg_34l4X7dHiMiNvt5tKR4r12kBXoRo" \
  -H "Content-Type: application/json" \
  -d '{
    "course_id": 1,
    "title": "Variables & Data Types",
    "video_timestamp": "12:45",
    "position": 3
  }'
```
> Note: Param `course_id` diambil dari route param. Body berisi detail chapter.

**Response (200):**
```json
{
  "success": true,
  "message": "[CREATE CHAPTER] Chapter created successfully",
  "data": {
    "chapter": {
      "chapter_id": 3,
      "course_id": 1,
      "title": "Variables & Data Types",
      "video_timestamp": "12:45",
      "position": 3,
      "is_completed": false,
      "created_at": "2026-03-13T12:00:00Z",
      "updated_at": "2026-03-13T12:00:00Z"
    }
  }
}
```

### Update Chapter
```bash
curl -X PUT http://localhost:3344/chapters/edit/3 \
  -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImhpbG1pbXVzeWFmYTE3QGdtYWlsLmNvbSIsImV4cCI6MTc3MzQ5NDk1NywibmFtZSI6IkhpbG1pIE11c3lhZmEiLCJ1c2VyX2lkIjoxfQ.lcKcpOQPe8Aybg_34l4X7dHiMiNvt5tKR4r12kBXoRo" \
  -H "Content-Type: application/json" \
  -d '{
    "title": "Variables, Constants & Data Types",
    "is_completed": true
  }'
```
**Response (200):**
```json
{
  "success": true,
  "message": "[UPDATE CHAPTER] Chapter updated successfully"
}
```

### Delete Chapter
```bash
curl -X DELETE http://localhost:3344/chapters/delete/3 \
  -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImhpbG1pbXVzeWFmYTE3QGdtYWlsLmNvbSIsImV4cCI6MTc3MzQ5NDk1NywibmFtZSI6IkhpbG1pIE11c3lhZmEiLCJ1c2VyX2lkIjoxfQ.lcKcpOQPe8Aybg_34l4X7dHiMiNvt5tKR4r12kBXoRo"
```
**Response (200):**
```json
{
  "success": true,
  "message": "[DELETE CHAPTER] Chapter deleted successfully"
}
```

---

## 4. Explore Courses (Public/Shared Courses)

### Get All Available Courses
```bash
curl -X GET http://localhost:3344/explore/courses \
  -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImhpbG1pbXVzeWFmYTE3QGdtYWlsLmNvbSIsImV4cCI6MTc3MzQ5NDk1NywibmFtZSI6IkhpbG1pIE11c3lhZmEiLCJ1c2VyX2lkIjoxfQ.lcKcpOQPe8Aybg_34l4X7dHiMiNvt5tKR4r12kBXoRo"
```
**Response (200):**
```json
{
  "success": true,
  "message": "[GET ALL COURSES FOR EXPLORE] Courses fetched successfully",
  "data": {
    "courses": [
      {
        "explore_course_id": 1,
        "creator_id": 5,
        "title": "Full Stack Web Development",
        "description": "Complete web dev course from HTML to deployment",
        "channel": "TechMaster",
        "channel_link": "https://youtube.com/@TechMaster",
        "video_link": "https://youtube.com/watch?v=xyz789",
        "thumbnail_image_url": "https://i.ytimg.com/vi/xyz789/maxresdefault.jpg",
        "version": 1,
        "created_at": "2026-03-10T08:00:00Z",
        "updated_at": "2026-03-10T08:00:00Z"
      }
    ]
  }
}
```

### Get Explore Course by ID
```bash
curl -X GET http://localhost:3344/explore/courses/1 \
  -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImhpbG1pbXVzeWFmYTE3QGdtYWlsLmNvbSIsImV4cCI6MTc3MzQ5NDk1NywibmFtZSI6IkhpbG1pIE11c3lhZmEiLCJ1c2VyX2lkIjoxfQ.lcKcpOQPe8Aybg_34l4X7dHiMiNvt5tKR4r12kBXoRo"
```
**Response (200):**
```json
{
  "success": true,
  "message": "[GET EXPLORED COURSE BY ID] Course fetched successfully",
  "data": {
    "course": {
      "explore_course_id": 1,
      "creator_id": 5,
      "title": "Full Stack Web Development",
      "description": "Complete web dev course from HTML to deployment",
      "channel": "TechMaster",
      "channel_link": "https://youtube.com/@TechMaster",
      "video_link": "https://youtube.com/watch?v=xyz789",
      "thumbnail_image_url": "https://i.ytimg.com/vi/xyz789/maxresdefault.jpg",
      "version": 1,
      "created_at": "2026-03-10T08:00:00Z",
      "updated_at": "2026-03-10T08:00:00Z"
    }
  }
}
```

### Search Courses
```bash
curl -X GET "http://localhost:3344/explore/courses/search?q=web" \
  -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImhpbG1pbXVzeWFmYTE3QGdtYWlsLmNvbSIsImV4cCI6MTc3MzQ5NDk1NywibmFtZSI6IkhpbG1pIE11c3lhZmEiLCJ1c2VyX2lkIjoxfQ.lcKcpOQPe8Aybg_34l4X7dHiMiNvt5tKR4r12kBXoRo"
```
**Response (200):**
```json
{
  "success": true,
  "message": "[SEARCH COURSES] Courses fetched successfully",
  "data": {
    "courses": [
      {
        "explore_course_id": 1,
        "creator_id": 5,
        "title": "Full Stack Web Development",
        "description": "Complete web dev course from HTML to deployment",
        "version": 1,
        "created_at": "2026-03-10T08:00:00Z",
        "updated_at": "2026-03-10T08:00:00Z"
      }
    ]
  }
}
```

### Enroll in Course
```bash
curl -X POST http://localhost:3344/explore/courses/1/enroll \
  -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImhpbG1pbXVzeWFmYTE3QGdtYWlsLmNvbSIsImV4cCI6MTc3MzQ5NDk1NywibmFtZSI6IkhpbG1pIE11c3lhZmEiLCJ1c2VyX2lkIjoxfQ.lcKcpOQPe8Aybg_34l4X7dHiMiNvt5tKR4r12kBXoRo"
```
> Akan meng-copy explore course + chapters-nya ke personal courses user.

**Response (200):**
```json
{
  "success": true,
  "message": "[ENROLL IN COURSE] Enrolled in course successfully"
}
```

---

## 5. Explore Chapters (Chapters dari Explore Course)

### Get All Chapters by Explore Course ID
```bash
curl -X GET http://localhost:3344/explore/courses/1/chapters/ \
  -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImhpbG1pbXVzeWFmYTE3QGdtYWlsLmNvbSIsImV4cCI6MTc3MzQ5NDk1NywibmFtZSI6IkhpbG1pIE11c3lhZmEiLCJ1c2VyX2lkIjoxfQ.lcKcpOQPe8Aybg_34l4X7dHiMiNvt5tKR4r12kBXoRo"
```
**Response (200):**
```json
{
  "success": true,
  "message": "[GET ALL EXPLORE CHAPTERS BY COURSE ID] Chapters fetched successfully",
  "data": {
    "chapters": [
      {
        "explore_chapter_id": 1,
        "explore_course_id": 1,
        "title": "HTML Basics",
        "description": "Learn the fundamentals of HTML",
        "order": 1,
        "created_at": "2026-03-10T08:00:00Z",
        "updated_at": "2026-03-10T08:00:00Z"
      },
      {
        "explore_chapter_id": 2,
        "explore_course_id": 1,
        "title": "CSS Styling",
        "description": "Master CSS layouts and styling",
        "order": 2,
        "created_at": "2026-03-10T08:00:00Z",
        "updated_at": "2026-03-10T08:00:00Z"
      }
    ]
  }
}
```

### Get One Explore Chapter by ID
```bash
curl -X GET http://localhost:3344/explore/courses/1/chapters/1 \
  -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImhpbG1pbXVzeWFmYTE3QGdtYWlsLmNvbSIsImV4cCI6MTc3MzQ5NDk1NywibmFtZSI6IkhpbG1pIE11c3lhZmEiLCJ1c2VyX2lkIjoxfQ.lcKcpOQPe8Aybg_34l4X7dHiMiNvt5tKR4r12kBXoRo"
```
**Response (200):**
```json
{
  "success": true,
  "message": "[GET ONE EXPLORE CHAPTER BY ID] Chapter fetched successfully",
  "data": {
    "chapter": {
      "explore_chapter_id": 1,
      "explore_course_id": 1,
      "title": "HTML Basics",
      "description": "Learn the fundamentals of HTML",
      "order": 1,
      "created_at": "2026-03-10T08:00:00Z",
      "updated_at": "2026-03-10T08:00:00Z"
    }
  }
}
```

### Create Explore Chapter
```bash
curl -X POST http://localhost:3344/explore/courses/1/chapters/create \
  -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImhpbG1pbXVzeWFmYTE3QGdtYWlsLmNvbSIsImV4cCI6MTc3MzQ5NDk1NywibmFtZSI6IkhpbG1pIE11c3lhZmEiLCJ1c2VyX2lkIjoxfQ.lcKcpOQPe8Aybg_34l4X7dHiMiNvt5tKR4r12kBXoRo" \
  -H "Content-Type: application/json" \
  -d '{
    "title": "JavaScript Fundamentals",
    "description": "Learn the basics of JavaScript programming",
    "order": 3
  }'
```
**Response (200):**
```json
{
  "success": true,
  "message": "[CREATE EXPLORE CHAPTER] Chapter created successfully",
  "data": {
    "chapter": {
      "explore_chapter_id": 3,
      "explore_course_id": 1,
      "title": "JavaScript Fundamentals",
      "description": "Learn the basics of JavaScript programming",
      "order": 3,
      "created_at": "2026-03-13T12:00:00Z",
      "updated_at": "2026-03-13T12:00:00Z"
    }
  }
}
```

### Update Explore Chapter
```bash
curl -X PUT http://localhost:3344/explore/courses/1/chapters/edit/3 \
  -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImhpbG1pbXVzeWFmYTE3QGdtYWlsLmNvbSIsImV4cCI6MTc3MzQ5NDk1NywibmFtZSI6IkhpbG1pIE11c3lhZmEiLCJ1c2VyX2lkIjoxfQ.lcKcpOQPe8Aybg_34l4X7dHiMiNvt5tKR4r12kBXoRo" \
  -H "Content-Type: application/json" \
  -d '{
    "title": "JavaScript ES6+ Fundamentals",
    "description": "Modern JavaScript with ES6 features"
  }'
```
**Response (200):**
```json
{
  "success": true,
  "message": "[UPDATE EXPLORE CHAPTER] Chapter updated successfully"
}
```

### Delete Explore Chapter
```bash
curl -X DELETE http://localhost:3344/explore/courses/1/chapters/delete/3 \
  -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImhpbG1pbXVzeWFmYTE3QGdtYWlsLmNvbSIsImV4cCI6MTc3MzQ5NDk1NywibmFtZSI6IkhpbG1pIE11c3lhZmEiLCJ1c2VyX2lkIjoxfQ.lcKcpOQPe8Aybg_34l4X7dHiMiNvt5tKR4r12kBXoRo"
```
**Response (200):**
```json
{
  "success": true,
  "message": "[DELETE EXPLORE CHAPTER] Chapter deleted successfully"
}
```

---

## Error Response Format

Semua error mengikuti format yang sama:

```json
{
  "success": false,
  "message": "[CONTEXT] Error description"
}
```

**Common HTTP Status Codes:**
| Code | Meaning |
|------|---------|
| `400` | Bad Request — data tidak valid atau parameter salah |
| `401` | Unauthorized — token tidak ada, expired, atau format salah |
| `403` | Forbidden — user tidak punya akses ke resource ini |
| `404` | Not Found — resource tidak ditemukan |
| `500` | Internal Server Error — error di server |

---

## Quick Start

```bash
# 1. Register
TOKEN=$(curl -s -X POST http://localhost:3344/auth/register \
  -H "Content-Type: application/json" \
  -d '{"full_name":"Test User","email":"test@mail.com","password":"pass123"}' \
  | jq -r '.data.token')

# 2. Create a course
curl -X POST http://localhost:3344/courses/create \
  -H "Authorization: Bearer $TOKEN" \
  -H "Content-Type: application/json" \
  -d '{"video_link":"https://youtube.com/watch?v=example"}'

# 3. Browse public courses
curl -X GET http://localhost:3344/explore/courses \
  -H "Authorization: Bearer $TOKEN"

# 4. Enroll in a public course
curl -X POST http://localhost:3344/explore/courses/1/enroll \
  -H "Authorization: Bearer $TOKEN"

# 5. View your courses
curl -X GET http://localhost:3344/courses/ \
  -H "Authorization: Bearer $TOKEN"
```
