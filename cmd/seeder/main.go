package main

import (
	"flag"
	"log"

	"github.com/bobchopperz/bahrululum/internal/config"
	"github.com/bobchopperz/bahrululum/internal/domain/models"
	"github.com/bobchopperz/bahrululum/internal/init/database"
	"gorm.io/gorm"
)

var (
	all      = flag.Bool("all", false, "Seed all tables")
	courses  = flag.Bool("courses", false, "Seed courses only")
	chapters = flag.Bool("chapters", false, "Seed course chapters only")
	contents = flag.Bool("contents", false, "Seed course contents only")
	clean    = flag.Bool("clean", false, "Clean all seeded data before seeding")
)

func main() {
	flag.Parse()

	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Initialize database
	db, err := database.InitDatabase(&cfg.DatabaseConfig)
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	// Auto-migrate models (ensure tables exist)
	if err := db.AutoMigrate(
		&models.Course{},
		&models.CourseChapter{},
		&models.CourseContent{},
	); err != nil {
		log.Fatalf("Failed to migrate models: %v", err)
	}

	// Clean data if requested
	if *clean {
		cleanData(db)
	}

	// Determine what to seed
	if *all {
		seedCourses(db)
		seedChapters(db)
		seedContents(db)
		log.Println("✓ All data seeded successfully!")
	} else if *courses {
		seedCourses(db)
		log.Println("✓ Courses seeded successfully!")
	} else if *chapters {
		seedChapters(db)
		log.Println("✓ Course chapters seeded successfully!")
	} else if *contents {
		seedContents(db)
		log.Println("✓ Course contents seeded successfully!")
	} else {
		// Default: seed everything
		seedCourses(db)
		seedChapters(db)
		seedContents(db)
		log.Println("✓ All data seeded successfully!")
	}
}

func cleanData(db *gorm.DB) {
	log.Println("Cleaning existing seeded data...")

	// Delete in reverse order of dependencies
	if err := db.Exec("DELETE FROM course_contents").Error; err != nil {
		log.Printf("Warning: Failed to clean course_contents: %v", err)
	}

	if err := db.Exec("DELETE FROM course_chapters").Error; err != nil {
		log.Printf("Warning: Failed to clean course_chapters: %v", err)
	}

	if err := db.Exec("DELETE FROM courses").Error; err != nil {
		log.Printf("Warning: Failed to clean courses: %v", err)
	}

	log.Println("✓ Data cleaned successfully!")
}

func seedCourses(db *gorm.DB) {
	log.Println("Seeding courses...")

	coursesData := []models.Course{
		{
			Name:        "Introduction to Go Programming",
			Description: "Learn the fundamentals of Go programming language, from basic syntax to advanced concepts. This comprehensive course covers everything you need to become a proficient Go developer.",
		},
		{
			Name:        "Web Development with Echo Framework",
			Description: "Master building RESTful APIs with Echo, one of the most popular Go web frameworks. Learn routing, middleware, authentication, and best practices.",
		},
		{
			Name:        "Database Design and PostgreSQL",
			Description: "Comprehensive guide to database design principles and PostgreSQL. Learn SQL, normalization, indexing, transactions, and performance optimization.",
		},
	}

	for i := range coursesData {
		// Check if course already exists
		var existing models.Course
		if err := db.Where("name = ?", coursesData[i].Name).First(&existing).Error; err == nil {
			log.Printf("  - Course '%s' already exists, skipping...", coursesData[i].Name)
			continue
		}

		if err := db.Create(&coursesData[i]).Error; err != nil {
			log.Printf("  ✗ Failed to create course '%s': %v", coursesData[i].Name, err)
		} else {
			log.Printf("  ✓ Created course: %s (ID: %d)", coursesData[i].Name, coursesData[i].ID)
		}
	}
}

func seedChapters(db *gorm.DB) {
	log.Println("Seeding course chapters...")

	// Fetch existing courses
	var allCourses []models.Course
	if err := db.Find(&allCourses).Error; err != nil {
		log.Fatalf("Failed to fetch courses: %v", err)
	}

	if len(allCourses) == 0 {
		log.Println("  ! No courses found. Please seed courses first.")
		return
	}

	// Define chapters for each course
	chaptersData := map[string][]models.CourseChapter{
		"Introduction to Go Programming": {
			{
				Title:        "Getting Started with Go",
				Description:  stringPtr("Learn how to install Go, set up your development environment, and write your first Go program."),
				ChapterOrder: 1,
				IsPublished:  true,
			},
			{
				Title:        "Go Basics: Variables and Types",
				Description:  stringPtr("Understanding Go's type system, variables, constants, and basic data types."),
				ChapterOrder: 2,
				IsPublished:  true,
			},
			{
				Title:        "Control Structures and Functions",
				Description:  stringPtr("Master if statements, loops, switch cases, and function definitions in Go."),
				ChapterOrder: 3,
				IsPublished:  true,
			},
			{
				Title:        "Data Structures: Arrays, Slices, and Maps",
				Description:  stringPtr("Deep dive into Go's built-in data structures and their practical applications."),
				ChapterOrder: 4,
				IsPublished:  true,
			},
			{
				Title:        "Structs and Methods",
				Description:  stringPtr("Learn object-oriented programming concepts in Go using structs and methods."),
				ChapterOrder: 5,
				IsPublished:  false,
			},
		},
		"Web Development with Echo Framework": {
			{
				Title:        "Introduction to Echo",
				Description:  stringPtr("Overview of Echo framework, installation, and creating your first API endpoint."),
				ChapterOrder: 1,
				IsPublished:  true,
			},
			{
				Title:        "Routing and Request Handling",
				Description:  stringPtr("Learn about route groups, parameters, query strings, and request binding."),
				ChapterOrder: 2,
				IsPublished:  true,
			},
			{
				Title:        "Middleware and Authentication",
				Description:  stringPtr("Implement middleware for logging, CORS, and JWT-based authentication."),
				ChapterOrder: 3,
				IsPublished:  true,
			},
			{
				Title:        "Database Integration with GORM",
				Description:  stringPtr("Connect your Echo application to PostgreSQL using GORM ORM."),
				ChapterOrder: 4,
				IsPublished:  false,
			},
		},
		"Database Design and PostgreSQL": {
			{
				Title:        "Database Fundamentals",
				Description:  stringPtr("Understanding relational databases, tables, rows, columns, and primary keys."),
				ChapterOrder: 1,
				IsPublished:  true,
			},
			{
				Title:        "SQL Basics: CRUD Operations",
				Description:  stringPtr("Learn to create, read, update, and delete data using SQL."),
				ChapterOrder: 2,
				IsPublished:  true,
			},
			{
				Title:        "Advanced Queries and Joins",
				Description:  stringPtr("Master complex queries, joins, subqueries, and aggregations."),
				ChapterOrder: 3,
				IsPublished:  true,
			},
			{
				Title:        "Indexing and Performance",
				Description:  stringPtr("Learn about database indexing strategies and query optimization."),
				ChapterOrder: 4,
				IsPublished:  false,
			},
		},
	}

	for _, course := range allCourses {
		chapters, exists := chaptersData[course.Name]
		if !exists {
			continue
		}

		for i := range chapters {
			chapters[i].CourseID = course.ID

			// Check if chapter already exists
			var existing models.CourseChapter
			if err := db.Where("course_id = ? AND title = ?", chapters[i].CourseID, chapters[i].Title).First(&existing).Error; err == nil {
				log.Printf("  - Chapter '%s' already exists for course '%s', skipping...", chapters[i].Title, course.Name)
				continue
			}

			if err := db.Create(&chapters[i]).Error; err != nil {
				log.Printf("  ✗ Failed to create chapter '%s': %v", chapters[i].Title, err)
			} else {
				log.Printf("  ✓ Created chapter: %s (ID: %d) for course '%s'", chapters[i].Title, chapters[i].ID, course.Name)
			}
		}
	}
}

func seedContents(db *gorm.DB) {
	log.Println("Seeding course contents...")

	// Fetch all chapters
	var allChapters []models.CourseChapter
	if err := db.Preload("Course").Find(&allChapters).Error; err != nil {
		log.Fatalf("Failed to fetch chapters: %v", err)
	}

	if len(allChapters) == 0 {
		log.Println("  ! No chapters found. Please seed chapters first.")
		return
	}

	// Define contents for specific chapters
	contentsData := map[string][]models.CourseContent{
		"Getting Started with Go": {
			{
				Title:           "Introduction Video",
				Description:     stringPtr("Welcome to the Go programming course! In this video, we'll introduce you to Go and what makes it special."),
				ContentType:     "video",
				FileURL:         stringPtr("https://example.com/videos/go-intro.mp4"),
				ContentOrder:    1,
				IsPublished:     true,
				DurationMinutes: intPtr(15),
			},
			{
				Title:           "Installing Go",
				Description:     stringPtr("Step-by-step guide to installing Go on Windows, macOS, and Linux."),
				ContentType:     "text",
				ContentText:     stringPtr("# Installing Go\n\n## Windows\n1. Download the installer from golang.org\n2. Run the MSI installer\n3. Verify installation with `go version`\n\n## macOS\nUsing Homebrew:\n```bash\nbrew install go\n```\n\n## Linux\n```bash\nsudo apt-get install golang\n```"),
				ContentOrder:    2,
				IsPublished:     true,
				DurationMinutes: intPtr(10),
			},
			{
				Title:           "Your First Go Program",
				Description:     stringPtr("Write and run your first Hello World program in Go."),
				ContentType:     "text",
				ContentText:     stringPtr("# Hello World in Go\n\n```go\npackage main\n\nimport \"fmt\"\n\nfunc main() {\n    fmt.Println(\"Hello, World!\")\n}\n```\n\nRun it with:\n```bash\ngo run main.go\n```"),
				ContentOrder:    3,
				IsPublished:     true,
				DurationMinutes: intPtr(5),
			},
			{
				Title:           "Quiz: Getting Started",
				Description:     stringPtr("Test your knowledge of Go basics."),
				ContentType:     "link",
				FileURL:         stringPtr("https://example.com/quiz/go-basics"),
				ContentOrder:    4,
				IsPublished:     true,
				DurationMinutes: intPtr(10),
			},
		},
		"Go Basics: Variables and Types": {
			{
				Title:           "Variables in Go",
				Description:     stringPtr("Learn about variable declaration and initialization in Go."),
				ContentType:     "video",
				FileURL:         stringPtr("https://example.com/videos/go-variables.mp4"),
				ContentOrder:    1,
				IsPublished:     true,
				DurationMinutes: intPtr(20),
			},
			{
				Title:           "Basic Data Types",
				Description:     stringPtr("Understanding int, float, string, bool, and other basic types."),
				ContentType:     "text",
				ContentText:     stringPtr("# Go Data Types\n\n## Numeric Types\n- int, int8, int16, int32, int64\n- uint, uint8, uint16, uint32, uint64\n- float32, float64\n\n## Other Types\n- string\n- bool\n- byte (alias for uint8)\n- rune (alias for int32)"),
				ContentOrder:    2,
				IsPublished:     true,
				DurationMinutes: intPtr(15),
			},
			{
				Title:           "Type Conversion",
				Description:     stringPtr("Converting between different types in Go."),
				ContentType:     "text",
				ContentText:     stringPtr("# Type Conversion\n\n```go\nvar i int = 42\nvar f float64 = float64(i)\nvar u uint = uint(f)\n```"),
				ContentOrder:    3,
				IsPublished:     true,
				DurationMinutes: intPtr(10),
			},
		},
		"Introduction to Echo": {
			{
				Title:           "What is Echo?",
				Description:     stringPtr("Introduction to Echo framework and its features."),
				ContentType:     "video",
				FileURL:         stringPtr("https://example.com/videos/echo-intro.mp4"),
				ContentOrder:    1,
				IsPublished:     true,
				DurationMinutes: intPtr(12),
			},
			{
				Title:           "Setting Up Echo",
				Description:     stringPtr("Install Echo and create your first server."),
				ContentType:     "text",
				ContentText:     stringPtr("# Setting Up Echo\n\n```bash\ngo get github.com/labstack/echo/v4\n```\n\n```go\npackage main\n\nimport (\n    \"github.com/labstack/echo/v4\"\n)\n\nfunc main() {\n    e := echo.New()\n    e.GET(\"/\", func(c echo.Context) error {\n        return c.String(200, \"Hello, Echo!\")\n    })\n    e.Start(\":8080\")\n}\n```"),
				ContentOrder:    2,
				IsPublished:     true,
				DurationMinutes: intPtr(15),
			},
		},
		"Database Fundamentals": {
			{
				Title:           "What is a Database?",
				Description:     stringPtr("Introduction to databases and PostgreSQL."),
				ContentType:     "video",
				FileURL:         stringPtr("https://example.com/videos/db-intro.mp4"),
				ContentOrder:    1,
				IsPublished:     true,
				DurationMinutes: intPtr(18),
			},
			{
				Title:           "Installing PostgreSQL",
				Description:     stringPtr("How to install and configure PostgreSQL."),
				ContentType:     "text",
				ContentText:     stringPtr("# Installing PostgreSQL\n\n## macOS\n```bash\nbrew install postgresql\nbrew services start postgresql\n```\n\n## Ubuntu\n```bash\nsudo apt-get install postgresql\n```"),
				ContentOrder:    2,
				IsPublished:     true,
				DurationMinutes: intPtr(12),
			},
			{
				Title:           "Database Concepts PDF",
				Description:     stringPtr("Comprehensive guide to database concepts."),
				ContentType:     "pdf",
				FileURL:         stringPtr("https://example.com/docs/database-concepts.pdf"),
				ContentOrder:    3,
				IsPublished:     true,
				DurationMinutes: intPtr(30),
			},
		},
	}

	for _, chapter := range allChapters {
		contents, exists := contentsData[chapter.Title]
		if !exists {
			continue
		}

		for i := range contents {
			contents[i].ChapterID = chapter.ID

			// Check if content already exists
			var existing models.CourseContent
			if err := db.Where("chapter_id = ? AND title = ?", contents[i].ChapterID, contents[i].Title).First(&existing).Error; err == nil {
				log.Printf("  - Content '%s' already exists for chapter '%s', skipping...", contents[i].Title, chapter.Title)
				continue
			}

			if err := db.Create(&contents[i]).Error; err != nil {
				log.Printf("  ✗ Failed to create content '%s': %v", contents[i].Title, err)
			} else {
				log.Printf("  ✓ Created content: %s (ID: %d) for chapter '%s'", contents[i].Title, contents[i].ID, chapter.Title)
			}
		}
	}
}

// Helper functions for pointer values
func stringPtr(s string) *string {
	return &s
}

func intPtr(i int) *int {
	return &i
}
