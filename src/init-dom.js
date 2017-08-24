import editor from './editor'
import { languages, keybindings } from './lists'

var langSelect = document.getElementById('language-select')
var keybindingSelect = document.getElementById('keybinding-select')

function populateLanguages () {
  languages.forEach(lang => {
    let opt = document.createElement('option')
    opt.setAttribute('value', lang.src)
    opt.innerHTML = lang.name
    langSelect.appendChild(opt)
  })
  if (mode) {
    editor.setMode(mode)
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

langSelect.onchange = function () {
  var selectedMode = langSelect.options[langSelect.selectedIndex].value
  editor.changeHighlighting(selectedMode)
}

keybindingSelect.onchange = function () {
  var selectedKeybinding = keybindingSelect.options[keybindingSelect.selectedIndex].value
  if (selectedKeybinding === 'ace') {
    selectedKeybinding = undefined
  }
  editor.setKeyboardHandler(selectedKeybinding)
}

populateLanguages()
populateKeyBindings()