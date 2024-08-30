package services

import (
	"log"

	email "restaurant-catalog/email"
)

func SendReservationEmail() {
	log.Println("Sending email")

	//create message body
	//recipient email
	//Topic
	var message = getMessageBody()

	//email service
	emailService := email.NewEmailService(587, "smtp.gmail.com", "", "")
	isSEnd, err := emailService.SendEmail("knives.thao@gmail.com", "Reservation Topic", message)
	if err != nil {
		log.Fatalf("Error sending email: %s", err)
	}
	if isSEnd {
		log.Println("Email sent successfully")
	} else {
		log.Println("Email unable to send")
	}

}

func getMessageBody() string {
	return "test messsage sent from restcatapi"
}

func ConfirmReservations() {

}
