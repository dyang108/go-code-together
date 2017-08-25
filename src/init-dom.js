import Editor from './EditorWrapper'
import { languages, keybindings } from './lists'

var langSelect = document.getElementById('language-select')
var keybindingSelect = document.getElementById('keybinding-select')
var userCount = document.getElementById('user-count')

function populateLanguages () {
  languages.forEach(lang => {
    let opt = document.createElement('option')
    opt.setAttribute('value', lang.src)
    opt.innerHTML = lang.name
    langSelect.appendChild(opt)
  })
  if (mode) {
    Editor.setMode(mode)
  }
}

function populateKeyBindings () {
  keybindings.forEach(kbtype => {
    let opt = document.createElement('option')
    opt.setAttribute('value', kbtype.src)
    opt.innerHTML = kbtype.name
    keybindingSelect.appendChild(opt)
  })
}

export function changeUserCount (change) {
  let currCount = parseInt(userCount.innerHTML)
  currCount += change
  userCount.innerHTML = currCount
}

langSelect.onchange = function () {
  var selectedMode = langSelect.options[langSelect.selectedIndex].value
  Editor.changeHighlighting(selectedMode)
}

keybindingSelect.onchange = function () {
  var selectedKeybinding = keybindingSelect.options[keybindingSelect.selectedIndex].value
  if (selectedKeybinding === 'ace') {
    selectedKeybinding = undefined
  }
  Editor.setKeyboardHandler(selectedKeybinding)
}

populateLanguages()
populateKeyBindings()
