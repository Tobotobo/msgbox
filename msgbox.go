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

type DefaultButton uint

const (
	DefaultButton1 DefaultButton = DefaultButton(walk.MsgBoxDefButton1)
	DefaultButton2 DefaultButton = DefaultButton(walk.MsgBoxDefButton2)
	DefaultButton3 DefaultButton = DefaultButton(walk.MsgBoxDefButton3)
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
	OwnerForm     *walk.Form
	Text          string
	Caption       string
	Icon          Icon
	Button        Button
	DefaultButton DefaultButton
	Result        Result
}

func New() *msgBox {
	return &msgBox{
		OwnerForm:     nil,
		Text:          "",
		Caption:       "",
		Icon:          IconNone,
		Button:        ButtonOK,
		DefaultButton: DefaultButton1,
		Result:        ResultNone,
	}
}

func (mb *msgBox) Owner(owner *walk.Form) *msgBox {
	mb.OwnerForm = owner
	return mb
}

func Owner(owner *walk.Form) *msgBox {
	mb := New()
	return mb.Owner(owner)
}

func (mb *msgBox) Show(text1caption2 ...string) *msgBox {
	if l := len(text1caption2); l >= 1 {
		mb.Text = text1caption2[0]
		if l >= 2 {
			mb.Caption = text1caption2[1]
		}
	}
	var owner walk.Form
	if mb.OwnerForm != nil {
		owner = *mb.OwnerForm
	} else {
		owner = nil
	}
	mb.Result = Result(walk.MsgBox(owner, mb.Caption, mb.Text, walk.MsgBoxStyle(mb.Icon)|walk.MsgBoxStyle(mb.Button)|walk.MsgBoxStyle(mb.DefaultButton)))
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
//  DefaultButton
// ----------------------------------------------------------------
func (mb *msgBox) DefBtn1() *msgBox {
	mb.DefaultButton = DefaultButton1
	return mb
}

func DefBtn1() *msgBox {
	mb := New()
	return mb.DefBtn1()
}

func (mb *msgBox) DefBtn2() *msgBox {
	mb.DefaultButton = DefaultButton2
	return mb
}

func DefBtn2() *msgBox {
	mb := New()
	return mb.DefBtn2()
}

func (mb *msgBox) DefBtn3() *msgBox {
	mb.DefaultButton = DefaultButton3
	return mb
}

func DefBtn3() *msgBox {
	mb := New()
	return mb.DefBtn2()
}

// ----------------------------------------------------------------
//  Result
// ----------------------------------------------------------------

func (r *Result) IsNone() bool {
	return *r == ResultNone
}

func (r *Result) IsOK() bool {
	return *r == ResultOK
}

func (r *Result) IsCancel() bool {
	return *r == ResultCancel
}

func (r *Result) IsYes() bool {
	return *r == ResultYes
}

func (r *Result) IsNo() bool {
	return *r == ResultNo
}

func (r *Result) IsAbort() bool {
	return *r == ResultAbort
}

func (r *Result) IsIgnore() bool {
	return *r == ResultIgnore
}

func (r *Result) IsRetry() bool {
	return *r == ResultRetry
}

func (mb *msgBox) IsNone() bool {
	return mb.Result.IsNone()
}

func (mb *msgBox) IsOK() bool {
	return mb.Result.IsOK()
}

func (mb *msgBox) IsCancel() bool {
	return mb.Result.IsCancel()
}

func (mb *msgBox) IsYes() bool {
	return mb.Result.IsYes()
}

func (mb *msgBox) IsNo() bool {
	return mb.Result.IsNo()
}

func (mb *msgBox) IsAbort() bool {
	return mb.Result.IsAbort()
}

func (mb *msgBox) IsIgnore() bool {
	return mb.Result.IsIgnore()
}

func (mb *msgBox) IsRetry() bool {
	return mb.Result.IsRetry()
}
