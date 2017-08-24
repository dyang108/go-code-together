import socket from './client-socket'
import editor from './editor'
import { EventTypes } from './lists'

socket.on(EventTypes.Connection, () => {
  socket.emit(EventTypes.SubscribeToRoom, editor.getRoomId())
})

socket.on(EventTypes.UserEdit, function (edit) {
  editor.serverEdit(edit)
})

socket.on(EventTypes.LanguageChange, function (change) {
  editor.setSyntax(change)
})
