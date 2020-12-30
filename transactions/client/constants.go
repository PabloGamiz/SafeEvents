package client

const (
	regexNoneSpecialChars = `^[a-zA-Z]+$`
	regexStandardEmail    = `^[A-Z0-9._%+-]+@[A-Z0-9.-]+\.[A-Z]{2,63}$`
	regexHash256          = `\b[A-Fa-f0-9]{64}\b`
	regexBase64           = `^(?:[A-Za-z0-9+/]{4})*(?:[A-Za-z0-9+/]{2}==|[A-Za-z0-9+/]{3}=)?$`

	errEmailFormat = "The provided email address does not has the expected format"
	errTokenFormat = "The provided token does no match with the expected format"

	envDummyEmail = "DUMMY_EMAIL"

	envSMTPHost = "SMTP_HOST"
	envSMTPPort = "SMTP_PORT"
	envMailUsr  = "FROM_MAIL"
	envMailPwd  = "FROM_PWD"

	emailSubject = "Subject: [SafeEvents] Notificació per contacte dirècte amb un positiu per COVID-19\n"
	emailMIME    = "MIME-version: 1.0;\nContent-Type: text/plain; charset=\"UTF-8\";\n\n"
	emailInfoURL = "https://www.mscbs.gob.es/profesionales/saludPublica/ccayes/alertasActual/nCov/documentos/COVID19_que_hago_si_conozco_alguien_con_sintomas.pdf"
)
