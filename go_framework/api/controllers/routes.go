package controllers

import "github.com/ethanol1310/go-forum/api/middlewares"

func (s *Server) initializeRoutes() {

	v1 := s.Router.Group("/api/v1")
	{
		// Captcha Route
		v1.GET("/captcha", s.InitCaptcha)
		v1.GET("/captcha/:id", s.GetCaptchaImage)

		// Login Route
		v1.POST("/login", s.Login)

		// Reset password:
		v1.POST("/password/forgot", s.ForgotPassword)
		v1.POST("/password/reset", s.ResetPassword)

		//Users routes
		v1.POST("/users", s.CreateUser)
		v1.GET("/users", s.GetUsers)
		v1.GET("/users/:id", s.GetUser)
		v1.PUT("/users/:id", middlewares.TokenAuthMiddleware(), s.UpdateUser)
		v1.PUT("/avatar/users/:id", middlewares.TokenAuthMiddleware(), s.UpdateAvatar)
		v1.DELETE("/users/:id", middlewares.TokenAuthMiddleware(), s.DeleteUser)

		//Posts routes
		v1.POST("/posts", middlewares.TokenAuthMiddleware(), s.CreatePost)
		v1.GET("/posts", s.GetPosts)
		v1.GET("/posts/:id", s.GetPost)
		v1.PUT("/posts/:id", middlewares.TokenAuthMiddleware(), s.UpdatePost)
		v1.DELETE("/posts/:id", middlewares.TokenAuthMiddleware(), s.DeletePost)
		v1.GET("/user_posts/:id", s.GetUserPosts)

		//Like route
		v1.GET("/likes/:id", s.GetLikes)
		v1.POST("/likes/:id", middlewares.TokenAuthMiddleware(), s.LikePost)
		v1.DELETE("/likes/:id", middlewares.TokenAuthMiddleware(), s.UnLikePost)

		//Comment routes
		v1.POST("/comments/:id", middlewares.TokenAuthMiddleware(), s.CreateComment)
		v1.GET("/comments/:id", s.GetComments)
		v1.PUT("/comments/:id", middlewares.TokenAuthMiddleware(), s.UpdateComment)
		v1.DELETE("/comments/:id", middlewares.TokenAuthMiddleware(), s.DeleteComment)
	}
}
