import io from 'socket.io-client'
import { EventTypes } from './lists'

const socket = io(socketName + '/index')
socket.connect()

let roomInput = document.getElementById('roomId')
let roomBtn = document.getElementById('roomBtn')
let countWarning = document.getElementById('countWarning')

function updateCount () {
  if (roomInput.value.includes(' ') || roomInput.value === '') {
    roomBtn.disabled = true
  } else {
    roomBtn.disabled = false
    socket.emit(EventTypes.TypeRoomName, roomInput.value, userCount => {
      if (!userCount) {
        countWarning.innerHTML = ''
        roomBtn.innerHTML = 'Create room'
      } else if (userCount === 1) {
        countWarning.innerHTML = 'There is 1 person in the room \'' + roomInput.value + '\'.'
        roomBtn.innerHTML = 'Join room'
      } else {
        countWarning.innerHTML = 'There are ' + userCount + ' people in the room \'' + roomInput.value + '\'.'
        roomBtn.innerHTML = 'Join room'
      }
    })
  }
}

roomInput.onkeyup = updateCount
roomInput.onchange = updateCount
