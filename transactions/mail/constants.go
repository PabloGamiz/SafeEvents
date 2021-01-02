package mail

const (
	envSMTPHost = "SMTP_HOST"
	envSMTPPort = "SMTP_PORT"
	envMailUsr  = "FROM_MAIL"
	envMailPwd  = "FROM_PWD"

	emailSubject = "Subject: [SafeEvents] Notificació per contacte dirècte amb un positiu en COVID-19\n"
	emailMIME    = "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	emailInfoURL = "https://www.mscbs.gob.es/profesionales/saludPublica/ccayes/alertasActual/nCov/documentos/COVID19_que_hago_si_conozco_alguien_con_sintomas.pdf"
)
