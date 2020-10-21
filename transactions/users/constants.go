package users

const (
	regexName    = `^[a-zA-Z]+$`
	regexEmail   = `^[A-Z0-9._%+-]+@[A-Z0-9.-]+\.[A-Z]{2,63}$`
	regexHash256 = `\b[A-Fa-f0-9]{64}\b`
	regexB64     = `^(?:[A-Za-z0-9+/]{4})*(?:[A-Za-z0-9+/]{2}==|[A-Za-z0-9+/]{3}=)?$`
)
