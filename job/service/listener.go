package service

const LISTENERS_AMOUNT = 10
const MAX_ATTEMPTS = 4

func processMail() {
    r := GetAccountRegistry().SendMail
    for email := range r {
        for i := 0; i < MAX_ATTEMPTS; i++ {    
            err := sendMail(email);
            if err == nil {
                break
            }
        }
    }
}

func SpawnListeners() {
    for i := 0; i < LISTENERS_AMOUNT; i++ {
        go processMail()
    }
}
