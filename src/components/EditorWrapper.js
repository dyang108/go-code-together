import Editor from './Editor'
import { EventTypes } from './lists'
import socket from './client-socket'
import SelectionManager from './SelectionManager'

export default class EditorWrapper extends Editor {
  constructor (mode, id) {
    super(mode, id)
    this.editor.session.on(EventTypes.EditorChange, function (e) {
      this.userEdit(e)
    }.bind(this))

    this.selectionManager = new SelectionManager(this.editor)
  }
  serverChangeSelection (msg) {
    this.selectionManager.selectionChanged(msg)
  }

  serverChangeCursor (msg) {
    this.selectionManager.cursorChanged(msg)
  }

  removeOtherUser (clientId) {
    this.selectionManager.removeOtherUser(clientId)
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
    let change = data.change
    switch (change.action) {
      case EventTypes.TextInsertion:
        this.editor.session.insert(change.start, change.lines.join('\n'))
        break
      case EventTypes.TextRemoval:
        this.editor.session.remove(change)
    }
    this.isChanging = false
  }
}
