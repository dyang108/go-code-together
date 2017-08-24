import { languages, EventTypes } from './lists'
import socket from './client-socket'

var langSelect = document.getElementById('language-select')

class EditorWrapper {
  constructor (mode) {
    this.editor = ace.edit('editor')
    this.editor.setTheme('ace/theme/monokai')
    this.editor.session.setUseWorker(false)
    this.editor.session.setMode(mode)

    let splitUrl = document.location.href.split('/')
    this.roomId = splitUrl[splitUrl.length - 1]
    this.isChanging = false
    this.editor.session.on(EventTypes.EditorChange, function (e) {
      this.userEdit(e)
    }.bind(this))
  }

  setMode (newMode) {
    this.editor.session.setMode(newMode)
    let ind = languages.findIndex(elem => elem.src === newMode)
    langSelect.selectedIndex = ind
  }

  changeHighlighting (newMode) {
    this.editor.session.setMode(newMode)
    if (!this.isChanging) {
      socket.emit(EventTypes.LanguageChange, JSON.stringify({
        roomId: this.roomId,
        mode: newMode
      }))
    }
  }

  setKeyboardHandler (kh) {
    this.editor.setKeyboardHandler(kh)
  }

  getRoomId () {
    return this.roomId
  }

  userEdit (e) {
    if (!this.isChanging) {
      socket.emit(EventTypes.UserEdit, JSON.stringify({
        change: e,
        text: this.editor.getValue(),
        roomId: this.roomId
      }))
    }
  }

  serverEdit (data) {
    this.isChanging = true
    let edit = JSON.parse(data)
    let change = edit.change
    switch (change.action) {
      case EventTypes.TextInsertion:
        this.editor.session.insert(change.start, change.lines.join('\n'))
        break
      case EventTypes.TextRemoval:
        this.editor.session.remove(change)
    }
    this.isChanging = false
  }

  setSyntax (data) {
    let edit = JSON.parse(data)
    this.isChanging = true
    this.setMode(edit.mode)
    this.isChanging = false
  }
}

var editor = new EditorWrapper(mode)
export default editor
