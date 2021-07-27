package msgbox

import "github.com/lxn/walk"

type Result uint

const (
	ResultNone   Result = walk.DlgCmdNone
	ResultOK     Result = walk.DlgCmdOK
	ResultCancel Result = walk.DlgCmdCancel
	ResultYes    Result = walk.DlgCmdYes
	ResultNo     Result = walk.DlgCmdNo
	ResultAbort  Result = walk.DlgCmdAbort
	ResultIgnore Result = walk.DlgCmdIgnore
	ResultRetry  Result = walk.DlgCmdRetry
)

type Button uint

const (
	ButtonAbortRetryIgnore Button = Button(walk.MsgBoxAbortRetryIgnore)
	ButtonOK               Button = Button(walk.MsgBoxOK)
	ButtonOKCancel         Button = Button(walk.MsgBoxOKCancel)
	ButtonRetryCancel      Button = Button(walk.MsgBoxRetryCancel)
	ButtonYesNo            Button = Button(walk.MsgBoxYesNo)
	ButtonYesNoCancel      Button = Button(walk.MsgBoxYesNoCancel)
)

type Icon uint

const (
	IconNone        Icon = Icon(0)
	IconError       Icon = Icon(walk.MsgBoxIconError)
	IconExclamation Icon = Icon(walk.MsgBoxIconExclamation)
	IconHand        Icon = Icon(walk.MsgBoxIconHand)
	IconInformation Icon = Icon(walk.MsgBoxIconInformation)
	IconQuestion    Icon = Icon(walk.MsgBoxIconQuestion)
	IconStop        Icon = Icon(walk.MsgBoxIconStop)
	IconWarning     Icon = Icon(walk.MsgBoxIconWarning)
)

type msgBox struct {
	Text    string
	Caption string
	Icon    Icon
	Button  Button
	Result  Result
}

func New() *msgBox {
	return &msgBox{
		Text:    "",
		Caption: "",
		Icon:    IconNone,
		Button:  ButtonOK,
		Result:  ResultNone,
	}
}

func (mb *msgBox) Show(text1caption2 ...string) *msgBox {
	if l := len(text1caption2); l >= 1 {
		mb.Text = text1caption2[0]
		if l >= 2 {
			mb.Caption = text1caption2[1]
		}
	}
	mb.Result = Result(walk.MsgBox(nil, mb.Caption, mb.Text, walk.MsgBoxStyle(mb.Icon)|walk.MsgBoxStyle(mb.Button)))
	return mb
}

func Show(text1caption2 ...string) *msgBox {
	mb := New()
	return mb.Show(text1caption2...)
}

// ----------------------------------------------------------------
//  Icon
// ----------------------------------------------------------------

func (mb *msgBox) Info() *msgBox {
	mb.Icon = IconInformation
	return mb
}

func Info() *msgBox {
	mb := New()
	return mb.Info()
}

func (mb *msgBox) Err() *msgBox {
	mb.Icon = IconError
	return mb
}

func Err() *msgBox {
	mb := New()
	return mb.Err()
}

func (mb *msgBox) Warn() *msgBox {
	mb.Icon = IconWarning
	return mb
}

func Warn() *msgBox {
	mb := New()
	return mb.Warn()
}

func (mb *msgBox) Ques() *msgBox {
	mb.Icon = IconQuestion
	return mb
}

func Ques() *msgBox {
	mb := New()
	return mb.Ques()
}

// ----------------------------------------------------------------
//  Button
// ----------------------------------------------------------------

func (mb *msgBox) AbortRetryIgnore() *msgBox {
	mb.Button = ButtonAbortRetryIgnore
	return mb
}

func AbortRetryIgnore() *msgBox {
	mb := New()
	return mb.AbortRetryIgnore()
}

func (mb *msgBox) OK() *msgBox {
	mb.Button = ButtonOK
	return mb
}

func OK() *msgBox {
	mb := New()
	return mb.OK()
}

func (mb *msgBox) OKCancel() *msgBox {
	mb.Button = ButtonOKCancel
	return mb
}

func OKCancel() *msgBox {
	mb := New()
	return mb.OKCancel()
}

func (mb *msgBox) RetryCancel() *msgBox {
	mb.Button = ButtonRetryCancel
	return mb
}

func RetryCancel() *msgBox {
	mb := New()
	return mb.RetryCancel()
}

func (mb *msgBox) YesNo() *msgBox {
	mb.Button = ButtonYesNo
	return mb
}

func YesNo() *msgBox {
	mb := New()
	return mb.YesNo()
}

func (mb *msgBox) YesNoCancel() *msgBox {
	mb.Button = ButtonYesNoCancel
	return mb
}

func YesNoCancel() *msgBox {
	mb := New()
	return mb.YesNoCancel()
}

// ----------------------------------------------------------------
//  Result
// ----------------------------------------------------------------

func (mb *msgBox) IsNone() bool {
	return mb.Result == ResultNone
}

func (mb *msgBox) IsOK() bool {
	return mb.Result == ResultOK
}

func (mb *msgBox) IsCancel() bool {
	return mb.Result == ResultCancel
}

func (mb *msgBox) IsYes() bool {
	return mb.Result == ResultYes
}

func (mb *msgBox) IsNo() bool {
	return mb.Result == ResultNo
}

func (mb *msgBox) IsAbort() bool {
	return mb.Result == ResultAbort
}

func (mb *msgBox) IsIgnore() bool {
	return mb.Result == ResultIgnore
}

func (mb *msgBox) IsRetry() bool {
	return mb.Result == ResultRetry
}
