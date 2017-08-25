import { EventTypes } from './lists'
import socket from './client-socket'

var Range = ace.require('ace/range').Range

export default class SelectionManager {
  constructor (editor) {
    this.editor = editor

    this.editor.selection.on(EventTypes.ChangeCursor, function (e) {
      socket.emit(EventTypes.ChangeCursor, JSON.stringify({
        change: this.editor.selection.getCursor(),
        clientId: socket.id
      }))
    }.bind(this))

    this.editor.selection.on(EventTypes.ChangeSelection, function (e) {
      socket.emit(EventTypes.ChangeSelection, JSON.stringify({
        changes: this.editor.selection.getAllRanges(),
        clientId: socket.id
      }))
    }.bind(this))
    this.clients = {}
  }

  cursorChanged (msg) {
    let cId = msg.clientId
    if (this.clients[cId]) {
      this.editor.session.removeMarker(this.clients[cId].cursor)
    } else {
      this.clients[cId] = {
        color: Math.floor(Math.random() * 6),
        selections: []
      }
    }
    let range = new Range(msg.change.row, msg.change.column, msg.change.row, msg.change.column + 1)
    let cls = 'other-cursor color' + this.clients[cId].color
    let newCursorPos = this.editor.session.addMarker(range, cls, 'text', true)
    this.clients[cId].cursor = newCursorPos
  }

  selectionChanged (msg) {
    let cId = msg.clientId
    if (this.clients[cId]) {
      for (let i = 0; i < this.clients[cId].selections.length; i++) {
        this.editor.session.removeMarker(this.clients[cId].selections[i])
      }
      this.clients[cId].selections = []
    } else {
      this.clients[cId] = {
        color: Math.floor(Math.random() * 6),
        selections: []
      }
    }
    for (let i = 0; i < msg.changes.length; i++) {
      let change = msg.changes[i]
      if (change.start.column !== change.end.column || change.start.row !== change.end.row) {
        let range = new Range(change.start.row, change.start.column, change.end.row, change.end.column)
        let cls = 'other-selection bg-color' + this.clients[cId].color
        let newCursorPos = this.editor.session.addMarker(range, cls, 'line', true)
        this.clients[cId].selections.push(newCursorPos)
      }
    }
  }

  removeOtherUser (clientId) {
    if (this.clients[clientId]) {
      this.clients[clientId].selections.forEach(sel => {
        this.editor.session.removeMarker(sel)
      })
      this.editor.session.removeMarker(this.clients[clientId].cursor)
      delete this.clients[clientId]
    }
  }
}
