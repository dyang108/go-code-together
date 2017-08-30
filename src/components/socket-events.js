import socket from './client-socket'
import cllbEditor from './CollabEditor'
import { EventTypes } from './lists'
import { changeUserCount, setUserCount } from './init-dom'

socket.on(EventTypes.Connection, () => {
  socket.emit(EventTypes.SubscribeToRoom, cllbEditor.getRoomId())
})

socket.on(EventTypes.UserEdit, edit => {
  cllbEditor.serverEdit(JSON.parse(edit))
})

socket.on(EventTypes.UserCountChange, change => {
  if (change === 'check') {
    socket.emit(EventTypes.GetRoomCount, cllbEditor.getRoomId(), userCount => {
      setUserCount(parseInt(userCount))
    })
  } else {
    changeUserCount(parseInt(change))
  }
})

socket.on(EventTypes.ChangeSelection, msg => {
  cllbEditor.serverChangeSelection(JSON.parse(msg))
})

socket.on(EventTypes.ChangeCursor, msg => {
  cllbEditor.serverChangeCursor(JSON.parse(msg))
})

socket.on(EventTypes.OtherUserDisconnect, change => {
  changeUserCount(-1)
  cllbEditor.removeOtherUser(change)
})

export default socket
