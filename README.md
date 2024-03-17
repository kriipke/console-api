# GitOps Console

## Overview

Secrets and other configuration for the `console-api` microserice should be located at:

	configs/app.env

You can create this file by executing

	cp ./configs/app.env.example ./configs/app.env


Afterwards, edit `./configs/app.env` to contain:

* credentials to a valid Postgres 14 database
* SMTP credentials to a Mailtrap Inbox

## API Overview

	METHOD  METHOD	                                 ROUTE                                                                                             DESCRIPTION
	----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------

	GET     /api/healthchecker                       main.main.func1
	POST    /api/auth/register                       github.com/kriipke/console-api/pkg/controllers.(*AuthController).SignUpUser-fm                     Create a new user
	POST    /api/auth/login                          github.com/kriipke/console-api/pkg/controllers.(*AuthController).SignInUser-fm                     Sign in the user
	GET     /api/auth/refresh                        github.com/kriipke/console-api/pkg/controllers.(*AuthController).RefreshAccessToken-fm             Refresh the access token
	GET     /api/auth/logout                         github.com/kriipke/console-api/pkg/controllers.(*AuthController).LogoutUser-fm                     Logout user
	POST    /api/auth/forgotpassword                 github.com/kriipke/console-api/pkg/controllers.(*AuthController).ForgotPassword-fm                 To request a rest link
	PATCH   /api/auth/resetpassword/:resetToken      github.com/kriipke/console-api/pkg/controllers.(*AuthController).ResetPassword-fm                  To reset the password
	GET     /api/auth/verifyemail/:verificationCode  github.com/kriipke/console-api/pkg/controllers.(*AuthController).VerifyEmail-fm                    Verify Email address w/ verification code

	GET     /api/users/me                            github.com/kriipke/console-api/pkg/controllers.(*UserController).GetMe-fm                          Return whoami info

	POST    /api/posts/                              github.com/kriipke/console-api/pkg/controllers.(*PostController).CreatePost-fm                     Create new post
	GET     /api/posts/                              github.com/kriipke/console-api/pkg/controllers.(*PostController).FindPosts-fm                      Get posts
	PUT     /api/posts/:postId                       github.com/kriipke/console-api/pkg/controllers.(*PostController).UpdatePost-fm                     Updtae post
	GET     /api/posts/:postId                       github.com/kriipke/console-api/pkg/controllers.(*PostController).FindPostById-fm                   Get post by ID
	DELETE  /api/posts/:postId                       github.com/kriipke/console-api/pkg/controllers.(*PostController).DeletePost-fm                     Delete Post

	POST    /api/clusters/                           github.com/kriipke/console-api/pkg/controllers.(*ClusterController).CreateCluster-fm               Add new cluster to DB
	GET     /api/clusters/                           github.com/kriipke/console-api/pkg/controllers.(*ClusterController).FindClusters-fm                Fetch clusters
	PUT     /api/clusters/:clusterId                 github.com/kriipke/console-api/pkg/controllers.(*ClusterController).UpdateCluster-fm               Update cluster info
	GET     /api/clusters/:clusterId                 github.com/kriipke/console-api/pkg/controllers.(*ClusterController).FindClusterById-fm             Get cluster info by ID
	DELETE  /api/clusters/:clusterId                 github.com/kriipke/console-api/pkg/controllers.(*ClusterController).DeleteCluster-fm               Delete Cluster


* Project Structure
 - https://github.com/golang-standards/project-layout

## References

### CORS

* https://fetch.spec.whatwg.org/#http-cors-protocol

## Backend
https://github.com/wpcodevo/golang-gorm-postgres

. https://codevoweb.com/setup-golang-gorm-restful-api-project-with-postgres
2. https://codevoweb.com/golang-gorm-postgresql-user-registration-with-refresh-tokens
3. https://codevoweb.com/golang-and-gorm-user-registration-email-verification
4. https://codevoweb.com/forgot-reset-passwords-in-golang-with-html-email
5. https://codevoweb.com/build-restful-crud-api-with-golang

## Frontend

https://github.com/wpcodevo/react-query-axios-tailwindcss
https://github.com/wpcodevo/node_typeorm

## Hostname Definitions



== `./pkg/controllers/auth.controller.go`

	ctx.SetCookie("access_token", access_token, config.AccessTokenMaxAge*60, "/", "localhost", false, true)
	ctx.SetCookie("refresh_token", refresh_token, config.RefreshTokenMaxAge*60, "/", "localhost", false, true)
	ctx.SetCookie("logged_in", "true", config.AccessTokenMaxAge*60, "/", "localhost", false, false)


== Turn on/off email verification


=== `./pkg/controllers/auth.controller.go`

	now := time.Now()
	newUser := models.User{
		Name:      payload.Name,
		Email:     strings.ToLower(payload.Email),
		Password:  hashedPassword,
		Role:      "user",
		Verified:  true, .   <----- This set to True disables verification
		Photo:     payload.Photo,
		Provider:  "local",
		CreatedAt: now,
		UpdatedAt: now,
	}

=== if `false` 
=== if `true` 

SMTP creds will be sourced from `configs/app.env`. See `configs/app.env.example`.
