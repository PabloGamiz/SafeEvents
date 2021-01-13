package radar

// Option values
const (
	queryFindByEventID            = "event_id = ?"
	queryFindByEventIDAndClientID = "event_id = ? and client_id = ?"
	queryFindByQR                 = "qr_code = ?"
	errRadarNotExists             = "Radar for MAC %s does not exists"
)
