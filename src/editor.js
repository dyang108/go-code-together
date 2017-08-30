import { initForEditors } from './components/init-dom'
import socket from './components/socket-events'
import cllbEditor from './components/CollabEditor'
import { EventTypes } from './components/lists'

initForEditors([cllbEditor])

socket.on(EventTypes.LanguageChange, change => {
  cllbEditor.setSyntax(JSON.parse(change))
})
