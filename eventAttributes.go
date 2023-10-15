package godom

import (
	"github.com/tbe/godom/helpers"
	"github.com/tbe/godom/types"
)

// The OnAfterPrint attribute specifies a script to be run after the document is printed.
//
// This is a global types.Attribute.
func OnAfterPrint(script string) types.Attribute {
	return helpers.SingleAttribute("onafterprint", script)
}

// The OnBeforePrint attribute specifies a script to be run before the document is printed.
//
// This is a global types.Attribute.
func OnBeforePrint(script string) types.Attribute {
	return helpers.SingleAttribute("onbeforeprint", script)
}

// The OnBeforeUnload attribute specifies a script to be run when the document is about to be unloaded.
//
// This is a global types.Attribute.
func OnBeforeUnload(script string) types.Attribute {
	return helpers.SingleAttribute("onbeforeunload", script)
}

// The OnError attribute specifies a script to be run when an error occurs.
//
// This is a global types.Attribute.
func OnError(script string) types.Attribute {
	return helpers.SingleAttribute("onerror", script)
}

// The OnHashChange attribute specifies a script to be run when there has been changes to the anchor part of the URL.
//
// This is a global types.Attribute.
func OnHashChange(script string) types.Attribute {
	return helpers.SingleAttribute("onhashchange", script)
}

// The OnLoad attribute specifies a script to be run after the page is finished loading.
//
// This is a global types.Attribute.
func OnLoad(script string) types.Attribute {
	return helpers.SingleAttribute("onload", script)
}

// The OnMessage attribute specifies a script to be run when the message is triggered.
//
// This is a global types.Attribute.
func OnMessage(script string) types.Attribute {
	return helpers.SingleAttribute("onmessage", script)
}

// The OnOffline attribute specifies a script to be run when the browser starts to work offline.
//
// This is a global types.Attribute.
func OnOffline(script string) types.Attribute {
	return helpers.SingleAttribute("onoffline", script)
}

// The OnOnline attribute specifies a script  to be run when the browser starts to work online.
//
// This is a global types.Attribute.
func OnOnline(script string) types.Attribute {
	return helpers.SingleAttribute("ononline", script)
}

// The OnPageHide attribute specifies a script to be run when a user navigates away from a page.
//
// This is a global types.Attribute.
func OnPageHide(script string) types.Attribute {
	return helpers.SingleAttribute("onpagehide", script)
}

// The OnPageShow attribute specifies a script to be run when a user navigates to a page.
//
// This is a global types.Attribute.
func OnPageShow(script string) types.Attribute {
	return helpers.SingleAttribute("onpageshow", script)
}

// The OnPopState attribute specifies a script to be run when the window's history changes.
//
// This is a global types.Attribute.
func OnPopState(script string) types.Attribute {
	return helpers.SingleAttribute("onpopstate", script)
}

// The OnResize attribute specifies a script to be run when the browser window is resized.
//
// This is a global types.Attribute.
func OnResize(script string) types.Attribute {
	return helpers.SingleAttribute("onresize", script)
}

// The OnStorage attribute specifies a script to be run when a Web Storage area is updated.
//
// This is a global types.Attribute.
func OnStorage(script string) types.Attribute {
	return helpers.SingleAttribute("onstorage", script)
}

// The OnUnload attribute specifies a script to be run once the page has unloaded (or the browser window has been closed).
//
// This is a global types.Attribute.
func OnUnload(script string) types.Attribute {
	return helpers.SingleAttribute("onunload", script)
}

// The OnBlur attribute specifies a script to be run the moment that the Element loses focus.
//
// This is a global types.Attribute.
func OnBlur(script string) types.Attribute {
	return helpers.SingleAttribute("onblur", script)
}

// The OnChange attribute specifies a script to be run the moment when the value of the Element is changed.
//
// This is a global types.Attribute.
func OnChange(script string) types.Attribute {
	return helpers.SingleAttribute("onchange", script)
}

// The OnContextMenu attribute specifies a script to be run when a context menu is triggered.
//
// This is a global types.Attribute.
func OnContextMenu(script string) types.Attribute {
	return helpers.SingleAttribute("oncontextmenu", script)
}

// The OnFocus attribute specifies a script to be run the moment when the Element gets focus.
//
// This is a global types.Attribute.
func OnFocus(script string) types.Attribute {
	return helpers.SingleAttribute("onfocus", script)
}

// The OnInput attribute specifies a script to be run when an Element gets user input.
//
// This is a global types.Attribute.
func OnInput(script string) types.Attribute {
	return helpers.SingleAttribute("oninput", script)
}

// The OnInvalid attribute specifies a script to be run when an Element is invalid.
//
// This is a global types.Attribute.
func OnInvalid(script string) types.Attribute {
	return helpers.SingleAttribute("oninvalid", script)
}

// The OnReset attribute specifies a script to be run when the Reset Button in a form is clicked.
//
// This is a global types.Attribute.
func OnReset(script string) types.Attribute {
	return helpers.SingleAttribute("onreset", script)
}

// The OnSearch attribute specifies a script to be run when the user writes something in a search field.
//
// This is a global types.Attribute.
func OnSearch(script string) types.Attribute {
	return helpers.SingleAttribute("onsearch", script)
}

// The OnSelect attribute specifies a script to be run after some text has been selected in an Element.
//
// This is a global types.Attribute.
func OnSelect(script string) types.Attribute {
	return helpers.SingleAttribute("onselect", script)
}

// The OnKeyDown attribute specifies a script to be run when a user is pressing a key.
//
// This is a global types.Attribute.
func OnKeyDown(script string) types.Attribute {
	return helpers.SingleAttribute("onkeydown", script)
}

// The OnKeyPress attribute specifies a script to be run when a user presses a key.
//
// This is a global types.Attribute.
func OnKeyPress(script string) types.Attribute {
	return helpers.SingleAttribute("onkeypress", script)
}

// The OnKeyUp attribute specifies a script to be run when a user releases a key.
//
// This is a global types.Attribute.
func OnKeyUp(script string) types.Attribute {
	return helpers.SingleAttribute("onkeyup", script)
}

// The OnClick attribute specifies a script to be run on a mouse click on the Element.
//
// This is a global types.Attribute.
func OnClick(script string) types.Attribute {
	return helpers.SingleAttribute("onclick", script)
}

// The OnDoubleClick attribute specifies a script to be run on a mouse double-click on the Element.
//
// This is a global types.Attribute.
func OnDoubleClick(script string) types.Attribute {
	return helpers.SingleAttribute("ondblclick", script)
}

// The OnMouseDown attribute specifies a script to be run when a mouse button is pressed down on an element.
//
// This is a global types.Attribute.
func OnMouseDown(script string) types.Attribute {
	return helpers.SingleAttribute("onmousedown", script)
}

// The OnMouseMove attribute specifies a script to be run when the mouse pointer is moving while it is over an element.
//
// This is a global types.Attribute.
func OnMouseMove(script string) types.Attribute {
	return helpers.SingleAttribute("onmousemove", script)
}

// The OnMouseOut attribute specifies a script to be run when the mouse pointer moves out of an element.
//
// This is a global types.Attribute.
func OnMouseOut(script string) types.Attribute {
	return helpers.SingleAttribute("onmouseout", script)
}

// The OnMouseOver attribute specifies a script to be run when the mouse pointer moves over an element.
//
// This is a global types.Attribute.
func OnMouseOver(script string) types.Attribute {
	return helpers.SingleAttribute("onmouseover", script)
}

// The OnMouseUp attribute specifies a script to be run when a mouse button is released over an element.
//
// This is a global types.Attribute.
func OnMouseUp(script string) types.Attribute {
	return helpers.SingleAttribute("onmouseup", script)
}

// The OnWheel attribute specifies a script to be run when the mouse wheel rolls up or down over an element.
//
// This is a global types.Attribute.
func OnWheel(script string) types.Attribute {
	return helpers.SingleAttribute("onwheel", script)
}

// The OnDrag attribute specifies a script to be run when an element is dragged.
//
// This is a global types.Attribute.
func OnDrag(script string) types.Attribute {
	return helpers.SingleAttribute("ondrag", script)
}

// The OnDragEnd attribute specifies a script to be run at the end of a drag operation.
//
// This is a global types.Attribute.
func OnDragEnd(script string) types.Attribute {
	return helpers.SingleAttribute("ondragend", script)
}

// The OnDragEnter attribute specifies a script to be run when an element has been dragged to a valid drop target.
//
// This is a global types.Attribute.
func OnDragEnter(script string) types.Attribute {
	return helpers.SingleAttribute("ondragenter", script)
}

// The OnDragLeave attribute specifies a script to be run when an element leaves a valid drop target.
//
// This is a global types.Attribute.
func OnDragLeave(script string) types.Attribute {
	return helpers.SingleAttribute("ondragleave", script)
}

// The OnDragOver attribute specifies a script to be run when an element is being dragged over a valid drop target.
//
// This is a global types.Attribute.
func OnDragOver(script string) types.Attribute {
	return helpers.SingleAttribute("ondragover", script)
}

// The OnDragStart attribute specifies a script to be run at the start of a drag operation.
//
// This is a global types.Attribute.
func OnDragStart(script string) types.Attribute {
	return helpers.SingleAttribute("ondragstart", script)
}

// The OnDrop attribute specifies a script to be run when dragged element is being dropped.
//
// This is a global types.Attribute.
func OnDrop(script string) types.Attribute {
	return helpers.SingleAttribute("ondrop", script)
}

// The OnScroll attribute specifies a script to be run when an element's scrollbar is being scrolled.
//
// This is a global types.Attribute.
func OnScroll(script string) types.Attribute {
	return helpers.SingleAttribute("onscroll", script)
}

// The OnCopy attribute specifies a script to be run when the user copies the content of an element.
//
// This is a global types.Attribute.
func OnCopy(script string) types.Attribute {
	return helpers.SingleAttribute("oncopy", script)
}

// The OnCut attribute specifies a script to be run when the user cuts the content of an element.
//
// This is a global types.Attribute.
func OnCut(script string) types.Attribute {
	return helpers.SingleAttribute("oncut", script)
}

// The OnPaste attribute specifies a script to be run when the user pastes some content in an element.
//
// This is a global types.Attribute.
func OnPaste(script string) types.Attribute {
	return helpers.SingleAttribute("onpaste", script)
}

// The OnAbort attribute specifies a script to be run on abort.
//
// This is a global types.Attribute.
func OnAbort(script string) types.Attribute {
	return helpers.SingleAttribute("onabort", script)
}

// The OnCanPlay attribute specifies a script to be run when a file is ready to start playing.
//
// This is a global types.Attribute.
func OnCanPlay(script string) types.Attribute {
	return helpers.SingleAttribute("oncanplay", script)
}

// The OnCanPlayThrough attribute specifies a script to be run when a file can be played all the way to the end without
// pausing for buffering.
//
// This is a global types.Attribute.
func OnCanPlayThrough(script string) types.Attribute {
	return helpers.SingleAttribute("oncanplaythrough", script)
}

// The OnCueChange attribute specifies a script to be run when the cue changes in a Track element.
//
// This is a global types.Attribute.
func OnCueChange(script string) types.Attribute {
	return helpers.SingleAttribute("oncuechange", script)
}

// The OnDurationChange attribute specifies a script to be run when the length of the media changes.
//
// This is a global types.Attribute.
func OnDurationChange(script string) types.Attribute {
	return helpers.SingleAttribute("ondurationchange", script)
}

// The OnEmptied attribute specifies a script to be run when something bad happens and the file is suddenly unavailable.
//
// This is a global types.Attribute.
func OnEmptied(script string) types.Attribute {
	return helpers.SingleAttribute("onemptied", script)
}

// The OnEnded attribute specifies a script to be run when the media has reach the end.
//
// This is a global types.Attribute.
func OnEnded(script string) types.Attribute {
	return helpers.SingleAttribute("onended", script)
}

// The OnLoadedData attribute specifies a script to be run when media data is loaded.
//
// This is a global types.Attribute.
func OnLoadedData(script string) types.Attribute {
	return helpers.SingleAttribute("onloadeddata", script)
}

// The OnLoadedMetaData attribute specifies a script to be run when meta data (like dimensions and duration) are loaded.
//
// This is a global types.Attribute.
func OnLoadedMetaData(script string) types.Attribute {
	return helpers.SingleAttribute("onloadedmetadata", script)
}

// The OnLoadStart attribute specifies a script to be run just as the file begins to load before anything is actually loaded.
//
// This is a global types.Attribute.
func OnLoadStart(script string) types.Attribute {
	return helpers.SingleAttribute("onloadstart", script)
}

// The OnPause attribute specifies a script to be run when the media is paused either by the user or programmatically.
//
// This is a global types.Attribute.
func OnPause(script string) types.Attribute {
	return helpers.SingleAttribute("onpause", script)
}

// The OnPlay attribute specifies a script to be run when the media is ready to start playing.
//
// This is a global types.Attribute.
func OnPlay(script string) types.Attribute {
	return helpers.SingleAttribute("onplay", script)
}

// The OnPlaying attribute specifies a script to be run when  the media actually has started playing.
//
// This is a global types.Attribute.
func OnPlaying(script string) types.Attribute {
	return helpers.SingleAttribute("onplaying", script)
}

// The OnProgress attribute specifies a script to be run when the browser is in the process of getting the media data.
//
// This is a global types.Attribute.
func OnProgress(script string) types.Attribute {
	return helpers.SingleAttribute("onprogress", script)
}

// The OnRateChange attribute specifies a script to be run each time the playback rate changes.
//
// This is a global types.Attribute.
func OnRateChange(script string) types.Attribute {
	return helpers.SingleAttribute("onratechange", script)
}

// The OnSeeked attribute specifies a script to be run when the seeking attribute is set to false indicating that seeking has ended.
//
// This is a global types.Attribute.
func OnSeeked(script string) types.Attribute {
	return helpers.SingleAttribute("onseeked", script)
}

// The OnSeeking attribute specifies a script to be run when the seeking attribute is set to true indicating that seeking is active.
//
// This is a global types.Attribute.
func OnSeeking(script string) types.Attribute {
	return helpers.SingleAttribute("onseeking", script)
}

// The OnStalled attribute specifies a script to be run when the browser is unable to fetch the media data for whatever reason.
//
// This is a global types.Attribute.
func OnStalled(script string) types.Attribute {
	return helpers.SingleAttribute("onstalled", script)
}

// The OnSuspend attribute specifies a script to be run when fetching the media data is stopped before it is completely
// loaded for whatever reason.
//
// This is a global types.Attribute.
func OnSuspend(script string) types.Attribute {
	return helpers.SingleAttribute("onsuspend", script)
}

// The OnTimeUpdate attribute specifies a script to be run when the playing position has changed.
//
// This is a global types.Attribute.
func OnTimeUpdate(script string) types.Attribute {
	return helpers.SingleAttribute("ontimeupdate", script)
}

// The OnVolumeChange attribute specifies a script to be run each time the volume is changed which.
//
// This is a global types.Attribute.
func OnVolumeChange(script string) types.Attribute {
	return helpers.SingleAttribute("onvolumechange", script)
}

// The OnWaiting attribute specifies a script to be run when the media has paused but is expected to resume.
//
// This is a global types.Attribute.
func OnWaiting(script string) types.Attribute {
	return helpers.SingleAttribute("onwaiting", script)
}

// The OnToggle attribute specifies a script to be run when the user opens or closes the Details element.
//
// This is a global types.Attribute.
func OnToggle(script string) types.Attribute {
	return helpers.SingleAttribute("ontoggle", script)
}

// The OnSubmit attribute specifies a script to be run when a Form is submitted.s
//
// This flag is allowed for:
// - Form
func OnSubmit(script string) types.Attribute {
	return helpers.SingleAttribute("onsubmit", script)
}
