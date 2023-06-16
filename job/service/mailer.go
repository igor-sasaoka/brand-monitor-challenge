package service 

import (
    "net/smtp"
    "os"
)

func sendMail(to string) error {
    pass := os.Getenv("MAIL_PASS")
    from := os.Getenv("MAIL_SENDER")
    msg := "From: " + from + "\n" + "To: " + to + "\n" +
    "Subject: Your account was successfully created!\n\n"

    err := smtp.SendMail(
        "smtp.gmail.com:587",
        smtp.PlainAuth("", from, pass, "smtp.gmail.com"),
        from,
        []string{to},
        []byte(msg),
    )
    return err
}
