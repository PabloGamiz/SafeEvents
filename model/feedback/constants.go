package feedback

const (
	queryFindByIDAssistantIDEventID = "id = ? AND assistant_id = ? AND event_ID = ?"
	queryFindByAssistantIDEventID   = "assistant_id = ? AND event_ID = ?"
	queryFindByEventID              = "event_ID = ?"
	queryFindByAssistantID          = "assistant_id = ?"

	errNoMatchingFeedbackForAssistantAndEvent = "There not exists any feedback matching the provided assistant and event!"
)
