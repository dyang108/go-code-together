import { initForEditors } from './components/init-dom'
import socket from './components/socket-events'
import cllbEditor from './components/CollabEditor'
import Whiteboard from './components/Whiteboard'
import { EventTypes } from './components/lists'

initForEditors([cllbEditor, Whiteboard])

socket.on(EventTypes.LanguageChange, change => {
  cllbEditor.setSyntax(JSON.parse(change))
  Whiteboard.setSyntax(JSON.parse(change))
})
