package services

import (
	"fmt"
	"gopkg.in/gomail.v2"
)

func SendEmail(to, body string) error {

	host := "smtp.gmail.com"
	port := 587
	username := "vinicius.devjvm@gmail.com"
	password := "14032005Vi*"

	m := gomail.NewMessage()
	m.SetHeader("From", username)
	m.SetHeader("To", to)
	m.SetHeader("Subject", "Token de recuperacao de senha")
	m.SetBody("text/plain", fmt.Sprintf("Olá, este é o seu token: %w", body))

	d := gomail.NewDialer(host, port, username, password)

	if err := d.DialAndSend(m); err != nil {
		fmt.Printf("Erro ao enviar e-mail: %v\n", err)
		return nil
	}

	fmt.Println("E-mail enviado com sucesso!")

	return nil
}
