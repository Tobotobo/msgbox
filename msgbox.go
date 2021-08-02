// msgbox.go
// Copyright (c) 2021 Tobotobo
// This software is released under the MIT License.
// http://opensource.org/licenses/mit-license.php

// Copyright 2010 The Walk Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
// https://github.com/lxn/walk/blob/master/LICENSE

// +build windows

//lint:file-ignore SA1019 syscall.StringToUTF16Ptr を使用します

package msgbox

import (
	"strings"
	"syscall"

	"github.com/lxn/walk"
	"github.com/lxn/win"
)

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
	Owner   win.HWND
	Text    string
	Caption string
	Icon    Icon
	Button  Button
}

type MsgBox struct {
	InnerValue msgBox
	Result     Result
}

func New() *MsgBox {
	return &MsgBox{
		InnerValue: msgBox{
			Owner:   0,
			Text:    "",
			Caption: "",
			Icon:    IconNone,
			Button:  ButtonOK,
		},
		Result: ResultNone,
	}
}

func (mb *MsgBox) Show(text1caption2 ...string) *MsgBox {
	if l := len(text1caption2); l >= 1 {
		mb.InnerValue.Text = text1caption2[0]
		if l >= 2 {
			mb.InnerValue.Caption = text1caption2[1]
		}
	}
	mb.Result = Result(show(mb.InnerValue.Owner, mb.InnerValue.Caption, mb.InnerValue.Text, uint32(mb.InnerValue.Icon)|uint32(mb.InnerValue.Button)))
	return mb
}

func Show(text1caption2 ...string) *MsgBox {
	mb := New()
	return mb.Show(text1caption2...)
}

func show(owner win.HWND, title, message string, style uint32) int {
	return int(win.MessageBox(
		owner,
		syscall.StringToUTF16Ptr(strings.ReplaceAll(message, "\x00", "␀")),
		syscall.StringToUTF16Ptr(strings.ReplaceAll(title, "\x00", "␀")),
		uint32(style)))
}

// ----------------------------------------------------------------
//  Icon
// ----------------------------------------------------------------

func (mb *MsgBox) Info() *MsgBox {
	mb.InnerValue.Icon = IconInformation
	return mb
}

func Info() *MsgBox {
	mb := New()
	return mb.Info()
}

func (mb *MsgBox) Err() *MsgBox {
	mb.InnerValue.Icon = IconError
	return mb
}

func Err() *MsgBox {
	mb := New()
	return mb.Err()
}

func (mb *MsgBox) Warn() *MsgBox {
	mb.InnerValue.Icon = IconWarning
	return mb
}

func Warn() *MsgBox {
	mb := New()
	return mb.Warn()
}

func (mb *MsgBox) Ques() *MsgBox {
	mb.InnerValue.Icon = IconQuestion
	return mb
}

func Ques() *MsgBox {
	mb := New()
	return mb.Ques()
}

// ----------------------------------------------------------------
//  Button
// ----------------------------------------------------------------

func (mb *MsgBox) AbortRetryIgnore() *MsgBox {
	mb.InnerValue.Button = ButtonAbortRetryIgnore
	return mb
}

func AbortRetryIgnore() *MsgBox {
	mb := New()
	return mb.AbortRetryIgnore()
}

func (mb *MsgBox) OK() *MsgBox {
	mb.InnerValue.Button = ButtonOK
	return mb
}

func OK() *MsgBox {
	mb := New()
	return mb.OK()
}

func (mb *MsgBox) OKCancel() *MsgBox {
	mb.InnerValue.Button = ButtonOKCancel
	return mb
}

func OKCancel() *MsgBox {
	mb := New()
	return mb.OKCancel()
}

func (mb *MsgBox) RetryCancel() *MsgBox {
	mb.InnerValue.Button = ButtonRetryCancel
	return mb
}

func RetryCancel() *MsgBox {
	mb := New()
	return mb.RetryCancel()
}

func (mb *MsgBox) YesNo() *MsgBox {
	mb.InnerValue.Button = ButtonYesNo
	return mb
}

func YesNo() *MsgBox {
	mb := New()
	return mb.YesNo()
}

func (mb *MsgBox) YesNoCancel() *MsgBox {
	mb.InnerValue.Button = ButtonYesNoCancel
	return mb
}

func YesNoCancel() *MsgBox {
	mb := New()
	return mb.YesNoCancel()
}

// ----------------------------------------------------------------
//  Result
// ----------------------------------------------------------------

func (mb *MsgBox) IsNone() bool {
	return mb.Result == ResultNone
}

func (mb *MsgBox) IsOK() bool {
	return mb.Result == ResultOK
}

func (mb *MsgBox) IsCancel() bool {
	return mb.Result == ResultCancel
}

func (mb *MsgBox) IsYes() bool {
	return mb.Result == ResultYes
}

func (mb *MsgBox) IsNo() bool {
	return mb.Result == ResultNo
}

func (mb *MsgBox) IsAbort() bool {
	return mb.Result == ResultAbort
}

func (mb *MsgBox) IsIgnore() bool {
	return mb.Result == ResultIgnore
}

func (mb *MsgBox) IsRetry() bool {
	return mb.Result == ResultRetry
}
