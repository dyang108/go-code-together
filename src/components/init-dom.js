import { languages, keybindings } from './lists'

var langSelect = document.getElementById('language-select')
var keybindingSelect = document.getElementById('keybinding-select')
var userCount = document.getElementById('user-count')

export function initForEditors (editors) {
  let populateLanguages = function () {
    languages.forEach(lang => {
      let opt = document.createElement('option')
      opt.setAttribute('value', lang.src)
      opt.innerHTML = lang.name
      langSelect.appendChild(opt)
    })
    if (mode) {
      editors.forEach(e => {
        e.setMode(mode)
      })
    }
  }

  let populateKeyBindings = function () {
    keybindings.forEach(kbtype => {
      let opt = document.createElement('option')
      opt.setAttribute('value', kbtype.src)
      opt.innerHTML = kbtype.name
      keybindingSelect.appendChild(opt)
    })
  }

  langSelect.onchange = function () {
    var selectedMode = langSelect.options[langSelect.selectedIndex].value
    editors.forEach(e => {
      e.changeHighlighting(selectedMode)
    })
  }

  keybindingSelect.onchange = function () {
    var selectedKeybinding = keybindingSelect.options[keybindingSelect.selectedIndex].value
    if (selectedKeybinding === 'ace') {
      selectedKeybinding = undefined
    }
    editors.forEach(e => {
      e.setKeyboardHandler(selectedKeybinding)
    })
  }

  populateLanguages()
  populateKeyBindings()
}

export function changeUserCount (change) {
  let currCount = parseInt(userCount.innerHTML)
  currCount += change
  setUserCount(currCount)
}

export function setUserCount (newCount) {
  userCount.innerHTML = newCount
}
