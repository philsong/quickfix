package quickfixgo

import (
	"fmt"
	"github.com/cbusbey/quickfixgo/tag"
)

const (
	RequiredTagMissing         RejectReason = 1
	UnsupportedMessageType     RejectReason = 3
	ValueIsIncorrect           RejectReason = 5
	CompIDProblem              RejectReason = 9
	SendingTimeAccuracyProblem RejectReason = 10
	InvalidMsgType             RejectReason = 11
)

func NewRequiredTagMissing(msg Message, tag tag.Tag) MessageReject {
	return messageRejectBase{rejectedMessage: msg, text: "Required tag missing", rejectReason: RequiredTagMissing,
		refTagID: tag}
}

func NewValueIsIncorrect(msg Message, tag tag.Tag) MessageReject {
	return messageRejectBase{rejectedMessage: msg, text: "Value is incorrect (out of range) for this tag", rejectReason: ValueIsIncorrect,
		refTagID: tag}
}

func NewInvalidMessageType(msg Message) MessageReject {
	return messageRejectBase{rejectedMessage: msg, text: "Invalid MsgType", rejectReason: InvalidMsgType}
}

func NewUnsupportedMessageType(msg Message) MessageReject {
	return messageRejectBase{rejectedMessage: msg, text: "Unsupported Message Type", rejectReason: UnsupportedMessageType, businessReject: true}
}

type IncorrectBeginString struct{ MessageReject }

func NewIncorrectBeginString(msg Message) IncorrectBeginString {
	return IncorrectBeginString{messageRejectBase{rejectedMessage: msg, text: "Incorrect BeginString"}}
}

func NewCompIDProblem(msg Message) MessageReject {
	return messageRejectBase{rejectedMessage: msg, text: "CompID problem", rejectReason: CompIDProblem}
}

func NewSendingTimeAccuracyProblem(msg Message) MessageReject {
	return messageRejectBase{rejectedMessage: msg, text: "SendingTime accuracy problem", rejectReason: SendingTimeAccuracyProblem}
}

type TargetTooHigh struct {
	MessageReject
	ReceivedTarget int
	ExpectedTarget int
}

type TargetTooLow struct {
	MessageReject
	ReceivedTarget int
	ExpectedTarget int
}

func NewTargetTooHigh(msg Message, receivedTarget, expectedTarget int) TargetTooHigh {
	return TargetTooHigh{
		MessageReject: messageRejectBase{
			rejectedMessage: msg,
			text:            fmt.Sprintf("MsgSeqNum too high, expecting %d but received %d", expectedTarget, receivedTarget)},
		ReceivedTarget: receivedTarget,
		ExpectedTarget: expectedTarget}
}

func NewTargetTooLow(msg Message, receivedTarget, expectedTarget int) TargetTooLow {
	return TargetTooLow{
		MessageReject: messageRejectBase{
			rejectedMessage: msg,
			text:            fmt.Sprintf("MsgSeqNum too low, expecting %d but received %d", expectedTarget, receivedTarget)},
		ReceivedTarget: receivedTarget,
		ExpectedTarget: expectedTarget}
}