package seeder

import (
	"log"
	"strings"

	"zurihaqi.github.io-backend/internal/database"
	"zurihaqi.github.io-backend/internal/model"
)

func SeedProjects() {
	db := database.DB

	if err := db.Exec("DELETE FROM projects").Error; err != nil {
		log.Fatal("Failed to delete existing projects:", err)
	}
	if err := db.Exec("DELETE FROM sqlite_sequence WHERE name = 'projects'").Error; err != nil {
		log.Fatal("Failed to reset projects ID sequence:", err)
	}

	projectsData := []struct {
		Title        string
		Description  string
		Repo         string
		Live         string
		Categories   []string
		Technologies []string
	}{
		{
			Title:        "Spatial AR",
			Description:  "This application is a Kotlin-based app utilizing ARCore technology to help users accurately measure room dimensions and place 3D furniture model directly in a virtual space. With the AR measurement feature, users can easily determine the length, width, and height of a room simply by pointing their device's camera. Additionally, the app provides a selection of 3D furniture model that can be placed in real-time, allowing users to visualize how the furniture would look in their space before making a purchase or rearranging their interior. Designed to offer an interactive and realistic experience, this app leverages augmented reality to enhance interior planning and design.",
			Repo:         "https://github.com/Zurihaqi/SpatialAR",
			Categories:   []string{"Mobile App"},
			Technologies: []string{"Kotlin", "ARCore", "Jetpack Compose"},
		},
		{
			Title:        "NestJs Chat App",
			Description:  "This is a real-time chat application built using NestJS, WebSockets, and MongoDB. It supports user registration, login, and the ability to send messages to specific users in a chat system.",
			Repo:         "https://github.com/Zurihaqi/nestjs-ws-mongodb",
			Categories:   []string{"API", "Website"},
			Technologies: []string{"NestJs", "NestJs WebSockets", "MongoDB", "Mongoose", "TypeScript"},
		},
		{
			Title:        "VTTP (Validated Tech Talent Pool)",
			Description:  "This web-based application is designed for investors and tech leads looking for developers or engineers to work on their projects. Through this platform, clients can easily search, filter, and select developers based on their needs, whether for hiring as part of their team or assigning specific projects. Key features include developer search by skills, detailed developer profiles, and a project management system that allows clients to directly assign tasks to their chosen developers. This application aims to streamline the recruitment process and project distribution within the tech ecosystem.",
			Categories:   []string{"Website"},
			Technologies: []string{"Spring Boot", "React", "Tailwind", "NextUI", "PostgreSQL"},
		},
		{
			Title:        "Ichiban Kuji",
			Description:  "Ichiban Kuji is a Kotlin-based application that allows users to collect action figures of their favorite anime characters. The app features a virtual coin system that can be used to either exchange for specific action figures or participate in a gacha system for a chance to obtain random figures, including rare and limited editions. Users can acquire these coins through in-app purchases using real money, enabling them to expand their collection more quickly. With an engaging interfaces and secure transaction system, CollectAnime provides a fun and immersive experience for anime figure collectors.",
			Categories:   []string{"Mobile App"},
			Technologies: []string{"Kotlin"},
		},
		{
			Title:        "Kerjain Aja",
			Description:  "'Kerjain Aja' is a platform designed to connect individuals who need help with daily tasks to skilled workers looking for flexible job opportunities. Customers can quickly find reliable assistance for service like home repairs, cleaning, and personal tasks, while workers gain access to jobs that fit their skills and schedules, promoting work-life balance and financial stability. \nBeyond convenience, 'Kerjain Aja' has a significant social impact by helping reduce unemployment, especially in the informal sector, and empowering workers with meaningful job opportunities. By bridging the gap between demand and supply, the platform fosters economic growth and accessible employment.",
			Repo:         "https://github.com/AlanPratama/palu-gada-mobile",
			Categories:   []string{"Website", "Mobile App"},
			Technologies: []string{"Spring Boot", "React Native", "Midtrans", "React", "Tailwind", "NextUI", "Docker", "PostgreSQL"},
		},
		{
			Title:        "Quran Verses and Calculator App",
			Description:  "This web app combines a sleek, functional calculator with a spiritually uplifting Qur'an Verse Randomizer, offering users both practical utility and moments of reflection; the calculator provides standard arithmetic operations with a clean, responsive interfaces, while the verse section allows users to browse Qur'anic ayahs with options to view a random verse, navigate to the next or previous one, and listen to the recitation—all in a minimal design that promotes focus and ease of use.",
			Repo:         "https://github.com/Zurihaqi/react-calculator-and-verse-randomizer",
			Live:         "https://react-calculator-and-verse-randomizer.vercel.app/",
			Categories:   []string{"Website"},
			Technologies: []string{"React", "Tailwind", "Axios"},
		},
		{
			Title:        "SecondHand",
			Description:  "This platform is a place to buy and sell goods online, especially used goods. This platform opens and provides various types of categories of needs. Users who register themselves on this application can act as sellers and buyers by using the same 1 (one) account. account. This platform will bring together sellers and buyers to be able to negotiate goods and make transactions directly outside the platform.",
			Repo:         "https://github.com/Zurihaqi/second-hand",
			Categories:   []string{"Website"},
			Technologies: []string{"React", "Bootstrap", "Axios", "jQuery"},
		},
		{
			Title:        "Kedai Mie Ayam",
			Description:  "This was my first attempt to create a project in TypeScript. This website is my own version of Colt Steele's YelpCamp, where users can post kedai and review others' kedai. The website was built using Next.js and Tailwind for the front-end and back-end, NextAuth for the authentication system, and MySQL for the database.",
			Repo:         "https://github.com/Zurihaqi/kedai-mie-ayam",
			Categories:   []string{"Website"},
			Technologies: []string{"NextJs", "Tailwind", "NextAuth", "PrismaORM", "MySQL", "JWT"},
		},
		{
			Title:        "Sentiment Analysis @ubhara_jaya",
			Description:  "Sentiment analysis is the process of analyzing digital text to determine if the emotional tone of the message is positive, negative, or neutral. In this case, the Data was taken from comments on the Instagram account posts of Universitas Bhayangkara Jakarta Raya using web scraping. The comment period ranges from January 2021 to December 2022.",
			Repo:         "https://colab.research.google.com/drive/16fiKZw49eoZLaWgruQ_B9_0rEc5QB7PQ?usp=sharing",
			Categories:   []string{"Machine Learning"},
			Technologies: []string{"Numpy", "Pandas", "NLTK", "Sastrawi"},
		},
		{
			Title:        "SPK Peminatan Informatika",
			Description:  "The purpose of this web-based system is to assess the suitability of the concentration selection of students of the informatics study program at Bhayangkara Jakarta Raya University with their interests and talents. This system is designed based on web with Express framework and Node.js for the backend, while the frontend uses Vue.js.",
			Repo:         "https://github.com/Zurihaqi/spk-informatika-frontend-skripsi",
			Categories:   []string{"Website"},
			Technologies: []string{"VueJs", "Bootstrap", "Axios", "ExpressJs", "Sequelize", "PostgreSQL"},
		},
		{
			Title:        "YelpCamp",
			Description:  "YelpCamp is a website where users can create and review campgrounds. In order to review or create a campground, you must have an account. This project was part of Colt Steele's web dev course on udemy. This project was created using Node.js, Express, MongoDB, and Bootstrap. Passport.js was used to handle authentication.",
			Repo:         "https://github.com/Zurihaqi/yelpcamp",
			Categories:   []string{"Website"},
			Technologies: []string{"ExpressJs", "EJS", "MongoDB", "Bootstrap", "Passport.js"},
		},
		{
			Title:        "Aplikasi Pemancingan",
			Description:  "This was a project made for a friend. This website was created to make it easier for anglers to find complete information about fishing in Bekasi Regency, and can also make it easier for fishing parties to disseminate information in fishing or become a fishing promotion media.",
			Repo:         "https://github.com/Zurihaqi/pemancingan",
			Categories:   []string{"Website"},
			Technologies: []string{"Laravel", "Eloquent", "MySQL", "Bootstrap"},
		},
		{
			Title:        "Personal Website V2",
			Description:  "An updated version of my personal website built using Next.js, Tailwind, and FramerMotion.",
			Repo:         "https://github.com/Zurihaqi/Zurihaqi.github.io",
			Live:         "https://zurihaqi.github.io/",
			Categories:   []string{"Website"},
			Technologies: []string{"NextJs", "Tailwind", "FramerMotion"},
		},
		{
			Title:        "Personal Website V1",
			Description:  "My first personal website built using Gatsby and Tailwind.",
			Repo:         "https://github.com/Zurihaqi/Zurihaqi.github.io/tree/old-gatsby",
			Categories:   []string{"Website"},
			Technologies: []string{"Gatsby", "React", "Tailwind", "GraphQL"},
		},
	}

	for _, p := range projectsData {
		project := model.Project{
			Title:       p.Title,
			Description: p.Description,
			Repo:        p.Repo,
			Live:        p.Live,
		}

		for _, catName := range p.Categories {
			catName = strings.TrimSpace(catName)
			var cat model.Category
			if err := db.Where("name = ?", catName).First(&cat).Error; err != nil {
				cat = model.Category{Name: catName}
				if err := db.Create(&cat).Error; err != nil {
					log.Fatal("Failed to create category:", err)
				}
			}
			project.Categories = append(project.Categories, cat)
		}

		for _, techName := range p.Technologies {
			techName = strings.TrimSpace(techName)
			var tech model.Technology
			if err := db.Where("name = ?", techName).First(&tech).Error; err != nil {
				tech = model.Technology{Name: techName}
				if err := db.Create(&tech).Error; err != nil {
					log.Fatal("Failed to create technology:", err)
				}
			}
			project.Technologies = append(project.Technologies, tech)
		}

		if err := db.Create(&project).Error; err != nil {
			log.Fatal("Failed to create project:", err)
		}
	}

	log.Println("Projects seeded successfully")
}
