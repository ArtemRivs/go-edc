  // объявление типа
    type MyType int 
  // объявление метода
  func (m MyType) String() string{
        return fmt.Sprintf("MyType: %d", m)      
  }
  // вызов метода
  var m MyType = 5
  var s string = m.String()
