export var languages = [{
  name: 'JavaScript',
  src: 'ace/mode/javascript'
}, {
  name: 'Python',
  src: 'ace/mode/python'
}, {
  name: 'Java',
  src: 'ace/mode/java'
}, {
  name: 'C/C++',
  src: 'ace/mode/c_cpp'
}, {
  name: 'C#',
  src: 'ace/mode/csharp'
}, {
  name: 'Objective-C',
  src: 'ace/mode/objectivec'
}, {
  name: 'Go',
  src: 'ace/mode/golang'
}, {
  name: 'Scala',
  src: 'ace/mode/scala'
}, {
  name: 'Haskell',
  src: 'ace/mode/haskell'
}, {
  name: 'Ruby',
  src: 'ace/mode/ruby'
}, {
  name: 'Lisp',
  src: 'ace/mode/lisp'
}, {
  name: 'Scheme',
  src: 'ace/mode/scheme'
}, {
  name: 'Swift',
  src: 'ace/mode/swift'
}, {
  name: 'Rust',
  src: 'ace/mode/rust'
}, {
  name: 'Pascal',
  src: 'ace/mode/pascal'
}, {
  name: 'Clojure',
  src: 'ace/mode/clojure'
}, {
  name: 'MATLAB',
  src: 'ace/mode/matlab'
}, {
  name: 'Plain Text',
  src: 'ace/mode/text'
}, {
  name: 'Markdown',
  src: 'ace/mode/markdown'
}]

export var keybindings = [{
  name: 'Default',
  src: 'ace'
}, {
//   name: 'Sublime',
//   src: 'sublime'
// }, {
  name: 'Vim',
  src: 'ace/keyboard/vim'
}, {
  name: 'Emacs',
  src: 'ace/keyboard/emacs'
}]

export var EventTypes = {
  UserEdit: 'edit',
  LanguageChange: 'syntaxChange',
  SubscribeToRoom: 'room',
  Connection: 'connect',
  EditorChange: 'change',
  TextInsertion: 'insert',
  TextRemoval: 'remove',
  UserCountChange: 'countChange',
  GetRoomCount: 'getRoomCount',
  ChangeCursor: 'changeCursor',
  ChangeSelection: 'changeSelection',
  OtherUserDisconnect: 'otherUserDisconnect'
}
