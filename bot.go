package flamingo

// AnswerChecker is a function that will determine if the provided
// answer is correct analysing the answer Message. If the function
// returns nil it is assumed the answer is valid, otherwise it is
// assumed to be invalid and the question will be repeated with the
// returned message. Take a look at the AskUntil method of the Bot interface.
type AnswerChecker func(Message) *OutgoingMessage

// Bot is the main interface to interact with the user. Either using controllers
// or action handlers these are all the exposed methods to communicate with
// the user.
type Bot interface {
	// ID returns the ID of the bot.
	ID() string

	// Reply replies a Message with an OutgoingMessage and returns the ID of the
	// reply along with an error, if any.
	Reply(Message, OutgoingMessage) (string, error)

	// Ask sends and OutgoingMessage (typically a question) and returns the ID of
	// the question asked, the Message the user replied and an error, if any.
	Ask(OutgoingMessage) (string, Message, error)

	// Conversation asks a series of questions and waits for answers to these
	// questions by the user. Then returns the IDs of all questions, all replies
	// and an error, if any.
	Conversation(Conversation) ([]string, []Message, error)

	// Say writes an OutgoingMessage and returns the ID of that message along with
	// an error, if any.
	Say(OutgoingMessage) (string, error)

	// Form posts a form and returns the ID of the form and an error, if any.
	Form(Form) (string, error)

	// Image posts an Image and returns the ID of the image message and an error,
	// if any.
	Image(Image) (string, error)

	// UpdateMessage updates the message with the given ID with the replacement
	// text and returns the ID of the new message along with an error, if any.
	UpdateMessage(id string, replacement string) (string, error)

	// UpdateForm updates the message with the given ID and replaces it with the
	// given form. Returns the ID of the new form and an error, if any.
	UpdateForm(id string, form Form) (string, error)

	// WaitForAction will block until an action with the given ID comes. Until
	// then, all the incoming messages or actions will be handled according to
	// the given waiting policy.
	WaitForAction(string, ActionWaitingPolicy) (Action, error)

	// AskUntil posts a question and checks the received message. If the
	// AnswerChecker considers it is correct, it will return. If not, the message
	// returned by the AnswerChecker will be posted and the process will repeat.
	// It returns the ID of the last question posted, the last user message and
	// an error, if any.
	AskUntil(OutgoingMessage, AnswerChecker) (string, Message, error)
}
