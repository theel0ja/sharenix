/*
   Copyright 2014 Franc[e]sco (lolisamurai@tfwno.gf)
   This file is part of sharenix.
   sharenix is free software: you can redistribute it and/or modify
   it under the terms of the GNU General Public License as published by
   the Free Software Foundation, either version 3 of the License, or
   (at your option) any later version.
   sharenix is distributed in the hope that it will be useful,
   but WITHOUT ANY WARRANTY; without even the implied warranty of
   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
   GNU General Public License for more details.
   You should have received a copy of the GNU General Public License
   along with sharenix. If not, see <http://www.gnu.org/licenses/>.
*/

package sharenixlib

import (
	"github.com/conformal/gotk3/gdk"
	"github.com/conformal/gotk3/gtk"
)

// GetClipboard returns the default display's GTK clipboard
func GetClipboard() (*gtk.Clipboard, error) {
	display, err := gdk.DisplayGetDefault()
	if err != nil {
		return nil, err
	}

	clipboard, err := gtk.ClipboardGetForDisplay(
		display, gdk.SELECTION_CLIPBOARD)
	if err != nil {
		return nil, err
	}

	return clipboard, nil
}

// SetClipbboardText sets the clipboard text contents and calls
// clipboard.Store().
// Note: this requires the program to run at least a few cycles of the main loop
// and it is not guaranteed to persist on all window managers once the program
// terminates.
func SetClipboardText(text string) (err error) {
	clipboard, err := GetClipboard()
	if err != nil {
		return err
	}

	clipboard.SetText(text)
	gtk.MainIterationDo(true)
	clipboard.Store()
	gtk.MainIterationDo(true)
	return
}
