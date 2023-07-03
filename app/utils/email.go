package utils

import (
	"crypto/rand"
	"crypto/tls"
	"fmt"
	"log"
	"net/smtp"

	"github.com/BisquitDubouche/CargoTrack_auth_microservice/app/models"
)

// Функция генерации кода подтверждения
func GenerateConfirmationCode() string {
	code := make([]byte, 3)
	rand.Read(code)

	// Преобразование кода в строку шестнадцатеричного формата
	return fmt.Sprintf("%x", code)
}

// Функция генерации ссылки для подтверждения почты
func GenerateConfirmationLink(userID string) string {
	baseURL := "http://localhost:8080/shipper/confirm-email" // Замените на ваш базовый URL

	// Формирование ссылки с параметрами userID и confirmationCode
	link := fmt.Sprintf("%s/%s", baseURL, userID)

	return link
}

// Функция генерация ссылки для подтверждения компании
func GenerateConfirmationLinkAdmin(freightCompanyID string) string {
	baseURL := "http://localhost:8080/admin/confirm-register-company" // Замените на ваш базовый URL

	// Формирование ссылки с параметрами userID и confirmationCode
	link := fmt.Sprintf("%s/%s", baseURL, freightCompanyID)

	return link
}

// Отправка письма на почту
func SendConfirmationEmail(email, code, link string) error {
	// Конфигурация SMTP-сервера
	smtpHost := "smtp.gmail.com"
	smtpPort := "465"
	smtpUsername := "switchinglanes.kg@gmail.com"
	smtpPassword := "idgkpbtidyuearpe"

	// Генерация ссылки для подтверждения

	// Формирование письма
	from := "CARGO_TRACK"
	to := []string{email}
	subject := "Сonfirm email"
	body := fmt.Sprintf("You have registered with the service Cargo Track \n\n Please, follow the link to verify your account \n\n%s", link)

	message := []byte(fmt.Sprintf("From: %s\r\n"+
		"To: %s\r\n"+
		"Subject: %s\r\n"+
		"\r\n"+
		"%s\r\n", from, to[0], subject, body))

	// Установка конфигурации для TLS
	tlsConfig := &tls.Config{
		InsecureSkipVerify: true, // Если сертификат не валидный
		ServerName:         smtpHost,
	}

	// Подключение к SMTP-серверу с использованием TLS
	conn, err := tls.Dial("tcp", smtpHost+":"+smtpPort, tlsConfig)
	if err != nil {
		log.Println("Ошибка при подключении к SMTP-серверу:", err)
		return err
	}
	defer conn.Close()

	// Аутентификация на SMTP-сервере
	auth := smtp.PlainAuth("", smtpUsername, smtpPassword, smtpHost)

	// Отправка письма через подключенное соединение
	client, err := smtp.NewClient(conn, smtpHost)
	if err != nil {
		log.Println("Ошибка при создании клиента SMTP:", err)
		return err
	}
	defer client.Quit()

	if err = client.Auth(auth); err != nil {
		log.Println("Ошибка при аутентификации на SMTP-сервере:", err)
		return err
	}

	if err = client.Mail(from); err != nil {
		log.Println("Ошибка при указании отправителя письма:", err)
		return err
	}

	for _, addr := range to {
		if err = client.Rcpt(addr); err != nil {
			log.Println("Ошибка при указании получателя письма:", err)
			return err
		}
	}

	w, err := client.Data()
	if err != nil {
		log.Println("Ошибка при отправке письма:", err)
		return err
	}
	defer w.Close()

	_, err = w.Write(message)
	if err != nil {
		log.Println("Ошибка при записи сообщения письма:", err)
		return err
	}

	return nil
}

// Отправка письма на почту
func SendConfirmationEmailToAdmins(admins []models.Administrator, link string, freithCompany models.FreightCompany) error {
	// Конфигурация SMTP-сервера
	smtpHost := "smtp.gmail.com"
	smtpPort := "465"
	smtpUsername := "switchinglanes.kg@gmail.com"
	smtpPassword := "idgkpbtidyuearpe"

	// Формирование письма
	from := smtpUsername
	subject := "Подтверждение регистрации"
	body := fmt.Sprintf("Эта компания хочет зарегистрироваться на нашей платформе: %s\n\nПодтвердить регистрацию компании: %s", freithCompany.Name, link)

	message := []byte(fmt.Sprintf("From: %s\r\n"+
		"Subject: %s\r\n"+
		"\r\n"+
		"%s\r\n", from, subject, body))

	// Установка конфигурации для TLS
	tlsConfig := &tls.Config{
		InsecureSkipVerify: true, // Если сертификат не валидный
		ServerName:         smtpHost,
	}

	// Подключение к SMTP-серверу с использованием TLS
	conn, err := tls.Dial("tcp", smtpHost+":"+smtpPort, tlsConfig)
	if err != nil {
		log.Println("Ошибка при подключении к SMTP-серверу:", err)
		return err
	}
	defer conn.Close()

	// Аутентификация на SMTP-сервере
	auth := smtp.PlainAuth("", smtpUsername, smtpPassword, smtpHost)

	// Отправка письма каждому администратору
	for _, admin := range admins {
		to := []string{admin.Email}

		// Подключение к SMTP-серверу для каждого администратора
		client, err := smtp.NewClient(conn, smtpHost)
		if err != nil {
			log.Println("Ошибка при создании клиента SMTP:", err)
			continue
		}

		if err = client.Auth(auth); err != nil {
			log.Println("Ошибка при аутентификации на SMTP-сервере:", err)
			client.Quit()
			continue
		}

		if err = client.Mail(from); err != nil {
			log.Println("Ошибка при указании отправителя письма:", err)
			client.Quit()
			continue
		}

		for _, addr := range to {
			if err = client.Rcpt(addr); err != nil {
				log.Println("Ошибка при указании получателя письма:", err)
				client.Quit()
				continue
			}
		}

		w, err := client.Data()
		if err != nil {
			log.Println("Ошибка при отправке письма:", err)
			client.Quit()
			continue
		}

		_, err = w.Write(message)
		if err != nil {
			log.Println("Ошибка при записи сообщения письма:", err)
			w.Close()
			client.Quit()
			continue
		}

		err = w.Close()
		if err != nil {
			log.Println("Ошибка при закрытии письма:", err)
			client.Quit()
			continue
		}

		err = client.Quit()
		if err != nil {
			log.Println("Ошибка при завершении соединения SMTP:", err)
			continue
		}
	}

	return nil
}
