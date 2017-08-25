import socket from './client-socket'
import editor from './editor'
import { EventTypes } from './lists'
import { changeUserCount } from './init-dom'

socket.on(EventTypes.Connection, () => {
  socket.emit(EventTypes.SubscribeToRoom, editor.getRoomId())
})

socket.on(EventTypes.UserEdit, edit => {
  editor.serverEdit(edit)
})

socket.on(EventTypes.LanguageChange, change => {
  editor.setSyntax(change)
})

socket.on(EventTypes.UserCountChange, change => {
  changeUserCount(change)
})
