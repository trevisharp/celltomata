package services

import (
	"os"
)

type ValidateAccountService struct {
	EmailService EmailService
}

func (s ValidateAccountService) SendEmail(username, email string) error {
	origin := os.Getenv("EMAIL_ORIGIN")
	password := os.Getenv("EMAIL_PASSWORD")
	title := "Welcome to celltomata"
	validationUrl := os.Getenv("VALIDATION")
	message := `
	<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Transitional//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd">
	<meta http-equiv="Content-Type" content="text/html; charset=UTF-8" />
	<html>
	<head></head>
	<body>
		<table cellpadding="0" cellspacing="0" border="0" style="width: 600px; height: 400px; padding: 20px;">
		<tr>
			<td>
			<table cellpadding="0" cellspacing="0" border="0" width="100%" style="
				border-radius: 10px;
				background-color: black;
				color: white;
				font-family: 'Segoe UI', Tahoma, Geneva, Verdana, sans-serif;
				padding: 40px;
				text-align: center;
			">
				<tr>
				<td style="padding-bottom: 30px;">
					<svg width="400" height="100" viewBox="0 0 400 100" xmlns="http://www.w3.org/2000/svg">
					<defs>
						<linearGradient id="grad" x1="0" y1="0" x2="1" y2="1">
						<stop offset="0%" stop-color="#28a745"/>
						<stop offset="100%" stop-color="#20c997"/>
						</linearGradient>
					</defs>
					<text x="50%" y="50%" dominant-baseline="middle" text-anchor="middle"
						font-family="Segoe UI, Roboto, sans-serif"
						font-size="42" fill="url(#grad)" letter-spacing="2">
						celltomata
					</text>
					</svg>
				</td>
				</tr>
				<tr>
				<td style="padding-bottom: 20px;">
					<h3 style="margin: 0;">Welcome ` + username + ` to celltomata!</h3>
				</td>
				</tr>
				<tr>
				<td>
					Your account has been created successfully, click the 
					<a href="` + validationUrl + `" style="color: orange;">link</a>
					below to verify your email.
				</td>
				</tr>
			</table>
			</td>
		</tr>
		</table>
	</body>
	</html>
	`

	return s.EmailService.Send(origin, password, email, title, message)
}
