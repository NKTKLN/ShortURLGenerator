package user

import (
	"fmt"
	"net/mail"
	"net/smtp"
	"os"

	"github.com/NKTKLN/ShortURLGenerator/pkg/others"
	"github.com/joho/godotenv"
	"github.com/scorredoira/email"
)

var htmlEmailVerification string = `
	<!DOCTYPE html>
	<html>
		<head>
			<style>
				body {font-family:arial,sans-serif!important;}
				.linkButton {width:150px;height:50px;border: 2px solid black;background:none;border-radius:10px;margin:40px;}
				.buttonText {font-size:20px;color:black;margin-top:10px;}
			</style>
		</head>
		<body>
			<div align="center">
				<p style="font-size:25px;font-weight:bold;">Verify your email address</p>
				<p>To confirm your login, just click the <b>verify email</b> button below:</p>
				<a href="https://%s/verification/%s" style="text-decoration:none;"><div class="linkButton"><p class="buttonText">Verify email</p></div></a>
				<p>If you haven't started registering on my site, just ignore this message.</p>
				<hr style="width:500px;margin:40px;">
				<p>2021 © | Created with ❤️ by <a href="https://nktkln.com" style="color:black;">NKTKLN</a>
			</div>
		</body>
	</html>
`

var htmlEmailReset string = `
	<!DOCTYPE html>
	<html>
		<head>
			<style>
				body {font-family:arial,sans-serif!important;}
				.linkButton {width:150px;height:50px;border: 2px solid black;background:none;border-radius:10px;margin:40px;}
				.buttonText {font-size:20px;color:black;margin-top:10px;}
			</style>
		</head>
		<body>
			<div align="center">
				<p style="font-size:25px;font-weight:bold;">Reset your password</p>
				<p>To confirm reset password, just click the <b>reset</b> button below:</p>
				<a href="https://%s/reset/%s" style="text-decoration:none;"><div class="linkButton"><p class="buttonText">Reset</p></div></a>
				<p>If you don't want to reset your password, just ignore this message.</p>
				<hr style="width:500px;margin:40px;">
				<p>2021 © | Created with ❤️ by <a href="https://nktkln.com" style="color:black;">NKTKLN</a>
			</div>
		</body>
	</html>
`

var htmlEmailTelegram string = `
	<!DOCTYPE html>
	<html>
		<head>
			<style>
				body {font-family:arial,sans-serif!important;}
				.linkButton {width:150px;height:50px;border: 2px solid black;background:none;border-radius:10px;margin:40px;}
				.buttonText {font-size:20px;color:black;margin-top:10px;}
			</style>
		</head>
		<body>
			<div align="center">
				<p style="font-size:25px;font-weight:bold;">Verify your telegram</p>
				<p>To confirm your login, just click the <b>verify telegram</b> button below:</p>
				<a href="https://t.me/%s?start=%s" style="text-decoration:none;"><div class="linkButton"><p class="buttonText">Verify telegram</p></div></a>
				<p>If you haven't started registering on my site, just ignore this message.</p>
				<hr style="width:500px;margin:40px;">
				<p>2021 © | Created with ❤️ by <a href="https://nktkln.com" style="color:black;">NKTKLN</a>
			</div>
		</body>
	</html>
`

func EmailSend(recipientEmail, hash, text string) {
	// Obtaining tedious information to send a message
	var emailText, site string
	if text == "verification" {
		others.ErrorChecking(godotenv.Load("env/app.env"))
		site = os.Getenv("VIRTUAL_HOST")
		emailText = htmlEmailVerification
	} else if text == "reset" {
		others.ErrorChecking(godotenv.Load("env/app.env"))
		site = os.Getenv("VIRTUAL_HOST")
		emailText = htmlEmailReset
	} else if text == "telegram" {
		others.ErrorChecking(godotenv.Load("env/bot.env"))
		site = os.Getenv("BOT_USERNAME")
		emailText = htmlEmailTelegram
	} 
	others.ErrorChecking(godotenv.Load("env/email.env"))
	senderEmail, password, smtpAddres := 
		os.Getenv("SENDER_EMAIL"),
		os.Getenv("EMAIL_PASSWORD"),
		os.Getenv("SMTP_ADDRESS")
	// Login and sending a message to the mail
	m := email.NewHTMLMessage("Login confirmation ", fmt.Sprintf(emailText, site, hash))
	m.From = mail.Address{Name: "ShortURLGenerator", Address: senderEmail}
	m.To = []string{recipientEmail}
	auth := smtp.PlainAuth("", senderEmail, password, smtpAddres)
	others.ErrorChecking(email.Send(smtpAddres+":587", auth, m))
}
