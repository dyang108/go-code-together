import io from 'socket.io-client'
const socket = io(socketName)
socket.connect()
export default socket
