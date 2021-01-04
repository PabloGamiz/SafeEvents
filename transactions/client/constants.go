package client

const (
	regexNoneSpecialChars = `^[a-zA-Z]+$`
	regexStandardEmail    = `^[A-Z0-9._%+-]+@[A-Z0-9.-]+\.[A-Z]{2,63}$`
	regexHash256          = `\b[A-Fa-f0-9]{64}\b`
	regexBase64           = `^(?:[A-Za-z0-9+/]{4})*(?:[A-Za-z0-9+/]{2}==|[A-Za-z0-9+/]{3}=)?$`

	errEmailFormat = "The provided email address does not has the expected format"
	errTokenFormat = "The provided token does no match with the expected format"

	envDummyEmail = "DUMMY_EMAIL"
	marginTime    = 24 * 15
)
