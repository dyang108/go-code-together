import { languages, EventTypes } from './lists'
import socket from './client-socket'
var langSelect = document.getElementById('language-select')

export default class Editor {
  constructor (mode, id) {
    this.editor = ace.edit(id)
    this.editor.setTheme('ace/theme/monokai')
    this.editor.session.setUseWorker(false)
    this.editor.session.setMode(mode)
    this.editor.$blockScrolling = Infinity
    this.editor.getSession().setUseWrapMode(true)
    this.roomId = window.location.pathname.slice(1)
    this.isChanging = false
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

  setSyntax (data) {
    this.isChanging = true
    this.setMode(data.mode)
    this.isChanging = false
  }
}
