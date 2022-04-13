package main

import (
	"github.com/ruang-guru/playground/backend/data-structure/assignment/text-editor/stack"
)

// Implementasikan teks editor sederhana
// Teks editor ini memiliki 4 method
// Write: digunakan untuk menulis per karakter
// Read: digunakan untuk mencetak karakter yang telah ditulis
// Undo: digunakan untuk melakukan operasi undo
// Redo: digunakan untuk melakukan operasi redo

type TextEditor struct {
	UndoStack *stack.Stack
	RedoStack *stack.Stack
}

func NewTextEditor() *TextEditor {
	return &TextEditor{
		UndoStack: stack.NewStack(),
		RedoStack: stack.NewStack(),
	}
}

func (te *TextEditor) Write(ch rune) {
	te.UndoStack.Data = append(te.UndoStack.Data, ch)
	te.UndoStack.Top++

	// TODO: answer here
}

func (te *TextEditor) Read() []rune {
	if len(te.UndoStack.Data) == 0 {
		return nil
	}
	return te.UndoStack.Data
}

func (te *TextEditor) Undo() {
	if te.UndoStack.IsEmpty() == false {
		te.RedoStack.Data = append(te.RedoStack.Data, te.UndoStack.Data[te.UndoStack.Top])
		te.RedoStack.Top++

		te.UndoStack.Data = te.UndoStack.Data[:te.UndoStack.Top]
		te.UndoStack.Top--
	}

}

func (te *TextEditor) Redo() {
	if te.RedoStack.IsEmpty() == false {
		te.UndoStack.Data = append(te.UndoStack.Data, te.RedoStack.Data[te.RedoStack.Top])
		te.UndoStack.Top++

		te.RedoStack.Data = te.RedoStack.Data[:te.RedoStack.Top]
		te.RedoStack.Top--
	}
}
