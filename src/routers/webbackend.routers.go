package routers

import "github.com/gin-gonic/gin"

func SetWebBackendRouters(router *gin.RouterGroup) {

	router.GET("/",func(c *gin.Context) {
		c.HTML(200, "sys_dashboard.html", gin.H{
			"Title": "Dashboard",
			"Breadcrumb": []string{
				"Dashboard", 
			},
		})
	})
	router.GET("/app-portal",func(c *gin.Context) {
		c.HTML(200, "sys_portal.html", gin.H{
			"Title": "Portal Aplikasi",
			"Breadcrumb": []string{
				"Dashboard",
				"Portal Aplikasi",
			},
		})
	})
	router.GET("/menu-portal",func(c *gin.Context) {
		c.HTML(200, "sys_menu.html", gin.H{
			"Title": "Portal Menu",
			"Breadcrumb": []string{
				"Dashboard",
				"Portal Menu",
			},
		})
	})
	router.GET("/group",func(c *gin.Context) {
		c.HTML(200, "sys_group.html", gin.H{
			"Title": "Group Pengguna",
			"Breadcrumb": []string{
				"Dashboard",
				"Group Pengguna",
			},
		})
	})
	router.GET("/users",func(c *gin.Context) {
		c.HTML(200, "sys_users.html", gin.H{
			"Title": "Pengguna",
			"Breadcrumb": []string{
				"Dashboard",
				"Pengguna",
			},
		})
	})
	router.GET("/permission",func(c *gin.Context) {
		c.HTML(200, "sys_permission.html", gin.H{
			"Title": "Hak Akses",
			"Breadcrumb": []string{
				"Dashboard",
				"Hak Akses",
			},
		})
	})
	router.GET("/log-auth",func(c *gin.Context) {
		c.HTML(200, "sys_log_auth.html", gin.H{
			"Title": "Login History",
			"Breadcrumb": []string{
				"Dashboard",
				"Login History",
			},
		})
	})
	router.GET("/settings",func(c *gin.Context) {
		c.HTML(200, "sys_settings.html", gin.H{
			"Title": "Pengaturan",
			"Breadcrumb": []string{
				"Dashboard",
				"Pengaturan",
			},
		})
	})
}