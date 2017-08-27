import socket from './client-socket'
import Editor from './EditorWrapper'
import { EventTypes } from './lists'
import { changeUserCount, setUserCount } from './init-dom'

socket.on(EventTypes.Connection, () => {
  socket.emit(EventTypes.SubscribeToRoom, Editor.getRoomId())
})

socket.on(EventTypes.UserEdit, edit => {
  Editor.serverEdit(JSON.parse(edit))
})

socket.on(EventTypes.LanguageChange, change => {
  Editor.setSyntax(JSON.parse(change))
})

socket.on(EventTypes.UserCountChange, change => {
  if (change === 'check') {
    socket.emit(EventTypes.GetRoomCount, Editor.getRoomId(), userCount => {
      setUserCount(parseInt(userCount))
    })
  } else {
    changeUserCount(parseInt(change))
  }
})

socket.on(EventTypes.ChangeSelection, msg => {
  Editor.serverChangeSelection(JSON.parse(msg))
})

socket.on(EventTypes.ChangeCursor, msg => {
  Editor.serverChangeCursor(JSON.parse(msg))
})

socket.on(EventTypes.OtherUserDisconnect, change => {
  changeUserCount(-1)
  Editor.removeOtherUser(change)
})
